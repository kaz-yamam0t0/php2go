package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrstr(t *testing.T) {

	var _err error
	var _ret interface{}
	_ret, _err = Strstr("name@example.com", "@")
	assert.NoError(t, _err)
	assert.Equal(t, "@example.com", _ret)

	_ret, _err = Strstr("name@example.com", "@", true)
	assert.NoError(t, _err)
	assert.Equal(t, "name", _ret)

	_ret, _err = Strstr("name@example.com", "name")
	assert.NoError(t, _err)
	assert.Equal(t, "name@example.com", _ret)

	_ret, _err = Strstr("name@example.com", "name", true)
	assert.NoError(t, _err)
	assert.Equal(t, "", _ret)

	_ret, _err = Strstr("name@example.com", "@example.com")
	assert.NoError(t, _err)
	assert.Equal(t, "@example.com", _ret)

	_ret, _err = Strstr("name@example.com", "@example.com", true)
	assert.NoError(t, _err)
	assert.Equal(t, "name", _ret)

}
