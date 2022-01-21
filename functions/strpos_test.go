package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrpos(t *testing.T) {
	assert.Equal(t, 0, Strpos("abc", "a"))
	assert.Equal(t, 1, Strpos("abc", "b"))
	assert.Equal(t, 0, Strpos("abc", "ab"))
	assert.Equal(t, 1, Strpos("abc", "bc"))
	assert.Equal(t, 0, Strpos("abc", "abc"))
	assert.Equal(t, 0, Strpos("あいうえお", "あ"))
	assert.Equal(t, 9, Strpos("あいう🐘えお", "🐘え"))
	assert.Equal(t, -1, Strpos("abc", "cb"))
	assert.Equal(t, -1, Strpos("abc", "B"))
	assert.Equal(t, 3, Strpos("abcdef", "de", 0))
	assert.Equal(t, 3, Strpos("abcdef", "de", 3))
	assert.Equal(t, -1, Strpos("abcdef", "de", 4))
}
