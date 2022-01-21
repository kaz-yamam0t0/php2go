package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCtypePrint(t *testing.T) {
	assert.Equal(t, true, CtypePrint("abcde "))
	assert.Equal(t, false, CtypePrint("\nabc \n"))
	assert.Equal(t, false, CtypePrint("あ")) // All multibyte chars are not regarded as printable

	// Irregular arguments
	assert.Equal(t, false, CtypePrint(""))
}
