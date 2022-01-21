package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrReplace(t *testing.T) {
	assert.Equal(t, "<body text='black'>", StrReplace("%body%", "black", "<body text='%body%'>"))
	assert.Equal(t, "You should eat pizza.", StrReplace("fruits", "pizza", "You should eat fruits."))
}
