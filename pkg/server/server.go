package server

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"text/template"

	"github.com/kamilsk/passport/pkg/config"
	"github.com/kamilsk/passport/pkg/errors"
	"github.com/kamilsk/passport/pkg/static"
	"github.com/kamilsk/passport/pkg/transfer/api/v1/tracker"
)

const sessionKey = "session"

// New returns a new instance of Passport server.
func New(cnf config.ServerConfig, service Service) *Server {
	u, err := url.Parse(cnf.BaseURL)
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
	cookie, err := req.Cookie(sessionKey)
	if err != nil {
		cookie = &http.Cookie{Name: sessionKey, HttpOnly: true, Secure: s.baseURL.Scheme == "https"}
	}

	response := s.service.HandleTrackerInstructionV1(tracker.InstructionRequest{EncryptedSession: cookie.Value})
	if response.Error != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	cookie.MaxAge, cookie.Path, cookie.Value = 0, "/", response.EncryptedSession
	http.SetCookie(rw, cookie)
	rw.Header().Set("Content-Type", "application/javascript")
	rw.WriteHeader(http.StatusOK)

	endpoint := *s.baseURL
	endpoint.Path = "/api/v1/tracker/fingerprint"
	_ = s.template.Execute(rw, struct {
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
	defer func() {
		_, _ = io.Copy(ioutil.Discard, req.Body)
		_ = req.Body.Close()
	}()
	cookie, err := req.Cookie(sessionKey)
	if err != nil {
		// Related articles:
		// - https://blog.mozilla.org/futurereleases/2018/08/30/changing-our-approach-to-anti-tracking/
		// - https://webkit.org/blog/7675/intelligent-tracking-prevention/
		log.Printf("\n\n[CRITICAL] cookie not found, skip this request (%q)\n\n", req.UserAgent())
		rw.WriteHeader(http.StatusAccepted)
		_, _ = io.Copy(ioutil.Discard, req.Body)
		_ = req.Body.Close()
		return
	}

	request := tracker.FingerprintRequest{EncryptedSession: cookie.Value, Header: req.Header}
	if err := req.ParseForm(); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	request.Payload.Fingerprint = req.PostFormValue("fingerprint")

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

	rw.WriteHeader(http.StatusCreated)
}

func passport() string {
	b, err := static.Asset("static/scripts/passport.min.js")
	if err != nil {
		panic(err)
	}
	return string(b)
}
