package utils

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

func GenerateHashFromJSON(data interface{}) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256(jsonData)

	return fmt.Sprintf("%x", hash), nil
}
