# php2go
Golang equivalent to PHP functions

## Usage

```go
import (
	"fmt"
	php2go "github.com/kaz-yamam0t0/php2go/functions"
)

func main() {
	// 2000-09-10 12:23:34
	fmt.Printf("%s\n", php2go.Date("Y-m-d H:i:s", php2go.Strtotime("10 September 2000 12:23:34")))
}
```

## Supported Functions

#### strings

* Addcslashes
* Addslashes
* Bin2hex
* Chr
* ChunkSplit
* CountChars
* Explode
* Hex2bin
* Htmlentities
* Htmlspecialchars
* HtmlspecialcharsDecode
* Implode
* Lcfirst
* Ltrim
* Md5
* Nl2br
* NumberFormat
* Ord
* Rtrim
* Sprintf
* StrContains
* StrEndsWith
* StrIreplace
* StrPad
* StrRepeat
* StrReplace
* StrStartsWith
* Strcasecmp
* Strcmp
* StripTags
* Stripcslashes
* Stripos
* Stripslashes
* Stristr
* Strlen
* Strncasecmp
* Strncmp
* Strpos
* Strrchr
* Strstr
* Strtolower
* Strtoupper
* Substr
* Trim
* Ucfirst
* Ucwords
* Wordwrap

#### filesystem

* Basename
* Dirname

#### array

* ArrayKeys

#### ctype

* CtypeAlnum
* CtypeAlpha
* CtypeCntrl
* CtypeDigit
* CtypeGraph
* CtypeLower
* CtypePrint
* CtypePunct
* CtypeSpace
* CtypeUpper
* CtypeXdigit

#### datetime

* Date
* Strtotime

#### url

* Base64Decode
* Base64Encode
* Rawurldecode
* Rawurlencode
* Urldecode
* Urlencode