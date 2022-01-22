package functions

// Generate a single-byte string from a number
func Chr(codepoint int) string {
	for codepoint < 0 {
		codepoint += 256
	}
	codepoint %= 256

	return string(byte(codepoint))
}
