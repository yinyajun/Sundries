package main

// 不断右移，判断最低位是否为1，最坏循环32次
func count1(n int32) int {
	var cnt int32
	for n != 0 {
		cnt += n & 1
		n >>= 1
	}
	return int(cnt)
}

// 如果降低复杂度呢？降为只和1的个数相关
// 从lowbit运算中获取灵感：x-1会将最低位1后面的所有位反号，lowbit是与原数异或，来获取包括lowbit的掩码
// 如果(x-1)&x，最低位的lowbit直接清零，其余不变，那么，不断的清除lowbit直到x为0
// (x-1)&x == x-lowbit(x)
func count2(n int32) int {
	var cnt int32
	for n != 0 {
		n &= n - 1
		cnt++
	}
	return int(cnt)
}

// 从上面推论可以知道，也可以直接利用lowbit运算
func count3(n int32) int {
	var cnt int32
	for n != 0 {
		n -= n & (^n + 1) // n-= n & (-n)
		cnt++
	}
	return int(cnt)
}

// 二进制平行算法 https://blog.csdn.net/erzr_zhang/article/details/55212676
// 这个算法很巧妙，类似于归并过程
// 0. 将每位视为1组，
// 1. 先将相邻的组（组大小为1）相加，形成新组（组大小为2），每组的值为组内1的个数
// 2. 再将相邻的组（组大小为2）相加，形成新组（组大小为4），每组的值为组内1的个数
// 3. 再将相邻的组（组大小为4）相加，形成新组（组大小为8），每组的值为组内1的个数
// ...
// 5. 再将相邻的组（组大小为16）相加，形成新组（组大小为32），每组的值为组内1的个数

// 如何使相邻的组相加，通过间隔的1位掩码
// 0101010101(归并大小为1的组)  00110011(归并大小为2的组)
func count4(n int32) int {
	n = (n & 0x55555555) + ((n >> 1) & 0x55555555)
	n = (n & 0x33333333) + ((n >> 2) & 0x33333333)
	n = (n & 0x0f0f0f0f) + ((n >> 4) & 0x0f0f0f0f)
	n = (n & 0x00ff00ff) + ((n >> 8) & 0x00ff00ff)
	n = (n & 0x0000ffff) + ((n >> 16) & 0x0000ffff)
	return int(n)
}

//func main() {
//	n := int32(56)
//	fmt.Println(count1(n))
//	fmt.Println(count2(n))
//	fmt.Println(count3(n))
//	fmt.Println(count4(n))
//}
