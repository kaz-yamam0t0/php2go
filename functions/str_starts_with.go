package functions

// Checks if a string starts with a given substring
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
