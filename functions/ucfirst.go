package functions

// Make a string's first character uppercase
func Ucfirst(s string) string {
	if len(s) > 0 && 'a' <= s[0] && s[0] <= 'z' {
		return string(s[0]-'a'+'A') + s[1:]
	}
	return s
}
