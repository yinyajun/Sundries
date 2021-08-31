package main

import (
	"CodeGuide/base/utils"
	"math"
)

// random在[0,x)上的概率为x
// 那么两次random，可能会有如下3种结果
// 1. 两次结果都在[0,x)：x^2
// 2. 两次结果都在[x,1): (1-x)^2
// 3. 一次在x内，一次在x外: 2x-2x^2

// 随机变量max(y1,y2) 在[0,x)上的概率为x^2

func randomPowerK(k int) float64 {
	if k <= 1 {
		return 0
	}
	var res float64
	for i := 0; i < k; i++ {
		res = math.Max(res, utils.Random.Float64())
	}
	return res
}
