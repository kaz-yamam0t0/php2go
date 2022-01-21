package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrIreplace(t *testing.T) {
	assert.Equal(t, "<body text='black'>", StrIreplace("%body%", "black", "<body text='%BODY%'>"))
	assert.Equal(t, "You should eat pizza.", StrIreplace("fruits", "pizza", "You should eat Fruits."))
}
