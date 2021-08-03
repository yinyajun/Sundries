package main

import "fmt"

func FindString(a []string, target string) int {
	return findString2(a, target, 0, len(a)-1)
}

// 这里的比较应该用compare接口实现，这里偷了个懒
func findString(a []string, target string, lo, hi int) int {
	res := -1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		fmt.Println(lo, hi, mid)
		if a[mid] == target { // 如果找到了，向左缩小搜索区间
			hi = mid - 1
			res = mid // 找到的情况下更新res
		} else if a[mid] != "" {
			if a[mid] < target {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		} else { // a[mid]==""，直接线性查找
			i := mid - 1
			for i >= lo && a[i] != "" { // 从右往左找第一个不为空的
				i--
			}
			// a[i] == "" || i== lo-1
			if i < lo { // 左边全部为空，去右边找
				lo = mid + 1
			} else if a[i] < target {
				lo = mid + 1
			} else if a[i] > target {
				hi = i - 1
			} else { // a[i]==target  如果找到了，向左缩小搜索区间
				res = i
				hi = i - 1
			}
		}
	}
	fmt.Println(lo) // 如果找不到，那么lo可能就不对
	return res
}

// 简化版本,
func findString2(a []string, target string, lo, hi int) int {
	res := -1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if a[mid] == target {
			res = mid
			hi = mid - 1
		} else if a[mid] != "" {
			if a[mid] < target {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		} else {
			i := mid - 1
			for i >= lo && a[i] != "" {
				i--
			}
			if i < lo || a[i] < target {
				lo = mid + 1
			} else {
				if a[i] == target {
					res = i
				}
				hi = i - 1
			}
		}
	}
	return res
}

func main() {
	ret := FindString([]string{"", "a", "", "a", "", "b", "", "c"}, "3")
	fmt.Println(ret)
}
