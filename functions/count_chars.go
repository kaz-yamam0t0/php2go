package functions

// Return information about characters used in a string
func CountChars(s string, args ...int) [256]int {
	/*	mode := 0
		if len(args) > 0 {
			mode = args[0]
		} */
	var dst [256]int
	s_len := len(s)
	for i := 0; i < s_len; i++ {
		dst[s[i]]++
	}
	return dst
}
