package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStripslashes(t *testing.T) {
	assert.Equal(t, "", Stripslashes("\\"))
	assert.Equal(t, " ", Stripslashes("\\ "))
	assert.Equal(t, "Is your name O'reilly?", Stripslashes("Is your name O'reilly?"))
	assert.Equal(t, "Is your name O\\'reilly?", Stripslashes("Is your name O\\\\'reilly?"))
}
