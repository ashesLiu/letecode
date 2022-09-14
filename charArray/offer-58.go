package charArray

func reverseLeftWords(s string, n int) string {

    // reverse [left,right)
    reverse := func(str []byte, left, right int){
        l,r := left, right-1
        for l<r{
            str[l], str[r] = str[r], str[l]
            l++
            r--
        }
    }

    str := []byte(s)
    reverse(str, 0, len(str))
    reverse(str, 0, len(str)-n)
    reverse(str, len(str)-n, len(str))

    return string(str)
}