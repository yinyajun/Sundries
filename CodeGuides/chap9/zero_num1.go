package main

// N!结果末尾0的个数

// 暴力解法就是计算出N!的结果后，计算末尾0的个数；计算阶乘结果容易上溢

// N!中结果末尾之所以有0，是因为有因子5和因子2，它两相结合就可以凑成10，而因子2是远多于因子5的，只要数一数0~N中有多少个5
// 这种方法效率还是比较低的，每个数计算因子5的个数的复杂度为O(lg i)，以5为底，总时间复杂度为O(N lg N)
func zeroNum1(n int) int {
	if n < 0 {
		return 0
	}
	var cur int
	var res int

	// 只有5的倍数才可能含有因子5
	for i := 5; i <= n; i += 5 {
		cur = i
		// cur中多少个因子5
		for cur%5 == 0 {
			cur /= 5
			res++
		}
	}
	return res
}

// 有个更trick的方法
// 1，2，3，4，5，6，7，8，9，10 ...
// 可以发现，每5个数字提供一个因子5，每25个数字额外提供一个因子5，每125个数字再额外提供一个因子5
// 类似于n转化为5进制后，有几位
func zeroNum2(n int) int {
	if n < 0 {
		return 0
	}
	res := 0
	for n != 0 {
		res += n / 5
		n /= 5
	}
	return res
}

// 给定非负整数N，二进制表达N!的结果，返回最低位的1在那个位置上，0开始计数
// 暴力方法，计算完N！后，通过lowbit方式获取

// 有几个因子2
func zeroNum3(n int) int {

	if n < 0 {
		return 0
	}

	res := 0
	for n != 0 {
		n >>= 1
		res += n
	}
	return res
}

//func main() {
//	//fmt.Println(zeroNum1(19))
//	//fmt.Println(zeroNum1(19))
//
//	fmt.Println(zeroNum3(5))
//}
