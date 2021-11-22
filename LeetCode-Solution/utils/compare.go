package utils

func MinInt(a ...int) int {
	ret := a[0]

	for i := 1; i < len(a); i++ {
		if a[i] < ret {
			ret = a[i]
		}
	}
	return ret
}

func MaxInt(a ...int) int {
	ret := a[0]

	for i := 1; i < len(a); i++ {
		if a[i] > ret {
			ret = a[i]
		}
	}
	return ret
}

func AbsInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
