package server

import "net/http"

// New returns a new instance of Passport server.
func New(service Service) *Server {
	return &Server{service: service}
}

// Server handles HTTP requests.
type Server struct {
	service Service
}

// GetTrackerInstructionV1 is responsible for `GET /api/v1/tracker/instruction` request handling.
func (s *Server) GetTrackerInstructionV1(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte(req.RequestURI))
}

// PostTrackerFingerprintV1 is responsible for `POST /api/v1/tracker/fingerprint` request handling.
func (s *Server) PostTrackerFingerprintV1(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte(req.RequestURI))
}
