package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCtypeXdigit(t *testing.T) {
	assert.Equal(t, true, CtypeXdigit("7f"))
	assert.Equal(t, false, CtypeXdigit("7f.45"))
	assert.Equal(t, false, CtypeXdigit("7f#\x24"))
	assert.Equal(t, false, CtypeXdigit("7f_3c"))
	assert.Equal(t, false, CtypeXdigit("7f-3c"))
	assert.Equal(t, false, CtypeXdigit("7,f3c"))

	// Irregular arguments
	assert.Equal(t, false, CtypeXdigit(""))
}
