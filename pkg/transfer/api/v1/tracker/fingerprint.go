package tracker

import "github.com/kamilsk/passport/pkg/domain"

// FingerprintRequest represents `GET /api/v1/tracker/fingerprint` request.
type FingerprintRequest struct {
	EncryptedSession string
	Header           map[string][]string
	Payload          struct {
		Fingerprint string
	}
}

// FingerprintResponse represents `GET /api/v1/tracker/fingerprint` response.
type FingerprintResponse struct {
	Fingerprint domain.Fingerprint
	Error       error
}
