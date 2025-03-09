package utils

func GenerateMapForResponseType(key string, message string, output interface{}) map[string]interface{} {
	quickMap := make(map[string]interface{})
	if key == "Error" {
		quickMap["message"] = message
	}
	quickMap[key] = output

	return quickMap
}
