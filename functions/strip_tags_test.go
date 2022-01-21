package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStripTags(t *testing.T) {
	assert.Equal(t, "sample text with tags", StripTags("<b>sample</b> text with <div>tags</div>"))
	assert.Equal(t, "<p>Test paragraph.</p> <a href=\"#fragment\">Other text</a>", StripTags("<p>Test paragraph.</p><!-- Comment --> <a href=\"#fragment\">Other text</a>", "<p><a>"))
	assert.Equal(t, "<p>Test paragraph.</p> <a href=\"#fragment\">Other text</a>", StripTags("<p>Test paragraph.</p><!-- Comment --> <a href=\"#fragment\">Other text</a>", []string{"p", "a"}))
	assert.Equal(t, "<p>Test paragraph.</p>", StripTags("<p>Test paragraph.</p>", map[string]string{"q": "p"})) // <p> is allowed.
	assert.Equal(t, "Test paragraph.", StripTags("<p>Test paragraph.</p>", map[string]string{"p": "q"}))        // <q> is allowed.
}
