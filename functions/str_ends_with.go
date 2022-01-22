package functions

// Checks if a string ends with a given substring
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
