package utils

import (
	"encoding/json"
	"io"
)

func ToJSON(r io.Reader, obj any) error {
	return json.NewDecoder(r).Decode(&obj)
}
