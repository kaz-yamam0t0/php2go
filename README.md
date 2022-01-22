# php2go
Golang equivalent to PHP functions

## Usage

```go
import (
	"fmt"
	php2go "github.com/kaz-yamam0t0/php2go/functions"
)

func main() {
	fmt.Printf("now: %s\n", php2go.Date("Y-m-d H:i:s"))
}
```

## Supported Functions

#### strings

* addcslashes
* addslashes
* bin2hex
* chr
* chunk_split
* count_chars
* explode
* hex2bin
* htmlentities
* htmlspecialchars
* htmlspecialchars_decode
* implode
* lcfirst
* ltrim
* md5
* nl2br
* number_format
* ord
* rtrim
* sprintf
* str_contains
* str_ends_with
* str_ireplace
* str_pad
* str_repeat
* str_replace
* str_starts_with
* strcasecmp
* strcmp
* strip_tags
* stripcslashes
* stripos
* stripslashes
* stristr
* strlen
* strncasecmp
* strncmp
* strpos
* strrchr
* strstr
* strtolower
* strtoupper
* substr
* trim
* ucfirst
* ucwords
* wordwrap

#### filesystem

* basename
* dirname

#### array

* array_keys

#### ctype

* ctype_alnum
* ctype_alpha
* ctype_cntrl
* ctype_digit
* ctype_graph
* ctype_lower
* ctype_print
* ctype_punct
* ctype_space
* ctype_upper
* ctype_xdigit

#### datetime

* date
* strtotime

#### url

* base64_decode
* base64_encode
* rawurldecode
* rawurlencode
* urldecode
* urlencode