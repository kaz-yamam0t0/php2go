package functions

// Convert the first byte of a string to a value between 0 and 255
func Ord(character string) int {
	if len(character) <= 0 {
		return 0
	}
	return int(character[0])
}
