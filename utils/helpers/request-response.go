package helpers

func SuccessGetResponseData(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"success": true,
		"message": "Success",
		"data":    data,
	}
}

func SuccessActionResponse(msg string) map[string]interface{} {
	return map[string]interface{}{
		"success": true,
		"message": msg,
	}
}

func FailedResponse(message string) map[string]interface{} {
	return map[string]interface{}{
		"success": false,
		"message": message,
		"data":    map[string]interface{}{},
	}
}
