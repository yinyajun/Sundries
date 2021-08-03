package main

// 问题1： 未排序数组中累加和为给定值的最长子数组？数组中元素有正有负有0

// 无序数组中的元素，如果有正有负有0，那么滑动窗口就不好用了，因为不知道什么时候去扩张窗口，什么时候去收缩窗口

// 除了滑动窗口来维护一个窗口内的sum值，能以O(1)的时间求得一段区间的和外，还有啥方法能以O(1)的时间求得一段区间的和呢？
// 前缀和数组！
// 在不断想用扩展j的过程中，寻找[i+1...j] = s(j) - s(i) == k
// 关键是怎么寻找左边界i? s(i) == s(j) - k
// 也就是说找到某个左边界，它的前缀和为s(j)-k

// 1. 每次都记录到左边的前缀和，然后遍历查找满足条件的左边界，O(N)
// 2. 每次不记录前缀和，每次现算，然后遍历查找满足条件的左边界，O(N^2)
// 3. 每次用map记录左边索引的前缀和，key为遍历过程中的前缀和sum，value为sum出现的最早位置，O(1)

// 使用第三种方式，需要稍微修改下前缀和数组，因为数组这种形式，无法达到O(1)的查找速度
// 使用map记录曾经遍历过的索引的前缀和。key为遇到的sum，value为该sum值最早出现的位置，这样就能在遍历到j的时候，确定满足条件的最长子数组

// 前缀和数组的定义，这里有个坑：因为前缀和 s[i] = sum(arr[0...i])，所以i>=0
// 也就是说计算的区间和为[i+1...j], 最大只能找到[1...j]的区间和。这样，第一个元素被忽略了，这个会造成错误
// 为了补偿第一个元素，i的初始值应该为-1，此时s(-1)=0，那么s(0)-s(-1) = a[0]，符合定义

// 在前缀和数组中，通常通过添加一个哨兵，来添加s(-1)=0
// 这里，直接在map中将s(-1)添加上。key为0，value为-1

// 在[0....j]中寻找i, 满足s(i) = s(j) - k，这样[i+1...j]的sum为k
// 同时注意i>=0, 为了保证[0...j]这样的区间，需要添加s(-1)=0这样的边界情况
func MaxLengthOfSumK(a []int, k int) int {
	sum := 0 // s(j)
	length := 0
	m := make(map[int]int) // 已经遍历过的前缀和map，为了加快查找速度
	m[0] = -1              // s(0)=-1
	for j := 0; j < len(a); j++ {
		sum += a[j]

		// 寻找前缀和=sum-k的左边界
		if i, ok := m[sum-k]; ok && j-i > length {
			length = j - i
		}
		// 将当前sum更新到map中
		if _, ok := m[sum]; !ok {
			m[sum] = j
		}
	}
	return length
}

// 问题2： 无序数组中，元素可正可负可0，求所有子数组中正数与负数相等的最长子数组
// 还是上一题，将正数记为1，负数记为-1，寻找最长的和为0的子数组长度
func MaxLengthOfSumK2(a []int) int {
	b := []int{}
	for i := 0; i < len(a); i++ {
		if a[i] > 0 {
			b = append(b, 1)
		} else if a[i] < 0 {
			b = append(b, -1)
		} else {
			b = append(b, 0)
		}
	}
	return MaxLengthOfSumK(b, 0)
}

// 问题3：无序数组中，元素都是1或0，求所有子数组中0和1相等的最长子数组
func MaxLengthOfSumK3(a []int) int {
	b := []int{}
	for i := 0; i < len(a); i++ {
		if a[i] == 0 {
			b = append(b, -1)
		} else {
			b = append(b, a[i])
		}
	}
	return MaxLengthOfSumK(b, 0)
}

//func main() {
//	a := []int{1, 2, 3, 3}
//	fmt.Println(MaxLengthOfSumK(a, 6))
//	b := []int{1, 2, 3, -2, 3, -6, 0, 1, 0}
//	fmt.Println(MaxLengthOfSumK2(b))
//
//	c := []int{1, 1, 0, 0, 0, 1, 0, 1, 1}
//	fmt.Println(MaxLengthOfSumK3(c))
//}
