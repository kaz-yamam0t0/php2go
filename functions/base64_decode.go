package functions

import (
	b64 "encoding/base64"
)

// Decodes data encoded with MIME base64
func Base64Decode(s string) (string, error) {
	b, err := b64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(b), nil

}
