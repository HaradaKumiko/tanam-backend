package base

type BaseErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func NewErrorResponse(message string) *BaseErrorResponse {
	return &BaseErrorResponse{
		Success: false,
		Message: message,
	}
}
