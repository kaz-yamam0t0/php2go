package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRtrim(t *testing.T) {
	assert.Equal(t, "\t\tThese are a few words :) ...", Rtrim("\t\tThese are a few words :) ...  "))
	assert.Equal(t, "\t\tThese are a few words :)", Rtrim("\t\tThese are a few words :) ...  ", " \t."))
	assert.Equal(t, "Hello Wor", Rtrim("Hello World", "Hdle"))
	assert.Equal(t, "\tExample string", Rtrim("\tExample string\n", "\x00..\x1f"))
}
