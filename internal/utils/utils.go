package utils

import (
	"database/sql/driver"
	"errors"
)

func GenerateMapForResponseType(key string, message string, output interface{}) map[string]interface{} {
	quickMap := make(map[string]interface{})
	if key == "Error" {
		quickMap["message"] = message
	}
	quickMap[key] = output

	return quickMap
}

// ✅ Define JSONB type properly
type JSONB []byte

// ✅ Implement Scanner interface to scan JSONB from DB
func (j *JSONB) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("JSONB: cannot scan non-byte value")
	}
	*j = JSONB(bytes) // Assign byte slice to JSONB type
	return nil
}

// ✅ Implement Value interface to store JSONB in DB
func (j JSONB) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return []byte(j), nil
}
