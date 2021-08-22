package main

// 无需整形数组中，找到未出现的最小正整数

// 理想情况下，数组中应该是[1, N]这N个数
// [1, l]已经确定的正整数，[r, N-1]无效元素;l=0, r=N
// arr[l] == l+1, 对应的元素在对应的位置上，直接扩展左边区间，[1, l+1]为包含的正整数
// arr[l] <= l，已经包含正整数为[1,l]，那么[l+1,r]内数字就少一个，右边区间变为[l, r-1]，并将有效数组的最后arr[r-1]放到当前位置
// arr[l] > r, 同上
// arr[l] \in [l+1, r], 此时arr[l]应该放到arr[l]-1这个位置上，交换（有个特殊情况，交换后元素不变，也就是说，出现那个位置上的元素和arr[l]重复，这个情况需要额外处理下）

// 时间复杂度为O(N)
func missNum(arr []int) int {
	l, r := 0, len(arr)

	for l < r {
		if arr[l] == l+1 {
			l++
		} else if arr[l] <= l || arr[l] > r {
			r--
			arr[l] = arr[r]
		} else {
			if arr[arr[l]-1] == arr[l] { //两者保留一个即可，保留位置正确的那个
				r--
				arr[l] = arr[r]
			} else {
				arr[l], arr[arr[l]-1] = arr[arr[l]-1], arr[l]
			}
		}
	}
	return l + 1
}

//func main() {
//	arr := []int{2, 1, 3, 4}
//	fmt.Println(missNum(arr))
//}
