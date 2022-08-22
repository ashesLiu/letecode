package contest

func largestPalindromic(num string) string {
    digits := make([]int, 10)
	for i:= 0;i< len(num);i++{
		digits[num[i]-'0']++
	}
	string
	for i:= 9;i>=0;i--{
		// 偶数
		if digits[i]!=0&&digits[i]%2==0{

		}
	}
}