/**
 * Golang equivalent to php `str_ends_with`
 *
 * Checks if a string ends with a given substring
 * @see https://www.php.net/manual/en/function.str-ends-with.php
 *
 * @param string haystack
 * @param string needle
 * @return bool
 */
package functions

func StrEndsWith(s string, needle string) bool {
	s_len := len(s)
	n_len := len(needle)
	if n_len > s_len {
		return false
	}
	for pos := 0; pos < n_len; pos++ {
		if s[s_len-1-pos] != needle[n_len-1-pos] {
			return false
		}
	}
	return true
}
