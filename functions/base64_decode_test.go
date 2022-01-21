package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBase64Decode(t *testing.T) {

	var _err error
	var _ret interface{}
	_ret, _err = Base64Decode("VGhpcyBpcyBhbiBlbmNvZGVkIHN0cmluZw==")
	assert.NoError(t, _err)
	assert.Equal(t, "This is an encoded string", _ret)

	_ret, _err = Base64Decode("RW5jb2RlcyB0aGUgZ2l2ZW4gc3RyaW5nIHdpdGggYmFzZTY0LgoKVGhpcyBlbmNvZGluZyBpcyBkZXNpZ25lZCB0byBtYWtlIGJpbmFyeSBkYXRhIHN1cnZpdmUgdHJhbnNwb3J0IHRocm91Z2ggdHJhbnNwb3J0IGxheWVycyB0aGF0IGFyZSBub3QgOC1iaXQgY2xlYW4sIHN1Y2ggYXMgbWFpbCBib2RpZXMuCgpCYXNlNjQtZW5jb2RlZCBkYXRhIHRha2VzIGFib3V0IDMzJSBtb3JlIHNwYWNlIHRoYW4gdGhlIG9yaWdpbmFsIGRhdGEu")
	assert.NoError(t, _err)
	assert.Equal(t, "Encodes the given string with base64.\n\nThis encoding is designed to make binary data survive transport through transport layers that are not 8-bit clean, such as mail bodies.\n\nBase64-encoded data takes about 33% more space than the original data.", _ret)

	_ret, _err = Base64Decode("5LuK5pel44Gv44GE44GE5aSp5rCX8J+QsQ==")
	assert.NoError(t, _err)
	assert.Equal(t, "今日はいい天気🐱", _ret)

}
