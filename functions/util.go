package functions

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


/**
 * ENT_COMPAT     : Will convert double-quotes and leave single-quotes alone.
 * ENT_QUOTES     : Will convert both double and single quotes.
 * ENT_NOQUOTES   : Will leave both double and single quotes unconverted.
 * ENT_IGNORE     : Silently discard invalid code unit sequences instead of returning an empty string. Using this flag is discouraged as it may have security implications.
 * ENT_DISALLOWED : Replace invalid code points for the given document type with a Unicode Replacement Character U+FFFD (UTF-8) or &#xFFFD; (otherwise) instead of leaving them as is. This may be useful, for instance, to ensure the well-formedness of XML documents with embedded external content.
 * ENT_SUBSTITUTE : Replace invalid code unit sequences with a Unicode Replacement Character U+FFFD (UTF-8) or &#xFFFD; (otherwise) instead of returning an empty string.
 * ENT_HTML401    : Handle code as HTML 4.01.
 * ENT_XML1       : Handle code as XML 1.
 * ENT_XHTML      : Handle code as XHTML.
 * ENT_HTML5      : Handle code as HTML 5.
 */
 const (
	ENT_COMPAT     = 2
	ENT_QUOTES     = 3
	ENT_NOQUOTES   = 0
	ENT_IGNORE     = 4
	ENT_DISALLOWED = 128
	ENT_SUBSTITUTE = 8
	ENT_HTML401    = 0
	ENT_XML1       = 16
	ENT_XHTML      = 32
	ENT_HTML5      = 48
)

const upperhex = "0123456789ABCDEF"
