package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBase64Encode(t *testing.T) {
	assert.Equal(t, "VGhpcyBpcyBhbiBlbmNvZGVkIHN0cmluZw==", Base64Encode("This is an encoded string"))
	assert.Equal(t, "RW5jb2RlcyB0aGUgZ2l2ZW4gc3RyaW5nIHdpdGggYmFzZTY0LgoKVGhpcyBlbmNvZGluZyBpcyBkZXNpZ25lZCB0byBtYWtlIGJpbmFyeSBkYXRhIHN1cnZpdmUgdHJhbnNwb3J0IHRocm91Z2ggdHJhbnNwb3J0IGxheWVycyB0aGF0IGFyZSBub3QgOC1iaXQgY2xlYW4sIHN1Y2ggYXMgbWFpbCBib2RpZXMuCgpCYXNlNjQtZW5jb2RlZCBkYXRhIHRha2VzIGFib3V0IDMzJSBtb3JlIHNwYWNlIHRoYW4gdGhlIG9yaWdpbmFsIGRhdGEu", Base64Encode("Encodes the given string with base64.\n\nThis encoding is designed to make binary data survive transport through transport layers that are not 8-bit clean, such as mail bodies.\n\nBase64-encoded data takes about 33% more space than the original data."))
	assert.Equal(t, "5LuK5pel44Gv44GE44GE5aSp5rCX8J+QsQ==", Base64Encode("今日はいい天気🐱"))
}
