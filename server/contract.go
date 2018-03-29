package server

import "github.com/kamilsk/passport/transfer/api/v1/tracker"

// Service defines the behavior of Passport service.
type Service interface {
	// HandleTrackerInstructionV1 handles an input request.
	HandleTrackerInstructionV1(tracker.InstructionRequest) tracker.InstructionResponse
	// HandleTrackerFingerprintV1 handles an input request.
	HandleTrackerFingerprintV1(tracker.FingerprintRequest) tracker.FingerprintResponse
}
