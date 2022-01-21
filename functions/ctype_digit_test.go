package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCtypeDigit(t *testing.T) {
	assert.Equal(t, true, CtypeDigit("12345"))
	assert.Equal(t, false, CtypeDigit("123.45"))
	assert.Equal(t, false, CtypeDigit("12345#\x24"))
	assert.Equal(t, false, CtypeDigit("123_45"))
	assert.Equal(t, false, CtypeDigit("123-45"))
	assert.Equal(t, false, CtypeDigit("12,345"))

	// Irregular arguments
	assert.Equal(t, false, CtypeDigit(""))
}
