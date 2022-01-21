package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrd(t *testing.T) {
	assert.Equal(t, 10, Ord("\n"))
	assert.Equal(t, 97, Ord("a"))
	assert.Equal(t, 227, Ord("あ")) // code of the first char
	assert.Equal(t, 240, Ord("🐘")) // code of the first char
}
