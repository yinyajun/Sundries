/*
There are N gas stations along a circular route, where the amount of gas at station i is gas[i].
You have a car with an unlimited gas tank, and it costs cost[i] of gas to travel from station i to its next
station (i+1). You begin the journey with an empty tank at one of the gas stations.
Return the starting gas station’s index if you can travel around the circuit once, otherwise return -1.
Note: The solution is guaranteed to be unique.

* @Author: Yajun
* @Date:   2021/11/26 10:42
*/

package chap2

// time: O(n^2); space: O(1)
func canCompleteCircuit(gas, cost []int) int {
	var (
		n        = len(gas)
		sum, idx int
	)
	for i := 0; i < n; i++ {
		sum = 0
		for j := 0; j < n; j++ {
			idx = (i + j) % n
			sum += gas[idx] - cost[idx]
			if sum < 0 {
				break
			}
			if j == n-1 {
				return i
			}
		}
	}
	return -1
}

// time: O(n); space: O(1)
// total<0 铁定没有解；sum用来筛选可能的起点，再用total验证一下
func canCompleteCircuitB(gas, cost []int) int {
	var (
		n          = len(gas)
		sum, total int
		j          = -1
	)
	for i := 0; i < n; i++ {
		sum += gas[i] - cost[i]
		total += gas[i] - cost[i]
		if sum < 0 {
			sum = 0
			j = i
		}
	}
	if total >= 0 {
		return j + 1
	}
	return -1
}
