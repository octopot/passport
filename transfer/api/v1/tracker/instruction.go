package tracker

import "github.com/kamilsk/passport/domain"

// RedirectRequest represents `GET /api/v1/tracker/instruction` request.
type InstructionRequest struct {
	Marker domain.UUID
}

// RedirectRequest represents `GET /api/v1/tracker/instruction` response.
type InstructionResponse struct {
	Marker domain.UUID
	Error  error
}
