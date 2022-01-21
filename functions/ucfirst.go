/**
 * Golang equivalent to php `ucfirst`
 *
 * Make a string's first character uppercase
 * @see https://www.php.net/manual/en/function.ucfirst.php
 *
 * @param string s
 * @return string
 */
package functions

func Ucfirst(s string) string {
	if len(s) > 0 && 'a' <= s[0] && s[0] <= 'z' {
		return string(s[0]-'a'+'A') + s[1:]
	}
	return s
}
