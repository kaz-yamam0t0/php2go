package functions

// utils

func lower(c byte) byte {
	if 'A' <= c && c <= 'Z' {
		return c + 'a' - 'A'
	}
	return c
}
func cmpLower(a byte, b byte) bool {
	a = lower(a)
	b = lower(b)
	return (a - b) == 0
}

func urlShouldEscape(c byte) bool {
	if c > 0x7E {
		return true
	}
	if 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || '0' <= c && c <= '9' {
		return false
	}
	switch c {
	case '-', '_', '.', '~':
		return false
	}
	return true
}

// used in Stripcslashes
func is_digit(c byte) bool {
	return '0' <= c && c <= '9'
}
func is_hex(c byte) bool {
	return ('0' <= c && c <= '9') || ('a' <= c && c <= 'f') || ('A' <= c && c <= 'F')
}
func c_hex(c byte) byte {
	if '0' <= c && c <= '9' {
		return c - '0'
	}
	if 'a' <= c && c <= 'f' {
		return c - 'a' + 0x0A
	}
	if 'A' <= c && c <= 'F' {
		return c - 'A' + 0x0A
	}
	return 0
}

// Constants for escaping or unescaping html strings.
const (
	ENT_COMPAT     = 2   // Will convert double-quotes and leave single-quotes alone.
	ENT_QUOTES     = 3   // Will convert both double and single quotes.
	ENT_NOQUOTES   = 0   // Will leave both double and single quotes unconverted.
	ENT_IGNORE     = 4   // Silently discard invalid code unit sequences instead of returning an empty string. Using this flag is discouraged as it may have security implications.
	ENT_DISALLOWED = 128 // Replace invalid code points for the given document type with a Unicode Replacement Character U+FFFD (UTF-8) or &#xFFFD; (otherwise) instead of leaving them as is. This may be useful, for instance, to ensure the well-formedness of XML documents with embedded external content.
	ENT_SUBSTITUTE = 8   // Replace invalid code unit sequences with a Unicode Replacement Character U+FFFD (UTF-8) or &#xFFFD; (otherwise) instead of returning an empty string.
	ENT_HTML401    = 0   // Handle code as HTML 4.01.
	ENT_XML1       = 16  // Handle code as XML 1.
	ENT_XHTML      = 32  // Handle code as XHTML.
	ENT_HTML5      = 48  // Handle code as HTML 5.
)

// Constants for padding string chars. (used in `StrPad()`)
const (
	STR_PAD_RIGHT = 1
	STR_PAD_LEFT  = 0
	STR_PAD_BOTH  = 2
)

const upperhex = "0123456789ABCDEF"
