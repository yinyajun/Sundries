package main

// 时间复杂度O(N)，空间复杂度O(1)

// 数组中均为正数
func GetMaxLength(a []int, k int) int {
	// [left, right)
	left, right := 0, 0
	sum := 0
	length := 0

	for right <= len(a) { // 合法范围[0, len(a))
		if sum < k {
			if right == len(a) { // right越界
				break
			}
			sum += a[right] // [left, right]
			right++         // [left, right)
		} else if sum > k {
			sum -= a[left]
			left++
		} else { // sum == k
			if right-left > length {
				length = right - left
			}
			sum -= a[left]
			left++
		}
	}
	return length
}

func GetMaxLength2(a []int, k int) int {
	// [left, right]
	left, right := 0, -1
	sum := 0
	length := 0

	for right < len(a) {
		if sum < k {
			right++
			if right == len(a) { // 越界
				break
			}
			sum += a[right]
		} else if sum > k {
			sum -= a[left]
			left++
		} else { // sum == k
			if right-left+1 > length {
				length = right - left + 1
			}
			sum -= a[left]
			left++
		}
	}
	return length
}

//func main() {
//	a := []int{1, 2, 1, 1, 1}
//	fmt.Println(GetMaxLength(a, 3))
//	fmt.Println(GetMaxLength2(a, 3))
//}
