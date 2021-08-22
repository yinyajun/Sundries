package main

func product(arr []int) []int {
	if len(arr) < 2 {
		panic("invalid arr")
	}

	all := 1
	count := 0

	for _, val := range arr {
		if val != 0 {
			all *= val
		} else {
			count++
		}
	}

	res := make([]int, len(arr))
	// 没有出现0
	if count == 0 {
		for i := 0; i < len(arr); i++ {
			res[i] = all / arr[i]
		}
	}
	// 只有一个0
	if count == 1 {
		for i := 0; i < len(arr); i++ {
			if arr[i] == 0 {
				res[i] = all
			}
		}
	}
	// 有多个0，其余乘积必为0
	// count > 1, res[i] == 0
	return res
}

// 通过预处理，left：从左到右累乘； right：从右到左累乘。
// 有了累乘数组，res[i]=left[i-1] * right[i+1]
// 但是这两个数组会额外占用空间，先从左到右，将left[i]保存在res[i]中，然后从右到左，计算right[i-1]，然后将结果保存在res总
func product2(arr []int) []int {
	res := make([]int, len(arr))

	prod := 1

	// left
	for i, val := range arr {
		prod *= val
		res[i] = prod
	}
	//right
	prod = 1

	for i := len(arr) - 1; i > 0; i-- {
		res[i] = res[i-1] * prod
		prod *= arr[i]
	}

	res[0] = prod

	return res
}

//func main() {
//	arr := []int{2, 3, 1, 4}
//	fmt.Println(product(arr))
//	fmt.Println(product2(arr))
//}
