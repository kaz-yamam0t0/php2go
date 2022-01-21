package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCtypeGraph(t *testing.T) {
	assert.Equal(t, true, CtypeGraph("abc!%de"))
	assert.Equal(t, false, CtypeGraph("abc!%de "))
	assert.Equal(t, false, CtypeGraph("\nab!%c \n"))
	assert.Equal(t, false, CtypeGraph("あ")) // All multibyte chars are regarded as not printable

	// Irregular arguments
	assert.Equal(t, false, CtypeGraph(""))
}
