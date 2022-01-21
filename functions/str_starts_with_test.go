package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrStartsWith(t *testing.T) {
	assert.Equal(t, true, StrStartsWith("abc", "a"))
	assert.Equal(t, true, StrStartsWith("abc", "abc"))
	assert.Equal(t, true, StrStartsWith("abc", ""))
	assert.Equal(t, false, StrStartsWith("abc", "A"))
}
