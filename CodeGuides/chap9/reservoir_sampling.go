package main

import "CodeGuide/base/utils"

//随机概率算法
//
// 应用场景：大数据流中随机抽样问题，当内存无法加载全部数据时，如何从包含未知大小的数据流中随机选取k个数据，
// 并且保证每个数据抽到的概率相等。

// 如果数据流中含有N个数，那么每个数被选取的概率要是k/N
// 结论：数据流中第i个数以k/i的概率保留，如果保留了，以1/k的概率替换之前的k个数中的任意一个
// 证明 见 水塘抽样的笔记

// 从包含未知大小的数据流nums中随机选取k个数据
func samplingK(nums []int, k int) []int {
	if len(nums) <= k {
		return nums
	}
	res := []int{}
	// 前k个数
	for i := 0; i < k; i++ {
		res = append(res, nums[i])
	}

	// 以k/i的概率保留
	for i := k; i < len(nums); i++ {
		if utils.Random.Intn(i+1) < k { // [0, k]  k+1个 中选 k个， 该数是否保留？
			// 替换
			res[utils.Random.Intn(k)] = nums[i]
		}
	}
	return res
}
