package main

import (
	"CodeGuide/base/utils"
	"fmt"
)

// !!该方法是错误的，如果数组中是3，4，6，5，遍历到6的时候不知道还有5的存在，没法构成最大子数组
// 当前状态可能和过去相关
func maxLengthConsecutiveArray(arr []int) int {
	m := make(map[int]int)
	var ans int
	for _, val := range arr {
		if _, ok := m[val-1]; ok {
			m[val] = m[val-1] + 1
		} else {
			m[val] = 1
		}
		fmt.Println(val, m[val])
		if m[val] > ans {
			ans = m[val]
		}
	}
	return ans
}

// 时间复杂度：O(N)  预处理O(N)，双重循环其实是O(N)的复杂度
// 因为外循环仅从连续子数组开头，而内循环继续数连续子数组，那么双重循环将会将数组中所有连续子数组遍历一遍，所以时间复杂度为O(N)
// 当然还有个简单方式，在外循环和内循环中分别打印一下访问元素，不难看出，所有元素仅被访问一遍。（参考单调栈，每次更新，都要删除不符合条件的元素，也是双重循环，复杂度为为O(N)）
func maxLengthConsecutiveArray2(arr []int) int {
	m := make(map[int]struct{})
	var ans int
	// preprocess
	for _, val := range arr {
		m[val] = struct{}{}
	}

	for val, _ := range m {
		if _, ok := m[val-1]; !ok {
			length := 1
			// 其他语言可以很容易实现
			//while m.contains(++val){
			//	length++
			//}

			// 实现1
			//val++
			//for _, exist := m[val]; exist; {
			//	length++
			//	val++
			//	_, exist = m[val]
			//}

			// 实现2
			for {
				val++
				if _, exist := m[val]; !exist {
					break
				} else {
					length++
				}
			}

			// 错误：这里的ok作为循环变量的判断条件并没有更新，一直是初始化的值
			//for _, ok :=m[val+1];ok ; val++{
			//	fmt.Println(val+1, m)
			//	time.Sleep(1*time.Second)
			//	length ++
			//}
			ans = utils.MaxInt(ans, length)
		}
	}
	return ans
}

//func main() {
//	arr := []int{10, 1, 3, 4, 7, 6, 20, 5, 13, 23, 14}
//	fmt.Println(maxLengthConsecutiveArray2(arr))
//}
