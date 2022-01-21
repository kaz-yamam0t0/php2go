/**
 * Golang equivalent to php `lcfirst`
 *
 * Make a string's first character lowercase
 * @see https://www.php.net/manual/en/function.lcfirst.php
 *
 * @param string s
 * @return string
 */
package functions

func Lcfirst(s string) string {
	if len(s) > 0 && 'A' <= s[0] && s[0] <= 'Z' {
		return string(s[0]-'A'+'a') + s[1:]
	}
	return s
}
