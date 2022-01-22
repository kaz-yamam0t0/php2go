package functions

import "net/url"

// Decode URL-encoded strings
func Rawurldecode(s string) (string, error) {
	return url.PathUnescape(s)
}
