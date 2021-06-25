package main

import "fmt"

// 不用任何比较找出两个数中较大的数
// 模拟一个指示函数，I(a>b)*a + I(a<b)*b = I(a-b>0) * a + I(a-b<0) * b
// 这里a-b的符号位可以辅助判断
// a-b>0, sign(a-b)=1
// a-b<0, sign(a-b)=0
// 由于使用a-b，有可能造成上溢，进而导致返回结果不正确

func getMax1(a, b int32) int32 {
	return sign2(a-b)*a + sign2(b-a)*b
}

// 阶跃函数，负数为0，非负数为1（直接判断符号位）
func sign(a int32) int32 {
	return ^(a >> 31) & 1 // 注意这里一定要&1，因为a>>31后还是int32，仍然带有符号位
}

func sign2(a int32) int32 {
	return int32((uint32(a) >> 31) ^ 1) // 通过uint32，先转为无符号整数，然后在亦或，此时再转为int32，注意转为int32时，一定要注意是否越界
}

// 需要考虑a,b的符号
// 首先判断ab是否异号，同号的话，必然不会溢出
// diff = sign(a) ^ sign(b)
// diff ==1 ,异号
// 1. a>=0, b<0, sign(a)=1, sign(b)=0
// 2. a<0, b>=0, sign(a)=0, sign(b)=1

// diff == 0，同号，此时不会溢出
// 1. a-b >=0 , sign(a-b) = 1
// 1. a-b <0 , sign(a-b) = 0

// I(diff==1) * (sA * a + SB * b) + I(diff==0) * (sC * a + sC * b)

func flip(a int32) int32 {
	return a ^ 1
}

// 注意，不要直接使用非来做flip，因为是补码存储，还有符号位
func getMax2(a, b int32) int32 {
	diff := sign2(a) ^ sign2(b)
	sa := sign2(a)
	sc := sign2(a - b)

	coefA := diff*sa + (diff^1)*sc
	fmt.Println(diff, sa, diff^1, sc)
	return coefA*a + (coefA^1)*b
}

//func main() {
//	fmt.Println(getMax1(3, -543))
//	fmt.Println(getMax2(3, -543))
//}
