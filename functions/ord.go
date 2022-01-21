/**
 * Golang equivalent to php `ord`
 *
 * Convert the first byte of a string to a value between 0 and 255
 * @see https://www.php.net/manual/en/function.ord.php
 *
 * @param string character
 * @return int
 */
package functions

func Ord(character string) int {
	if len(character) <= 0 {
		return 0
	}
	return int(character[0])
}
