/**
 * Golang equivalent to php `chr`
 *
 * Generate a single-byte string from a number
 * @see https://www.php.net/manual/en/function.chr.php
 *
 * @param int codepoint
 * @return string
 */
package functions

func Chr(codepoint int) string {
	for codepoint < 0 {
		codepoint += 256
	}
	codepoint %= 256

	return string(byte(codepoint))
}
