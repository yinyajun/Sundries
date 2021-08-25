package main

import "math"

// 将矩阵旋转到平行坐标轴，这样判断点是否在矩阵中就比较容易了
// 旋转矩阵 =
//  cos    -sin
//  sin    cos

func _isInside(x1, y1, x4, y4, x, y float64) bool {
	if x <= x1 || x >= x4 || y >= y1 || y <= y4 {
		return false
	}
	return true
}

func insideRect(x1, y1, x2, y2, x3, y3, x4, y4, x, y float64) bool {
	if y1 == y2 {
		return _isInside(x1, y1, x4, y4, x, y)
	}

	a := y4 - y3
	b := x4 - x3
	c := math.Sqrt(a*a + b*b)
	sin := a / c
	cos := b / c

	_x1 := x1*cos - y1*sin
	_y1 := x1*sin + y1*cos
	_x4 := x4*cos - y4*sin
	_y4 := x4*sin + y4*cos

	_x := x*cos - y*sin
	_y := x*sin + y*cos

	return _isInside(_x1, _y1, _x4, _y4, _x, _y)
}
