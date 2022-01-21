/**
 * Golang equivalent to php `rawurldecode`
 *
 * Decode URL-encoded strings
 * @see https://www.php.net/manual/en/function.rawurldecode.php
 *
 * @param string s
 * @return string
 */
package functions

import "net/url"

func Rawurldecode(s string) (string, error) {
	return url.PathUnescape(s)
}
