package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStripcslashes(t *testing.T) {

	// Some characters are converted in C-like style
	assert.Equal(t, "\n", Stripcslashes("\\n"))   // Line Feed
	assert.Equal(t, "\t", Stripcslashes("\\t"))   // Horizontal Tab
	assert.Equal(t, "\r", Stripcslashes("\\r"))   // Carriage Return
	assert.Equal(t, "\x07", Stripcslashes("\\a")) // Bell
	assert.Equal(t, "\x0b", Stripcslashes("\\v")) // Vertical Tab
	assert.Equal(t, "\x08", Stripcslashes("\\b")) // Backspace
	assert.Equal(t, "\x0c", Stripcslashes("\\f")) // Form Feed
	assert.Equal(t, "\\", Stripcslashes("\\\\"))  // Double backslashes

	// other non-alphanumeric characters with ASCII codes lower than 32 and higher than 126 are converted to octal representation.
	assert.Equal(t, "\x00", Stripcslashes("\\000")) // NUll
	assert.Equal(t, "\x01", Stripcslashes("\\001"))
	assert.Equal(t, "\x02", Stripcslashes("\\002"))
	assert.Equal(t, "\x11", Stripcslashes("\\x11"))
	assert.Equal(t, "\x12", Stripcslashes("\\x12"))

	// Other characters are not converted.
	assert.Equal(t, "あ", Stripcslashes("\\あ"))

	// More than 2 byte characters are handles as like binary data
	assert.Equal(t, "abcあいう", Stripcslashes("ab\\cあいう"))
	assert.Equal(t, "abcあいう", Stripcslashes("ab\\c\\343\\201\\202\\343\\201\\204\\343\\201\\206"))
	assert.Equal(t, "abcあいう😃", Stripcslashes("ab\\c\\343\\201\\202\\343\\201\\204\\343\\201\\206😃"))
	assert.Equal(t, "abcあいう😃", Stripcslashes("ab\\c\\343\\201\\202\\343\\201\\204\\343\\201\\206\\360\\237\\230\\203"))
}
