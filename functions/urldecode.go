/**
 * Golang equivalent to php `urldecode`
 *
 * Decodes URL-encoded string
 * @see https://www.php.net/manual/en/function.urldecode.php
 *
 * @param string s
 * @return string
 */
package functions

import "net/url"

func Urldecode(s string) (string, error) {
	return url.QueryUnescape(s)
}
