package main

import (
	"math"
)

// 只用位运算不用算术运算实现整数的加减乘除运算
// 不用考虑溢出情况

// 加法：
// 不考虑进位的加法：按位亦或， a^b
// 进位：该位相加均为1，才需要进位，(a&b)<<1

// 那么直接将 a^b 和 进位 相加就可以了吗？
// 相加后仍然会产生进位
// 所以需要不断相加，直到 进位为0

func add(a, b int32) int32 {
	ret := a ^ b
	carry := (a & b) << 1
	for carry != 0 {
		ret, carry = ret^carry, (ret&carry)<<1 // 注意如果分两行写，此时ret可能已经重新更新了
	}
	return ret
}

func add2(a, b int32) int32 {
	ret := a
	for b != 0 {
		ret = a ^ b // 之所以多用一个变量，因为不能先污染a变量
		b = (a & b) << 1
		a = ret
	}
	return ret
}

// 直接复用形参变量的内存地址
// 这里可以将 [a,b]视为列向量，每一步都是一个非线性变换，得到新的[a,b]
// 类似的有斐波那契额数列的dp解法
func add3(a, b int32) int32 {
	for b != 0 {
		a, b = a^b, (a&b)<<1
	}
	return a
}

// a-b = a + (-b)
// 也就是说，只要找到b的相反数，就可以完成减法
func minus(a, b int32) int32 {
	return add3(a, negative(b))
}

// 负数的补码：正数的原码，取反+1
func negative(a int32) int32 {
	return add3(^a, 1)
}

// 乘法运算: 将乘法改为加法
// a*b= a * (2^0 * b0 + 2^1 * b1 + 2^2 * b2 + ...)
func multiply(a, b int32) int32 {
	var res int32
	for b != 0 {
		if b&1 > 0 { // b在该位上为1
			res += a
		}
		a <<= 1
		b >>= 1
	}
	return res
}

// 除法运算：将除法改为减法
// 如果 b * res = a ( a / b = res)
// a = b * (2^0 * res0 + 2^1 * res1 + ... + 2^31 * res31)
// 只要分别找到res上各位的值即可
// 注意，如果b*2^31>a，那么a中不可能包含b*2^31，所以res31=0

// 类似于将十进制转为二进制
// 先找到a能包含的最大部分，然后a-最大部分，然后在剩余的a中找到次大部分，并依此找下去（两者皆为非负数）
// 如果有一个为负数，先转为正数计算，然后添加上符号

func _divide(a, b int32) int32 {
	x, y := a, b
	if a < 0 {
		x = negative(a)
	}
	if b < 0 {
		y = negative(b)
	}

	var i, res int32
	for i = 31; i >= 0; i = minus(i, 1) {
		// a > (b * 2^i) ?  => a / 2^i  > b
		if (x >> i) >= y { // 代表res_i有值
			res |= 1 << i
			x = minus(x, y<<i) // a - b* 2^i * res_i
		}
	}

	if (a > 0 && b < 0) || (a < 0 && b > 0) {
		return negative(res)
	}
	return res
}

// 除法中，有个步骤，将负数转为正数
// 这里有个巨坑：正数的表达范围 比 负数的表达范围 小1
// 如果a,b中有一个正好是负数中的最小值，是没法直接转为对应的正数的
// 1. 如果a，b均不是最小值，上面正常除法
// 2. 如果a，b均是最小值，a/b == 1
// 3. 仅b是最小值，a/b == 0
// 4. 仅a是最小值，此时可以通过一点修正方法

// a/b =》 (a + 1 -1) / b =》 (a+1)/b + (-1)/b
// c = (a+1)/b , -1/b = (a - c * b) /b
// res = c + (a - c* b) /b

func divide(a, b int32) int32 {
	if b == 0 {
		panic("divisor is zero")
	}
	if a == math.MinInt32 && b == math.MinInt32 {
		return 1
	}
	if b == math.MinInt32 {
		return 0
	}
	if a == math.MinInt32 {
		c := _divide(add(a, 1), b)
		return add(c, _divide(minus(a, multiply(c, b)), b))
	}
	return _divide(a, b)
}

//func main() {
//	a, b := int32(-27), int32(9)
//	fmt.Println(add3(a, b))
//	fmt.Println(minus(a, b))
//	fmt.Println(multiply(a, b))
//	fmt.Println(divide(a, b))
//}
