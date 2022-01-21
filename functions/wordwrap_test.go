package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWordwrap(t *testing.T) {

	var _err error
	var _ret string
	_ret, _err = Wordwrap("The quick brown fox jumped over the lazy dog.", 20, "<br />\n")
	assert.NoError(t, _err)
	assert.Equal(t, "The quick brown fox<br />\njumped over the lazy<br />\ndog.", _ret)

	_ret, _err = Wordwrap("A very long woooooooooooord.", 8, "\n", true)
	assert.NoError(t, _err)
	assert.Equal(t, "A very\nlong\nwooooooo\nooooord.", _ret)

	_ret, _err = Wordwrap("A very long wooooooooooooooooord.", 8, "\n", true)
	assert.NoError(t, _err)
	assert.Equal(t, "A very\nlong\nwooooooo\noooooooo\noord.", _ret)

	_ret, _err = Wordwrap("A very long woooooooooooord.", 8, "\n", false)
	assert.NoError(t, _err)
	assert.Equal(t, "A very\nlong\nwoooooooooooord.", _ret)

}
