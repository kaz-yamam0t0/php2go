package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCtypePunct(t *testing.T) {
	assert.Equal(t, true, CtypePunct("!\"#\x24%&'()*+,-./:;<=>?@[]^_\x60{|}~"))
	assert.Equal(t, false, CtypePunct("\n!@\x60 \n"))
	assert.Equal(t, false, CtypePunct("あ")) // All multibyte chars are regarded as not puctuations

	// Irregular arguments
	assert.Equal(t, false, CtypePunct(""))
}
