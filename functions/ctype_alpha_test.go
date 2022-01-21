package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCtypeAlpha(t *testing.T) {
	assert.Equal(t, true, CtypeAlpha("abcDE"))
	assert.Equal(t, false, CtypeAlpha("abcDE123"))
	assert.Equal(t, false, CtypeAlpha("abcDE#\x24"))
	assert.Equal(t, false, CtypeAlpha("abc_dE"))
	assert.Equal(t, false, CtypeAlpha("abC-dE"))

	// Irregular arguments
	assert.Equal(t, false, CtypeAlpha(""))
}
