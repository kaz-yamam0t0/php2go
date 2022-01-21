package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrContains(t *testing.T) {
	assert.Equal(t, true, StrContains("abc", "a"))
	assert.Equal(t, true, StrContains("abc", "b"))
	assert.Equal(t, true, StrContains("abc", "ab"))
	assert.Equal(t, true, StrContains("abc", "bc"))
	assert.Equal(t, true, StrContains("abc", "abc"))
	assert.Equal(t, true, StrContains("あいうえお", "あ"))
	assert.Equal(t, true, StrContains("あいう🐘えお", "🐘え"))
	assert.Equal(t, false, StrContains("abc", "cb"))
	assert.Equal(t, false, StrContains("abc", "B"))
}
