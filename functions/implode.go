package functions

import "strings"

// Join array elements with a string
func Implode(separator string, a []string) string {
	return strings.Join(a, separator)
}
