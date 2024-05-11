package response

type ErrorResponse struct {
	Info map[string]interface{} `json:"info"`
	Data interface{}            `json:"data"`
}

func ErrorFormatter(message string) ErrorResponse {
	response := ErrorResponse{
		Info: map[string]interface{}{
			"success": false,
			"message": message,
		},
		Data: nil,
	}

	return response
}
