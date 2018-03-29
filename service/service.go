package service

import "github.com/kamilsk/passport/transfer/api/v1/tracker"

// New returns a new instance of Passport service.
func New(dao Storage) *Passport {
	return &Passport{dao: dao}
}

// Passport is the primary application service.
type Passport struct {
	dao Storage
}

// HandleTrackerInstructionV1 handles an input request.
func (s *Passport) HandleTrackerInstructionV1(tracker.InstructionRequest) tracker.InstructionResponse {
	var response tracker.InstructionResponse
	return response
}

// HandleTrackerFingerprintV1 handles an input request.
func (s *Passport) HandleTrackerFingerprintV1(tracker.FingerprintRequest) tracker.FingerprintResponse {
	var response tracker.FingerprintResponse
	return response
}
