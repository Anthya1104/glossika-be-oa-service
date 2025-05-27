package util

import "encoding/json"

// Safe log struct by converting it to JSON string
func StructToJsonString(data interface{}) string {
	jsonData, _ := json.Marshal(data)
	return string(jsonData)
}
