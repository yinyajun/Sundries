package main

import (
	"CodeGuide/base/utils"
	"fmt"
)

// ** 只要实现某一个区间上的等概随机函数，就可以实现任意区间上的随机函数

// ==> 只用rand1to5，扩展到rand1to7
func rand1to5() int {
	return utils.Random.Intn(5) + 1
}

func rand1to7() int {
	num := 0
	for {
		num = (rand1to5()-1)*5 + rand1to5() // rand[0,24]
		if num < 21 {
			break
		}
	} // 多出的[21,24]均匀的摊在了[0,20]上
	return num%7 + 1 // rand[0,6] + 1
}

// ==> rand01p，实现rand1to6
func rand01p() int {
	p := 0.83
	if utils.Random.Float64() < p {
		return 0
	}
	return 1
}

// 用rand01p出现01和10的概率都为p(1-p)，等概率
func rand01() int {
	var num int
	for {
		num = rand01p()
		if num != rand01p() {
			break
		}
	}
	return num
}

func rand0to3() int {
	return rand01()*2 + rand01()
}

func rand0to7() int {
	return rand0to3()*2 + rand01()
}

func rand1to6() int {
	var num int
	for {
		num = rand0to7()
		if num <= 5 {
			break
		}
	}
	return num + 1
}

// --> 给定一个等概率随机产生1~M的随机函数
// 请用rand1toM(m)实现rand1toN
// *** 如果仅用rand1toM，它能实现[0,M), [0, M^2), [0, M^3)...上的随机取值
// [0, M^2):  rand1toM*M + rand1toM
// 也可以这么认为是两位M进制的数，每个位置上去[0,M)的一个值，那么它的取值范围也是[0, M^2)
func rand1toM(m int) int {
	return utils.Random.Intn(m) + 1
}

// M<N的话，可能需要多位，根据N的大小来定位数；方便起见，统一转到M进制下处理，这样位数就是N在M进制下的位数
func rand0toNInMSys(N [32]int, m int) [32]int {
	res := [32]int{}

	var start int
	for N[start] == 0 {
		start++
	}
	// N[start]!=0

	// 此时用rand1toM来创建一个范围比[0,N]大的随机数，大于N的随机数丢弃
	lastEqual := true // 上一位是否相等，相等的话，下一位需要比较

	index := start
	for index < 32 {
		res[index] = rand1toM(m) - 1
		if lastEqual {
			if res[index] > N[index] { // 大于N了，重新来
				index = start
				continue
			} else {
				if res[index] < N[index] {
					lastEqual = false
				}
			}
		}
		index++
	}
	return res
}

func rand1toN(m, n int) int {
	nMsys := getMsystemNum(n-1, m)
	r := rand0toNInMSys(nMsys, m)
	return getNumFromMsystem(r, m) + 1
}

// 长除法: 将value转成m进制数
func getMsystemNum(val, m int) [32]int {
	res := [32]int{}
	index := len(res) - 1 // 从最高位开始
	for val != 0 {
		res[index] = val % m
		index--
		val /= m
	}
	return res
}

// 将m进制数转为十进制：从最低位开始，每次位置上的值乘以base
func getNumFromMsystem(mSysNum [32]int, m int) int {
	base := 1
	var res int
	for i := len(mSysNum) - 1; i >= 0; i-- {
		res += mSysNum[i] * base
		base *= m
	}
	return res
}

// 将m进制转为十进制：从最高位开始
func getNumFromMsystem2(mSysNum [32]int, m int) int {
	var res int
	for i := 0; i < len(mSysNum); i++ {
		res = res*m + mSysNum[i]
	}
	return res
}

func count(r func() int) {
	m := make(map[int]int)
	for i := 0; i < 10000; i++ {
		m[r()] += 1
	}
	fmt.Println(m)
}

func count2(m, n int, r func(int, int) int) {
	mm := make(map[int]int)
	for i := 0; i < 10000; i++ {
		mm[r(m, n)] += 1
	}
	fmt.Println(mm)
}

//func main() {
//	//count2(5, 9 , rand1toN)
//	fmt.Println(getMsystemNum(24, 2))
//	//fmt.Println(getNumFromMsystem2(getMsystemNum(9, 2), 2))
//}
