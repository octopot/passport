package service

import (
	"github.com/kamilsk/passport/domain"
	"github.com/kamilsk/passport/transfer/api/v1/tracker"
)

// New returns a new instance of Passport service.
func New(dao Storage) *Passport {
	return &Passport{dao: dao}
}

// Passport is the primary application service.
type Passport struct {
	dao Storage
}

// HandleTrackerInstructionV1 handles an input request.
func (s *Passport) HandleTrackerInstructionV1(request tracker.InstructionRequest) tracker.InstructionResponse {
	var response tracker.InstructionResponse

	{ // TODO encrypt/decrypt marker
		marker := domain.UUID(request.Marker)
		if !marker.IsValid() {
			marker, response.Error = s.dao.UUID()
		}
		response.Marker = string(marker)
	}

	return response
}

// HandleTrackerFingerprintV1 handles an input request.
func (s *Passport) HandleTrackerFingerprintV1(tracker.FingerprintRequest) tracker.FingerprintResponse {
	var response tracker.FingerprintResponse
	return response
}
