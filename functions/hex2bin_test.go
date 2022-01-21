package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHex2bin(t *testing.T) {

	var _err error
	var _ret interface{}
	_ret, _err = Hex2bin("6578616d706c65206865782064617461")
	assert.NoError(t, _err)
	assert.Equal(t, "example hex data", _ret)

	_ret, _err = Hex2bin("e38182f09f9098")
	assert.NoError(t, _err)
	assert.Equal(t, "あ🐘", _ret)

}
