/**
 * Golang equivalent to php `str_starts_with`
 *
 * Checks if a string starts with a given substring
 * @see https://www.php.net/manual/en/function.str-starts-with.php
 *
 * @param string haystack
 * @param string needle
 * @return bool
 */
package functions

func StrStartsWith(s string, needle string) bool {
	s_len := len(s)
	n_len := len(needle)
	if n_len > s_len {
		return false
	}
	for pos := 0; pos < n_len; pos++ {
		if s[pos] != needle[pos] {
			return false
		}
	}
	return true
}
