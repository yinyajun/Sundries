package main

// 找到最高位和最低位

func IsPalindromeNumber(n int) bool {
	base := 1

	for n >= 10*base {
		base *= 10
	}

	for n != 0 {
		high := n / base
		low := n % 10

		if high != low {
			return false
		}
		n = (n % base) / 10
		base /= 100

	}
	return true
}

//func main() {
//	fmt.Println(IsPalindromeNumber(131))
//}
