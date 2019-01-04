package service

import (
	"context"
	core "errors"

	"github.com/kamilsk/passport/pkg/config"
	"github.com/kamilsk/passport/pkg/domain"
	"github.com/kamilsk/passport/pkg/errors"
	"github.com/kamilsk/passport/pkg/transfer/api/v1/tracker"
)

// New returns a new instance of Passport service.
func New(_ config.ServiceConfig, storage Storage) *Passport {
	return &Passport{storage: storage}
}

// Passport is the primary application service.
type Passport struct {
	storage Storage
}

// HandleTrackerFingerprintV1 handles an input request.
func (s *Passport) HandleTrackerFingerprintV1(request tracker.FingerprintRequest) tracker.FingerprintResponse {
	var response tracker.FingerprintResponse

	{ // TODO encrypt/decrypt session
		session := domain.UUID(request.EncryptedSession)
		if !session.IsValid() {
			response.Error = errors.Validation(errors.ClientErrorMessage, core.New("invalid session"),
				"trying to validate user session %q", session)
			return response
		}
		fingerprint := domain.Fingerprint{Session: session, Marker: request.Payload.Fingerprint}
		if !fingerprint.IsValid() {
			response.Error = errors.Validation(errors.ClientErrorMessage, core.New("invalid fingerprint"),
				"trying to validate user fingerprint %q", fingerprint.Marker)
			return response
		}
		response.Fingerprint, response.Error = s.storage.StoreFingerprint(context.TODO(), fingerprint)
	}

	return response
}

// HandleTrackerInstructionV1 handles an input request.
func (s *Passport) HandleTrackerInstructionV1(request tracker.InstructionRequest) tracker.InstructionResponse {
	var response tracker.InstructionResponse

	{ // TODO encrypt/decrypt session
		session := domain.UUID(request.EncryptedSession)
		if !session.IsValid() {
			session, response.Error = s.storage.UUID(context.TODO())
		}
		response.EncryptedSession = session.String()
	}

	return response
}
