package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddcslashes(t *testing.T) {

	// Some characters are converted in C-like style
	assert.Equal(t, "\\n", Addcslashes("\n", "\x00..\xff"))   // Line Feed
	assert.Equal(t, "\\t", Addcslashes("\t", "\x00..\xff"))   // Horizontal Tab
	assert.Equal(t, "\\r", Addcslashes("\r", "\x00..\xff"))   // Carriage Return
	assert.Equal(t, "\\a", Addcslashes("\x07", "\x00..\xff")) // Bell
	assert.Equal(t, "\\v", Addcslashes("\x0b", "\x00..\xff")) // Vertical Tab
	assert.Equal(t, "\\b", Addcslashes("\x08", "\x00..\xff")) // Backspace
	assert.Equal(t, "\\f", Addcslashes("\x0c", "\x00..\xff")) // Form Feed

	// other non-alphanumeric characters with ASCII codes lower than 32 and higher than 126 are converted to octal representation.
	assert.Equal(t, "\\000", Addcslashes("\x00", "\x00..\xff")) // NUll
	assert.Equal(t, "\\001", Addcslashes("\x01", "\x00..\xff"))
	assert.Equal(t, "\\002", Addcslashes("\x02", "\x00..\xff"))

	// alphanumeric characters are not converted.
	assert.Equal(t, "\\a", Addcslashes("a", "\x00..\xff"))
	assert.Equal(t, "\\z", Addcslashes("z", "\x00..\xff"))
	assert.Equal(t, "\\A", Addcslashes("A", "\x00..\xff"))
	assert.Equal(t, "\\Z", Addcslashes("Z", "\x00..\xff"))

	// [\]^_` are included between "A" and "z"
	assert.Equal(t, "\\f\\o\\o\\[ \\]", Addcslashes("foo[ ]", "A..z"))
	assert.Equal(t, "\\f\\o\\o[ ]", Addcslashes("foo[ ]", "a..zA..Z"))

	// No range is constructed when first char > last char, but this usage will actually show warning.
	//assert.Equal(t, "\\zoo['\\.']", Addcslashes("zoo['.']","z..A"))

	// More than 2 byte characters are handles as like binary data
	assert.Equal(t, "\\a\\b\\cあいう", Addcslashes("abcあいう", "abc"))
	assert.Equal(t, "\\a\\b\\c\\343\\201\\202\\343\\201\\204\\343\\201\\206", Addcslashes("abcあいう", "abcあいう"))
	assert.Equal(t, "\\a\\b\\c\\343\\201\\202\\343\\201\\204\\343\\201\\206", Addcslashes("abcあいう", "abcあ..う"))
	assert.Equal(t, "\\a\\b\\c\\343\\201\\202\\343\\201\x84\\343\\201\\206", Addcslashes("abcあいう", "abcあう"))
	assert.Equal(t, "\\a\\b\\c\\343\\201\\202\\343\\201\x84\\343\\201\\206😃", Addcslashes("abcあいう😃", "abcあう"))
	assert.Equal(t, "\\a\\b\\c\\343\\201\\202\\343\\201\x84\\343\\201\\206\\360\\237\\230\\203", Addcslashes("abcあいう😃", "abc😃あう"))
}
