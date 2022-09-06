package charArray

import (
	"fmt"
	"testing"
)

func Test_reverseStr(t *testing.T) {
	str := reverseStr("abbcdaefabd", 2)
	fmt.Println(str)
}
