package service

import (
	origin "errors"

	"github.com/kamilsk/passport/domain"
	"github.com/kamilsk/passport/errors"
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
		marker := domain.UUID(request.EncryptedMarker)
		if !marker.IsValid() {
			marker, response.Error = s.dao.UUID()
		}
		response.EncryptedMarker = string(marker)
	}

	return response
}

// HandleTrackerFingerprintV1 handles an input request.
func (s *Passport) HandleTrackerFingerprintV1(request tracker.FingerprintRequest) tracker.FingerprintResponse {
	var response tracker.FingerprintResponse

	{ // TODO encrypt/decrypt marker
		marker := domain.UUID(request.EncryptedMarker)
		if !marker.IsValid() {
			response.Error = errors.Validation(errors.ClientErrorMessage, origin.New("invalid marker"),
				"trying to validate user marker %q", marker)
			return response
		}
		fingerprint := domain.Fingerprint{Marker: string(marker), Value: request.Payload.Fingerprint}
		if !fingerprint.IsValid() {
			response.Error = errors.Validation(errors.ClientErrorMessage, origin.New("invalid fingerprint"),
				"trying to validate user fingerprint %q", fingerprint.Value)
			return response
		}
		response.Fingerprint, response.Error = s.dao.TakeFingerprint(fingerprint)
	}

	return response
}
