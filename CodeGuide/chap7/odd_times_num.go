package main

// 给定一个整形数组，只有一个数出现了奇数次，其他的数都出现了偶数次，打印这个数
// 时间复杂度为O(N)，额外空间复杂度为O(1)

// 异或操作满足交换律和结合律， a^a=0
func printOddTimesNum(a []int) int {
	ret := 0
	for _, num := range a {
		ret ^= num
	}
	return ret
}

// 如果两个数出现了奇数次，其他数都出现偶数次
// 那么亦或之后，肯定会是a^b != 0
// a^b中至少有1位是1，该位上ab异号
// 据此，再次遍历数组，即可区分出a，b

func printOddTimesNum2(a []int) (int, int) {
	ret := 0
	ret2 := 0
	for _, num := range a {
		ret ^= num
	}
	// ret == a^b
	lowbit := ret & (^ret + 1) // lowbit往右都一样，往左仍然相反
	for _, num := range a {
		if (num & lowbit) != 0 {
			ret2 ^= num
		}
	}
	// ret2 == a || ret2  == b
	return ret2, ret ^ ret2

}

//func main() {
//	a := []int{1, 1, 3, 2, 1, 2, 1}
//	fmt.Println(printOddTimesNum(a))
//	b := []int{1,1,2,2,3,4}
//	fmt.Println(printOddTimesNum2(b))
//}
