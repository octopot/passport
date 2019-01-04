package tracker

// InstructionRequest represents `GET /api/v1/tracker/instruction` request.
type InstructionRequest struct {
	EncryptedSession string
}

// InstructionResponse represents `GET /api/v1/tracker/instruction` response.
type InstructionResponse struct {
	EncryptedSession string
	Error            error
}
