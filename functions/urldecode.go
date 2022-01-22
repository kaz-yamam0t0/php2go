package functions

import "net/url"

// Decodes URL-encoded string
func Urldecode(s string) (string, error) {
	return url.QueryUnescape(s)
}
