package router

import "net/http"

// Server defines the behavior of Passport server.
type Server interface {
	// GetTrackerInstructionV1 is responsible for `GET /api/v1/tracker/instruction` request handling.
	GetTrackerInstructionV1(http.ResponseWriter, *http.Request)
	// PostTrackerFingerprintV1 is responsible for `POST /api/v1/tracker/fingerprint` request handling.
	PostTrackerFingerprintV1(http.ResponseWriter, *http.Request)
}
