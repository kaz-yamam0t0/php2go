package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCtypeUpper(t *testing.T) {
	assert.Equal(t, true, CtypeUpper("ABCDE"))
	assert.Equal(t, false, CtypeUpper("ABcDE"))
	assert.Equal(t, false, CtypeUpper("ABCDE#\x24"))
	assert.Equal(t, false, CtypeUpper("ABC_DE"))
	assert.Equal(t, false, CtypeUpper("ABC-DE"))

	// Irregular arguments
	assert.Equal(t, false, CtypeUpper(""))
}
