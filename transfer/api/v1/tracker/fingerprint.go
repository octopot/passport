package tracker

import "github.com/kamilsk/passport/domain"

// FingerprintRequest represents `GET /api/v1/tracker/fingerprint` request.
type FingerprintRequest struct {
	Fingerprint domain.Fingerprint
}

// FingerprintRequest represents `GET /api/v1/tracker/fingerprint` response.
type FingerprintResponse struct {
	Fingerprint domain.Fingerprint
	Error       error
}
