package tracker

// RedirectRequest represents `GET /api/v1/tracker/instruction` request.
type InstructionRequest struct {
	Marker string
}

// RedirectRequest represents `GET /api/v1/tracker/instruction` response.
type InstructionResponse struct {
	Marker string
	Error  error
}
