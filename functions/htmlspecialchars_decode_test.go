package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHtmlspecialcharsDecode(t *testing.T) {

	// Tests for quotes
	assert.Equal(t, "<a href='test'>\"Test\"</a>", HtmlspecialcharsDecode("&lt;a href=&#039;test&#039;&gt;&quot;Test&quot;&lt;/a&gt;", ENT_QUOTES))
	assert.Equal(t, "<a href=&#039;test&#039;>\"Test\"</a>", HtmlspecialcharsDecode("&lt;a href=&#039;test&#039;&gt;&quot;Test&quot;&lt;/a&gt;", ENT_COMPAT))
	assert.Equal(t, "<a href=&#039;test&#039;>&quot;Test&quot;</a>", HtmlspecialcharsDecode("&lt;a href=&#039;test&#039;&gt;&quot;Test&quot;&lt;/a&gt;", ENT_NOQUOTES))

	// (Unlike `htmlspecialchars`, )`htmlspecialchars_decode` doesn't convert invalid character sequences.
	assert.Equal(t, "\x81\x81<>", HtmlspecialcharsDecode("\x81\x81&lt;&gt;", ENT_QUOTES))
	assert.Equal(t, "\x81\x81<>", HtmlspecialcharsDecode("\x81\x81&lt;&gt;", ENT_SUBSTITUTE))
}
