package tracker

import "github.com/kamilsk/passport/domain"

// FingerprintRequest represents `GET /api/v1/tracker/fingerprint` request.
type FingerprintRequest struct {
	EncryptedMarker string
	Header          map[string][]string
	Payload         struct {
		Fingerprint string                   `json:"fingerprint"`
		Metadata    []map[string]interface{} `json:"metadata"`
	}
}

// FingerprintRequest represents `GET /api/v1/tracker/fingerprint` response.
type FingerprintResponse struct {
	Fingerprint domain.Fingerprint
	Error       error
}
