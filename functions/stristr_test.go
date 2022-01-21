package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStristr(t *testing.T) {

	var _err error
	var _ret interface{}
	_ret, _err = Stristr("NAME@EXAMPLE.COM", "e")
	assert.NoError(t, _err)
	assert.Equal(t, "E@EXAMPLE.COM", _ret)

	_ret, _err = Stristr("NAME@EXAMPLE.COM", "e", true)
	assert.NoError(t, _err)
	assert.Equal(t, "NAM", _ret)

	_ret, _err = Stristr("NAME@EXAMPLE.COM", "name")
	assert.NoError(t, _err)
	assert.Equal(t, "NAME@EXAMPLE.COM", _ret)

	_ret, _err = Stristr("NAME@EXAMPLE.COM", "name", true)
	assert.NoError(t, _err)
	assert.Equal(t, "", _ret)

	_ret, _err = Stristr("NAME@EXAMPLE.COM", "@example.com")
	assert.NoError(t, _err)
	assert.Equal(t, "@EXAMPLE.COM", _ret)

	_ret, _err = Stristr("NAME@EXAMPLE.COM", "@example.com", true)
	assert.NoError(t, _err)
	assert.Equal(t, "NAME", _ret)

}
