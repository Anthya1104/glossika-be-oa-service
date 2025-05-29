package util

import "encoding/json"

// Safe log struct by converting it to JSON string
func StructToJsonString(data interface{}) string {
	jsonData, _ := json.Marshal(data)
	return string(jsonData)
}

func SliceDataByPaging[T any](products []T, offset, limit int) ([]T, int) {
	total := len(products)
	if offset > total {
		return []T{}, total
	}
	end := offset + limit
	if end > total {
		end = total
	}
	return products[offset:end], total
}
