package charArray

import (
	"fmt"
	"testing"
)

func Test_reverseStr(t *testing.T) {
	str := reverseStr("abbcdaefabd", 2)
	fmt.Println(str)
}

func TestCharArray(t *testing.T) {
	var arr = [...]int{1, 2, 323}
	fmt.Println(arr)
	var sli = arr[0:]
	fmt.Printf("arr=%p sli=%p len=%d cap=%d\n", &arr, sli, len(sli), cap(sli))
	sli = append(sli, 1)
	fmt.Printf("arr=%p sli=%p len=%d cap=%d\n", &arr, sli, len(sli), cap(sli))
}

func TestType(t *testing.T){
	var runes []rune
	for _, r := range "Hello, world" {
    	runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes) // "['H' 'e' 'l' 'l' 'o' ',' ' ' '世' '界']"
}

func Test_28(t *testing.T){
	strStr("aabaabaafl","aabaaf")
}