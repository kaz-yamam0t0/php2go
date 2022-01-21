package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrim(t *testing.T) {
	assert.Equal(t, "These are a few words :) ...", Trim("\t\tThese are a few words :) ...  "))
	assert.Equal(t, "These are a few words :)", Trim("\t\tThese are a few words :) ...  ", " \t."))
	assert.Equal(t, "o Wor", Trim("Hello World", "Hdle"))
	assert.Equal(t, "Example string", Trim("\tExample string\n", "\x00..\x1f"))
}
