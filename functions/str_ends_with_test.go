package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrEndsWith(t *testing.T) {
	assert.Equal(t, true, StrEndsWith("abc", "c"))
	assert.Equal(t, true, StrEndsWith("abc", "abc"))
	assert.Equal(t, false, StrEndsWith("abc", "b"))
	assert.Equal(t, true, StrEndsWith("abc", ""))
	assert.Equal(t, false, StrEndsWith("abc", "C"))
}
