package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCtypeCntrl(t *testing.T) {
	assert.Equal(t, true, CtypeCntrl("\x00\x01\x02\x03\x04\x05\n\r\t\x08\x0b"))
	assert.Equal(t, false, CtypeCntrl("\x00\x01\x02\x03 "))
	assert.Equal(t, false, CtypeCntrl("あ")) // All multibyte chars are regarded as not puctuations

	// Irregular arguments
	assert.Equal(t, false, CtypeCntrl(""))
}
