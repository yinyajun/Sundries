package main

// O(2^n)
func NumSteps1(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 {
		return n
	}
	return NumSteps1(n-1) + NumSteps1(n-2)
}

func NumSteps2(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 {
		return n
	}

	dp_0, dp_1 := 1, 2
	for i := 3; i <= n; i++ {
		dp_1, dp_0 = dp_1+dp_0, dp_1
	}
	return dp_1
}

func NumSteps3(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 {
		return n
	}
	A := [][]int{{1, 1}, {1, 0}}
	res := MatrixPower(A, n-2)
	return 2*res[0][0] + res[0][1]
}
