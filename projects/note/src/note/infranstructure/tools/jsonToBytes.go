package tools

import (
	"encoding/json"
)

func JsonToBytes(p interface{}) ([]byte, error) {
	bytesJson, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return bytesJson, nil
}
