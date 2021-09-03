package main

// A    B    C    AA    AB    AC       BABC
// 1    2    3    4     5     6        72

// 看起来像k进制，但是注意k进制中每一位的取值范围为[0, k-1]
// 而这里每一个位的取值范围为[1, k]

// 72 = 3^3 * 2 + 3^2 * 2 + 3^1 * 0 + 3^0 * 0 = (2200)3
// 72 = 3^3 * 2 + 3^2 * 1 + 3^1 * 2 + 3^0 * 3 = (2123)

// 可以发现这里和k进制还是有区别

func getKCharFromNum(num int, k int) [32]int {
	res := [32]int{}
	idx := len(res) - 1
	for num != 0 {
		res[idx] = num % k
		num /= k
		idx--
	}
	return res
}

func getCharFromNum(num, k int) [32]int {
	// 确定位数：1,1,1,1
	cur := 1
	length := 0
	sum := 0
	for n := num; n >= cur; {
		length++
		sum += cur
		n -= cur
		cur *= k
	}

	// num - (1111)k，减去这些数，保证每一位至少有1
	// 剩余的数用k进制表示，每一位在[0, k-1]
	// 两者相加起来，正好每位都在[1, k]
	res := getKCharFromNum(num-sum, k)
	for i := len(res) - 1; i >= len(res)-length; i-- {
		res[i] += 1
	}
	return res
}

func getNumFromKchar(num [32]int, k int) int {
	base := 1
	res := 0
	for i := 31; i >= 0; i-- {
		res += num[i] * base
		base *= k
	}
	return res
}

func getNumFromKChar2(num [32]int, k int) int {
	var res int
	for i := 0; i < 32; i++ {
		res = res*k + num[i]
	}
	return res
}

func getNumFromChar(num [32]int, k int) int {
	var res int
	for i := 0; i < 32; i++ {
		res = res*k + num[i]
	}
	return res
}

//func main() {
//	fmt.Println(getKCharFromNum(72, 3))
//	fmt.Println(getCharFromNum(72, 3))
//	fmt.Println(getNumFromKchar(getKCharFromNum(72, 3), 3))
//	fmt.Println(getNumFromKChar2(getKCharFromNum(72, 3), 3))
//	fmt.Println(getNumFromChar(getCharFromNum(72, 3), 3))
//}
