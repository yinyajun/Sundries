package main

// 三角形逆时针顺序，如果点都在左边，那么该点是在三角形内部
// 如何判断所给点1、点2、点3能够构成逆时针？
// 1->3有向边，如果点2在其右边，则为逆时针

// 如何判断一个点在 有向边的右侧？
// 叉积：向量，模为两个有向边构成平行四边形面积，方向为其法向量方向（右手定则）

// 通过右手定则，纸面朝外为正
// 可以发现 1->2 叉乘 1->3 > 0, 2在1->3右边（以有向边角度）

func crossProduct(x1, y1, x2, y2 float64) float64 {
	return x1*y2 - x2*y1
}

func isInside2(x1, y1, x2, y2, x3, y3, x, y float64) bool {
	// 先判断是否是逆时针
	if crossProduct(x2-x1, y2-y1, x3-x1, y3-y1) <= 0 { //顺时针
		//将点2和点3兑换，1-2-3 实际上是 1-3-2，这样就变成逆时针了
		x2, y2, x3, y3 = x3, y3, x2, y2
	}

	//1->2 叉乘 1->O > 0 ，在左边
	if crossProduct(x2-x1, y2-y1, x-x1, y-y1) < 0 {
		return false
	}
	// 2->3 , 2->O
	if crossProduct(x3-x2, y3-y2, x-x2, y-y2) < 0 {
		return false
	}
	// 3—>1, 3->O
	if crossProduct(x1-x3, y1-y3, x-x3, y-y3) < 0 {
		return false
	}
	return true
}
