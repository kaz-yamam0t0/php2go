package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLtrim(t *testing.T) {
	assert.Equal(t, "These are a few words :) ...  ", Ltrim("\t\tThese are a few words :) ...  "))
	assert.Equal(t, "These are a few words :) ...  ", Ltrim("\t\tThese are a few words :) ...  ", " \t."))
	assert.Equal(t, "o World", Ltrim("Hello World", "Hdle"))
	assert.Equal(t, "Example string\n", Ltrim("\tExample string\n", "\x00..\x1f"))
}
