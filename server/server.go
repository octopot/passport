package server

import (
	"encoding/json"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/kamilsk/passport/errors"
	"github.com/kamilsk/passport/static"
	"github.com/kamilsk/passport/transfer/api/v1/tracker"
)

const (
	// MarkerKey is used to find and store required cookie value.
	MarkerKey = "marker"
)

// New returns a new instance of Passport server.
func New(baseURL string, service Service) *Server {
	u, err := url.Parse(baseURL)
	if err != nil {
		panic(err)
	}
	return &Server{baseURL: u, service: service, template: template.Must(template.New("script").Parse(passport()))}
}

// Server handles HTTP requests.
type Server struct {
	baseURL  *url.URL
	service  Service
	template *template.Template
}

// GetTrackerInstructionV1 is responsible for `GET /api/v1/tracker/instruction` request handling.
func (s *Server) GetTrackerInstructionV1(rw http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie(MarkerKey)
	if err != nil || !cookie.HttpOnly || !cookie.Secure {
		cookie = &http.Cookie{Name: MarkerKey, Secure: true, HttpOnly: true}
	}

	response := s.service.HandleTrackerInstructionV1(tracker.InstructionRequest{EncryptedMarker: cookie.Value})
	if response.Error != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	cookie.MaxAge, cookie.Path, cookie.Value = 0, "/", response.EncryptedMarker
	http.SetCookie(rw, cookie)
	rw.Header().Set("Content-Type", "application/javascript")
	rw.WriteHeader(http.StatusOK)

	endpoint := *s.baseURL
	endpoint.Path = "/api/v1/tracker/fingerprint"
	s.template.Execute(rw, struct {
		Endpoint  string
		Limit     uint8
		Threshold uint8
		Correct   int // Milliseconds
		Watch     int // Milliseconds
		Debug     bool
	}{
		Endpoint: endpoint.String(),
		Limit:    5, Threshold: 3,
		Correct: 250, Watch: 1000,
		Debug: false,
	})
}

// PostTrackerFingerprintV1 is responsible for `POST /api/v1/tracker/fingerprint` request handling.
func (s *Server) PostTrackerFingerprintV1(rw http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie(MarkerKey)
	if err != nil {
		// issue #19: Safari sends cookies in `demo-cross-origin`-mode, but doesn't send it in production
		log.Printf("\n\n[CRITICAL] cookie not found, skip this request (%q)\n\n", req.UserAgent())
		rw.WriteHeader(http.StatusAccepted)
		io.Copy(ioutil.Discard, req.Body)
		req.Body.Close()
		return
	}
	if !cookie.HttpOnly || !cookie.Secure {
		// issue #22: prevent cookie manipulation
		log.Printf("\n\n[CRITICAL] cookie is not safe, skip this request (%+v)\n\n", *cookie)
		io.Copy(ioutil.Discard, req.Body)
		req.Body.Close()
		return
	}

	defer req.Body.Close()
	request := tracker.FingerprintRequest{EncryptedMarker: cookie.Value, Header: req.Header}
	if err := json.NewDecoder(req.Body).Decode(&request.Payload); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		io.Copy(ioutil.Discard, req.Body)
		req.Body.Close()
		return
	}
	req.Body.Close()

	response := s.service.HandleTrackerFingerprintV1(request)
	if response.Error != nil {
		if err, is := response.Error.(errors.ApplicationError); is {
			if _, is := err.IsClientError(); is {
				rw.WriteHeader(http.StatusBadRequest)
				return
			}
		}
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	if response.Fingerprint.UpdatedAt.Valid {
		rw.WriteHeader(http.StatusOK)
		return
	}
	rw.WriteHeader(http.StatusCreated)
}

func passport() string {
	b, err := static.Asset("static/scripts/passport.min.js")
	if err != nil {
		panic(err)
	}
	return string(b)
}
