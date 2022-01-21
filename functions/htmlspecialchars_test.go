package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHtmlspecialchars(t *testing.T) {

	// Tests for quotes
	assert.Equal(t, "&lt;a href=&#039;test&#039;&gt;&quot;Test&quot;&lt;/a&gt;", Htmlspecialchars("<a href='test'>\"Test\"</a>", ENT_QUOTES))
	assert.Equal(t, "&lt;a href='test'&gt;&quot;Test&quot;&lt;/a&gt;", Htmlspecialchars("<a href='test'>\"Test\"</a>", ENT_COMPAT))
	assert.Equal(t, "&lt;a href='test'&gt;\"Test\"&lt;/a&gt;", Htmlspecialchars("<a href='test'>\"Test\"</a>", ENT_NOQUOTES))

	// Returns an empty string if contains invalid code unit sequences
	assert.Equal(t, "", Htmlspecialchars("\x81\x81<>", ENT_QUOTES, "utf-8"))

	// Convert invalid code unit sequences to ï¿½ (U+FFFD) with `ENT_SUBSTITUTE`
	assert.Equal(t, "\xef\xbf\xbd\xef\xbf\xbd&lt;&gt;", Htmlspecialchars("\x81\x81<>", ENT_SUBSTITUTE, "utf-8"))

	// Usually invalid chars within Non-UTF-8 strings are converted to "&#xFFFD;" but seem to be ignored when called from the command line.
	//assert.Equal(t, "&#xFFFD;&#xFFFD;&lt;&gt;", Htmlspecialchars("\x81\x81<>",ENT_SUBSTITUTE,"Shift_JIS"))
	assert.Equal(t, "\x81\x81&lt;&gt;", Htmlspecialchars("\x81\x81<>", ENT_SUBSTITUTE, "Shift_JIS"))

	// HTML entities tests for ENT_HTML401
	assert.Equal(t, "&amp;", Htmlspecialchars("&amp;", ENT_NOQUOTES|ENT_SUBSTITUTE|ENT_HTML401, "utf-8", false))
	assert.Equal(t, "&quot;", Htmlspecialchars("&quot;", ENT_NOQUOTES|ENT_SUBSTITUTE|ENT_HTML401, "utf-8", false))
	assert.Equal(t, "&amp;apos;", Htmlspecialchars("&apos;", ENT_NOQUOTES|ENT_SUBSTITUTE|ENT_HTML401, "utf-8", false))
	assert.Equal(t, "&lt;", Htmlspecialchars("&lt;", ENT_NOQUOTES|ENT_SUBSTITUTE|ENT_HTML401, "utf-8", false))
	assert.Equal(t, "&gt;", Htmlspecialchars("&gt;", ENT_NOQUOTES|ENT_SUBSTITUTE|ENT_HTML401, "utf-8", false))
	assert.Equal(t, "&#x26;", Htmlspecialchars("&#x26;", ENT_NOQUOTES|ENT_SUBSTITUTE|ENT_HTML401, "utf-8", false))
	assert.Equal(t, "&amp;plus;", Htmlspecialchars("&plus;", ENT_NOQUOTES|ENT_SUBSTITUTE|ENT_HTML401, "utf-8", false))

	// HTML entities tests for ENT_XML1
	assert.Equal(t, "&amp;", Htmlspecialchars("&amp;", ENT_NOQUOTES|ENT_SUBSTITUTE|ENT_XML1, "utf-8", false))
	assert.Equal(t, "&quot;", Htmlspecialchars("&quot;", ENT_NOQUOTES|ENT_SUBSTITUTE|ENT_XML1, "utf-8", false))
	assert.Equal(t, "&apos;", Htmlspecialchars("&apos;", ENT_NOQUOTES|ENT_SUBSTITUTE|ENT_XML1, "utf-8", false))
	assert.Equal(t, "&lt;", Htmlspecialchars("&lt;", ENT_NOQUOTES|ENT_SUBSTITUTE|ENT_XML1, "utf-8", false))
	assert.Equal(t, "&gt;", Htmlspecialchars("&gt;", ENT_NOQUOTES|ENT_SUBSTITUTE|ENT_XML1, "utf-8", false))
	assert.Equal(t, "&#x26;", Htmlspecialchars("&#x26;", ENT_NOQUOTES|ENT_SUBSTITUTE|ENT_XML1, "utf-8", false))
	assert.Equal(t, "&amp;plus;", Htmlspecialchars("&plus;", ENT_NOQUOTES|ENT_SUBSTITUTE|ENT_XML1, "utf-8", false))

	// HTML entities tests for ENT_XHTML
	assert.Equal(t, "&amp;", Htmlspecialchars("&amp;", ENT_NOQUOTES|ENT_SUBSTITUTE|ENT_XHTML, "utf-8", false))
	assert.Equal(t, "&quot;", Htmlspecialchars("&quot;", ENT_NOQUOTES|ENT_SUBSTITUTE|ENT_XHTML, "utf-8", false))
	assert.Equal(t, "&apos;", Htmlspecialchars("&apos;", ENT_NOQUOTES|ENT_SUBSTITUTE|ENT_XHTML, "utf-8", false))
	assert.Equal(t, "&lt;", Htmlspecialchars("&lt;", ENT_NOQUOTES|ENT_SUBSTITUTE|ENT_XHTML, "utf-8", false))
	assert.Equal(t, "&gt;", Htmlspecialchars("&gt;", ENT_NOQUOTES|ENT_SUBSTITUTE|ENT_XHTML, "utf-8", false))
	assert.Equal(t, "&#x26;", Htmlspecialchars("&#x26;", ENT_NOQUOTES|ENT_SUBSTITUTE|ENT_XHTML, "utf-8", false))
	assert.Equal(t, "&amp;plus;", Htmlspecialchars("&plus;", ENT_NOQUOTES|ENT_SUBSTITUTE|ENT_XHTML, "utf-8", false))

	// HTML entities tests for ENT_HTML5
	assert.Equal(t, "&amp;", Htmlspecialchars("&amp;", ENT_NOQUOTES|ENT_SUBSTITUTE|ENT_HTML5, "utf-8", false))
	assert.Equal(t, "&quot;", Htmlspecialchars("&quot;", ENT_NOQUOTES|ENT_SUBSTITUTE|ENT_HTML5, "utf-8", false))
	assert.Equal(t, "&apos;", Htmlspecialchars("&apos;", ENT_NOQUOTES|ENT_SUBSTITUTE|ENT_HTML5, "utf-8", false))
	assert.Equal(t, "&lt;", Htmlspecialchars("&lt;", ENT_NOQUOTES|ENT_SUBSTITUTE|ENT_HTML5, "utf-8", false))
	assert.Equal(t, "&gt;", Htmlspecialchars("&gt;", ENT_NOQUOTES|ENT_SUBSTITUTE|ENT_HTML5, "utf-8", false))
	assert.Equal(t, "&#x26;", Htmlspecialchars("&#x26;", ENT_NOQUOTES|ENT_SUBSTITUTE|ENT_HTML5, "utf-8", false))
	assert.Equal(t, "&plus;", Htmlspecialchars("&plus;", ENT_NOQUOTES|ENT_SUBSTITUTE|ENT_HTML5, "utf-8", false))
}
