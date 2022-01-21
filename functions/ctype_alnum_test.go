package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCtypeAlnum(t *testing.T) {
	assert.Equal(t, true, CtypeAlnum("abcDE123"))
	assert.Equal(t, false, CtypeAlnum("abcDE123#\x24"))
	assert.Equal(t, false, CtypeAlnum("abc_dE123"))
	assert.Equal(t, false, CtypeAlnum("abc-dE123"))

	// Irregular arguments
	assert.Equal(t, false, CtypeAlnum(""))
}
