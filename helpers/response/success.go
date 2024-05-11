package response

type SuccessResponse struct {
	Info map[string]interface{} `json:"info"`
	Data interface{}            `json:"data"`
}

func SuccessSingularFormatter(message string, data interface{}) SuccessResponse {
	response := SuccessResponse{
		Info: map[string]interface{}{
			"success": true,
			"message": message,
		},
		Data: data,
	}

	return response
}

func SuccessPluralFormatter(message string, data interface{}) SuccessResponse {
	response := SuccessResponse{
		Info: map[string]interface{}{
			"message": message,
			"success": true,
		},
		Data: data,
	}

	return response
}
