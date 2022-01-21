package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStripos(t *testing.T) {
	assert.Equal(t, 0, Stripos("abc", "a"))
	assert.Equal(t, 1, Stripos("abc", "b"))
	assert.Equal(t, 0, Stripos("abc", "ab"))
	assert.Equal(t, 1, Stripos("abc", "bc"))
	assert.Equal(t, 0, Stripos("abc", "abc"))
	assert.Equal(t, 0, Stripos("abc", "A"))
	assert.Equal(t, 1, Stripos("abc", "B"))
	assert.Equal(t, 0, Stripos("abc", "AB"))
	assert.Equal(t, 1, Stripos("abc", "BC"))
	assert.Equal(t, 1, Stripos("abc", "bC"))
	assert.Equal(t, 0, Stripos("abc", "AbC"))
	assert.Equal(t, 0, Stripos("あいうえお", "あ"))
	assert.Equal(t, 9, Stripos("あいう🐘えお", "🐘え"))
	assert.Equal(t, -1, Stripos("abc", "cb"))
	assert.Equal(t, 3, Stripos("abcdef", "De", 0))
	assert.Equal(t, 3, Stripos("abcdef", "dEF", 3))
	assert.Equal(t, -1, Stripos("abcdef", "DeF", 4))
}
