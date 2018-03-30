package server

import (
	"encoding/json"
	"html/template"
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
	if err != nil {
		cookie = &http.Cookie{Name: MarkerKey}
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
	s.template.Execute(rw, struct {
		BaseURL   *url.URL
		Endpoint  string
		Threshold uint8
		Correct   int
		Watch     int
		Debug     bool
	}{
		BaseURL:   s.baseURL,
		Endpoint:  "/api/v1/tracker/fingerprint",
		Threshold: 3,
		Correct:   100,  // Milliseconds
		Watch:     1000, // Milliseconds
		Debug:     false,
	})
}

// PostTrackerFingerprintV1 is responsible for `POST /api/v1/tracker/fingerprint` request handling.
func (s *Server) PostTrackerFingerprintV1(rw http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie(MarkerKey)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	defer req.Body.Close()
	request := tracker.FingerprintRequest{EncryptedMarker: cookie.Value, Header: req.Header}
	if err := json.NewDecoder(req.Body).Decode(&request.Payload); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

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
