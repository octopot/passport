package server

import (
	"html/template"
	"net/http"
	"net/url"

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
	response := s.service.HandleTrackerInstructionV1(tracker.InstructionRequest{Marker: cookie.Value})
	if response.Error != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	{ // domain logic
		cookie.MaxAge, cookie.Path, cookie.Value = 0, "/", response.Marker
		http.SetCookie(rw, cookie)
	}

	rw.Header().Set("Content-Type", "application/javascript")
	rw.WriteHeader(http.StatusOK)
	s.template.Execute(rw, struct{ BaseURL *url.URL }{s.baseURL})
}

// PostTrackerFingerprintV1 is responsible for `POST /api/v1/tracker/fingerprint` request handling.
func (s *Server) PostTrackerFingerprintV1(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte(req.RequestURI))
}

func passport() string {
	b, err := static.Asset("static/scripts/passport.js")
	if err != nil {
		panic(err)
	}
	return string(b)
}
