package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrlen(t *testing.T) {
	assert.Equal(t, 6, Strlen("abcdef"))
	assert.Equal(t, 7, Strlen(" ab cd "))
	assert.Equal(t, 0, Strlen(""))
}
