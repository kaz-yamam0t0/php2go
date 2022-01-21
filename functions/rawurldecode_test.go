package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRawurldecode(t *testing.T) {

	var _err error
	var _ret string
	_ret, _err = Rawurldecode("%E3%81%82%E3%81%84%E3%81%86%E3%81%88%E3%81%8A")
	assert.NoError(t, _err)
	assert.Equal(t, "あいうえお", _ret)

	// ascii code chars
	_ret, _err = Rawurldecode("%00%01%02%03%04%05%06%07%08%09%0A%0B%0C%0D%0E%0F%10%11%12%13%14%15%16%17%18%19%1A%1B%1C%1D%1E%1F%20%21%22%23%24%25%26%27%28%29%2A%2B%2C-.%2F0123456789%3A%3B%3C%3D%3E%3F%40ABCDEFGHIJKLMNOPQRSTUVWXYZ%5B%5C%5D%5E_%60abcdefghijklmnopqrstuvwxyz%7B%7C%7D%7E")
	assert.NoError(t, _err)
	assert.Equal(t, "\x00\x01\x02\x03\x04\x05\x06\x07\x08\t\n\x0b\x0c\r\x0e\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f !\"#\x24%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_\x60abcdefghijklmnopqrstuvwxyz{|}~", _ret)

	_ret, _err = Rawurldecode("+") // "+" is not converted.
	assert.NoError(t, _err)
	assert.Equal(t, "+", _ret)

	_ret, _err = Rawurldecode("%20")
	assert.NoError(t, _err)
	assert.Equal(t, " ", _ret)

	_ret, _err = Rawurldecode("~")
	assert.NoError(t, _err)
	assert.Equal(t, "~", _ret)

	_ret, _err = Rawurldecode("%7E")
	assert.NoError(t, _err)
	assert.Equal(t, "~", _ret)

}
