package functions

import (
	b64 "encoding/base64"
)

// Encodes data with MIME base64
func Base64Encode(s string) string {
	return b64.StdEncoding.EncodeToString([]byte(s))
}
