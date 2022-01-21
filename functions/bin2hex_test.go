package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBin2hex(t *testing.T) {
	assert.Equal(t, "000a0d", Bin2hex("\x00\n\r"))
	assert.Equal(t, "e38182f09f9098", Bin2hex("あ🐘"))
}
