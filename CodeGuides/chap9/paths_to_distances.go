package main

import (
	"fmt"
)

// time O(N); space O(1)

// 暴力解法，时间复杂度是满足要求的，空间复杂度远远超出了要求
func paths2Distances(paths []int) []int {
	res := make([]int, len(paths))
	memo := make(map[int]int)

	for i := 0; i < len(paths); i++ {
		res[i] = path(paths, i, memo)
	}

	// 根据距离数组转换为统计数组
	memo = make(map[int]int)
	for i := 0; i < len(paths); i++ {
		memo[res[i]]++
	}

	for i := 0; i < len(paths); i++ {
		res[i] = memo[i]
	}
	return res
}

var cnt int

func indent(cnt int) {
	for i := 0; i < cnt; i++ {
		fmt.Print("    ")
	}
}

// 这里的树，是指向父节点的树，通过记忆化递归，可以用O(N)的时间，计算出所有节点的深度
// 额外空间：O(N) + O(树的深度)
func path(paths []int, s int, memo map[int]int) int {
	//indent(cnt)
	//cnt++
	//fmt.Println("enter", s)

	if s == paths[s] {
		memo[s] = 0
		//cnt--
		//indent(cnt)
		//fmt.Println("return", s, memo)
		return 0
	}

	c, ok := memo[paths[s]]
	if ok {
		memo[s] = c + 1
	} else {
		memo[s] = 1 + path(paths, paths[s], memo)
	}

	//cnt--
	//indent(cnt)
	//fmt.Println("return", s, ok, memo)
	return memo[s]
}

func paths2Distances2(paths []int) []int {
	res := make([]int, len(paths))
	memo := make(map[int]int)

	for i := 0; i < len(paths); i++ {
		res[i] = -1 * path(paths, i, memo)
	}

	fmt.Println(res)
	// 根据距离数组转换为统计数组，这里不使用额外空间
	// 为了复用res数组，这里将path置为相反数，为了区别统计数组
	dist2stat2(res)
	return res
}

// 思路就是在遍历dist数组的过程中，每个位置代表当前节点到根节点的距离
// 当遍历到一个位置的时候，记录下该位置的距离后，该位置就可以空出来记录别的东西了
// 这里有个trick，dist数组使用负数表示，当dist<=0的时候，代表该位置到根节点的距离的相反数
// 而转化为统计数组的时候，距离作为index，值为统计个数，此时的值为非负
func dist2stat(dist []int) {
	for i := 0; i < len(dist); i++ {
		d := dist[i]
		if d < 0 {
			dist[i] = 0
		}

		for d < 0 { // 代表距离
			// 将-d记录
			if dist[d*-1] < 0 { // 记录位置仍然存储距离
				temp := dist[d*-1]
				dist[d*-1] = 1
				d = temp
			} else { // 记录位置存储统计
				dist[d*-1]++
				d = dist[d*-1]
			}
		}
	}
	// dist[i]==0 omitted
	dist[0] = 1
}

func dist2stat2(dist []int) {
	for i := 0; i < len(dist); i++ { // 开始更新距离首都为i的地方
		d := dist[i]
		if d <= 0 { // 代表距离，将距离为i的位置初始化为统计值
			dist[i] = 0

			// -d >= 0
			for { // d此时代表距离，表示距离首都为-d的城市，需要将d转为统计存储到dist[-d]
				if dist[-d] >= 0 { // dist[-d]表征统计，可以存储
					dist[-d]++
					break
				}
				// dist[-d]表征距离，通过循环更新，先存储dist[-d]，然后将dist[-d]置1，表示已经遇到1个d距离的城市
				dist[-d], d = 1, dist[-d]
			}
		}
	}
}

func paths2Distances3(paths []int) {
	// 直接在paths数组上，求得距离数组
	paths2dist2(paths)
	fmt.Println(paths)
	// 根据距离数组转换为统计数组，这里不使用额外空间, 为了复用res数组，这里将path置为相反数，为了区别统计数组
	dist2stat2(paths)
	fmt.Println(paths)
}

func paths2dist(paths []int) {
	//paths[i] < 0 表征距离， paths[i] >=0 表征父节点
	i := 0

	if paths[i] >= 0 {
		cur := paths[i]
		paths[i] = -1
		pre := i

		for paths[cur] != cur {
			if paths[cur] >= 0 {
				next := paths[cur]
				paths[cur] = pre
				pre = cur
				cur = next
			} else {
				break
			}
		}
		var value int
		if paths[cur] == cur {
			value = 0
		} else {
			value = paths[cur]
		}

		for paths[pre] != -1 {
			lastPre := paths[pre]
			value--
			paths[pre] = value
			cur = pre
			pre = lastPre
		}
		value--
		paths[pre] = value
	}
	fmt.Println(paths)
}

func paths2dist2(paths []int) {
	//paths[i] < 0 表征距离， paths[i] >=0 表征父节点
	capital := 0
	var cur, pre int
	var value int
	for i := 0; i < len(paths); i++ {
		if paths[i] == i {
			capital = i
			continue // 不写的话，paths[root]将会出错，
		}

		if paths[i] >= 0 { // 节点表征父节点
			// 标志开始节点
			cur = paths[i]
			paths[i] = -1 // 确保paths[i]不是root
			pre = i

			for paths[cur] != cur { // 未达到根节点
				if paths[cur] < 0 { // 节点表征距离
					break
				}
				// 节点表征父节点
				paths[cur], pre, cur = pre, cur, paths[cur] //向根节点移动，同时paths[cur]记录子节点，以便返回
			}
			if paths[cur] == cur {
				value = 0
			} else { // 在寻找根节点的路径上遇到祖先节点已经计算完距离了
				value = paths[cur]
			}

			// 从根节点或者已经计算完距离的节点开始，反推到开始节点
			for paths[pre] != -1 { // 未遍历到开始节点
				value--
				pre, paths[pre] = paths[pre], value
			}
			// 起始节点
			value--
			paths[pre] = value
		}

		fmt.Println(i, paths)
	}
	fmt.Println(capital)
	//paths[capital] = 0
}

//func main() {
//	paths := []int{9, 1, 4, 9, 0, 4, 8, 9, 0, 1}
//	paths2Distances3(paths)
//}
