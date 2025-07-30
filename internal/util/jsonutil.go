package util

import (
	"encoding/json"
	"errors"
	"io"
)

func DecodeJSON(body io.ReadCloser, v any) error {
	data, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	if len(data) < 1 {
		return errors.New("Missing request body")
	}
	err = json.Unmarshal(data, v)
	if err != nil {
		return err
	}
	return nil
}

func EncodeJson(v any) ([]byte, error) {
	if data, err := json.Marshal(v); err != nil {
		return nil, err
	} else {
		return data, nil
	}
}
