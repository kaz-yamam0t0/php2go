package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrrchr(t *testing.T) {

	var _err error
	var _ret interface{}
	_ret, _err = Strrchr("/usr/bin:/bin:/usr/local/bin", ":")
	assert.NoError(t, _err)
	assert.Equal(t, ":/usr/local/bin", _ret)

	_ret, _err = Strrchr("aaa\nbbbb\ncccccc", "\n")
	assert.NoError(t, _err)
	assert.Equal(t, "\ncccccc", _ret)

	_ret, _err = Strrchr("/www/public_html/index.html", "/")
	assert.NoError(t, _err)
	assert.Equal(t, "/index.html", _ret)

}
