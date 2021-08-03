package main

import "fmt"

// 两个字符串s和p，最常见的方式就是分别有两个指针在s和p上移动
// 如果指针都能移动到最后，那么就匹配成功

// str中不能有.和*，pattern中*不能是首字符，任意两个*不能相邻
func isValidString(str, pattern string) bool {
	for i := range str {
		if str[i] == '.' || str[i] == '*' {
			return false
		}
	}

	for i := range pattern {
		if pattern[i] == '*' && (i == 0 || pattern[i-1] == '*') {
			return false
		}
	}
	return true
}

// 不含有*的pattern
func match1(str, pattern string) bool {
	i, j := 0, 0
	for i < len(str) && j < len(pattern) {
		if str[i] == pattern[j] || pattern[i] == '.' {
			i++
			j++
		} else {
			return false
		}
	}
	return i == j
}

/*

if (s[i] == p[j] || p[j] == '.') {
    // 匹配
    if (j < p.size() - 1 && p[j + 1] == '*') {
        // 有 * 通配符，可以匹配 0 次或多次
    } else {
        // 无 * 通配符，老老实实匹配 1 次
        i++; j++;
    }
} else {
    // 不匹配
    if (j < p.size() - 1 && p[j + 1] == '*') {
        // 有 * 通配符，只能匹配 0 次
    } else {
        // 无 * 通配符，匹配无法进行下去了
        return false;
    }
}

bool dp(string& s, int i, string& p, int j) {
    if (s[i] == p[j] || p[j] == '.') {
        // 匹配
        if (j < p.size() - 1 && p[j + 1] == '*') {
            // 1.1 通配符匹配 0 次或多次
            return dp(s, i, p, j + 2)
                || dp(s, i + 1, p, j);
        } else {
            // 1.2 常规匹配 1 次
            return dp(s, i + 1, p, j + 1);
        }
    } else {
        // 不匹配
        if (j < p.size() - 1 && p[j + 1] == '*') {
            // 2.1 通配符匹配 0 次
            return dp(s, i, p, j + 2);
        } else {
            // 2.2 无法继续匹配
            return false;
        }
    }
}
*/

func StringPattern(str, pattern string) bool {
	return process(str, pattern, 0, 0)
}

func process(str, pattern string, si, pi int) bool {
	if pi == len(pattern) {
		return si == len(str)
	}
	if pi+1 == len(pattern) || pattern[pi+1] != '*' { // 不用*通配符匹配
		return si != len(str) &&
			(pattern[pi] == str[si] || pattern[pi] == '.') && // 匹配
			process(str, pattern, si+1, pi+1) // 匹配且常规匹配1次（否则匹配失败）
	}
	// 使用通配符匹配
	for si != len(str) && (pattern[pi] == str[si] || pattern[pi] == '.') { // 匹配
		// 通配符匹配0次或者多次
		if process(str, pattern, si, pi+2) { // 能否匹配0次
			return true
		}
		si++ // 匹配多次
	}
	// str已经end或者不匹配
	return process(str, pattern, si, pi+2) // 通配符匹配0次
}

func StringPattern2(str, pattern string) bool {
	return process2(str, pattern, 0, 0)
}

// 定义递归函数process(str, pattern string, si, pi int)： str[si...]是否能匹配pattern[pi...]
// case1：如果pattern匹配到最后，那么如果str还有未匹配的，自然匹配失败
// case2：如果str匹配到最后，那么如果pattern还有未匹配的，关注pattern[pi...]能否匹配""(通配符匹配下才有可能)
// case3：如果pattern的下一个pattern[pi+1]不为*，使用常规匹配
//   # 如果pi和si相匹配，那么关注pattern[pi+1...]是否匹配str[si...]
//   # 如果pi和si不匹配，直接返回false
// case4：如果pattern的下一个pattern[pi+1]为*，使用通配符匹配
//   # 如果pi和si相匹配，那么可以通配符匹配0次或者多次
//       $ 匹配0次：(aaaXXX, a*aaaYYY), 且XXX和YYY匹配则可以返回true，关注str[si...]是否匹配pattern[pi+2...]
//       $ 匹配1次：(aaaXXX, a*aaYYY), 且XXX和YYY匹配则可以返回true，关注str[si+1...]是否匹配pattern[pi+2...]
//       $ 匹配2次：(aaaXXX, a*aYYY), 且XXX和YYY匹配则可以返回true，关注str[si+2...]是否匹配pattern[pi+2...]
//       $ 匹配3次：(aaaXXX, a*YYY), 且XXX和YYY匹配则可以返回true，关注str[si+3...]是否匹配pattern[pi+2...]
//       $ ...
//		 str[si]==pattern[pi]&&process(si, pi) = (str[si]==pattern[pi]&&process(si, pi+2)) ||
//												 (str[si]==pattern[pi] && process(si+1 ,pi+2)) ||
//											     (str[si]==pattern[pi] && process(si+2 ,pi+2)) || ...
//   # 如果pi和si不相匹配，那么只能通配符匹配0次，那么关注pattern[pi...]是否匹配str[si+2...]

func process2(str, pattern string, si, pi int) bool {
	// case1
	if pi == len(pattern) {
		return si == len(str)
	}
	// case4: 通配符匹配模式
	if pi+1 < len(pattern) && pattern[pi+1] == '*' {
		// pi和si能匹配，那么可以匹配0次或者多次
		// zero:= process(str, pattern, si, pi+2)
		// once:= process(str, pattern, si+1, pi+2)
		// twice:= process(str, pattern, si+2, pi+2)
		// ...
		for si < len(str) && pattern[pi] == str[si] || pattern[pi] == '.' {
			if process2(str, pattern, si, pi+2) {
				return true
			} else {
				si++ // 向后多匹配一次
			}
		}
		// pi和si不能匹配，只能匹配0次(或者case2)
		return process2(str, pattern, si, pi+2)
	} else {
		// case3：常规匹配模式
		if si < len(str) {
			if pattern[pi] == str[si] || pattern[pi] == '.' {
				return process2(str, pattern, si+1, pi+1) // 常规匹配1次
			} else {
				return false // 无法常规匹配
			}
		} else {
			// case2: str结束而pattern未结束，在常规匹配模式下，无法继续匹配
			return false
		}
	}
}

// 优化case4, 使用类似完全背包的方式降低复杂度
// 在通配符匹配模式下且当前可以匹配时，可以0次或者多次匹配
// str[si]==pattern[pi]&&process(si, pi) = (str[si]==pattern[pi]&&process(si, pi+2)) ||
//										   (str[si]==pattern[pi] && process(si+1 ,pi+2)) ||
//										   (str[si]==pattern[pi] && process(si+2 ,pi+2)) || ...
// 可以发现
// str[si]==pattern[pi]&&process(si+1, pi) = (str[si]==pattern[pi]&&process(si+1, pi+2)) ||
//										   (str[si]==pattern[pi] && process(si+2 ,pi+2)) || ...
// 所以
// str[si]==pattern[pi]&&process(si, pi) = (str[si]==pattern[pi]&&process(si, pi+2)) ||
//										   str[si]==pattern[pi]&&process(si+1, pi)

func StringPattern3(str, pattern string) bool {
	return process3(str, pattern, 0, 0)
}

func process3(str, pattern string, si, pi int) bool {
	// case1
	if pi == len(pattern) {
		return si == len(str)
	}
	// case4: 通配符匹配模式
	if pi+1 < len(pattern) && pattern[pi+1] == '*' {
		if si < len(str) && pattern[pi] == str[si] || pattern[pi] == '.' {
			return process3(str, pattern, si, pi+2) || process3(str, pattern, si+1, pi)
		}
		// pi和si不能匹配，只能匹配0次(或者case2)
		return process3(str, pattern, si, pi+2)
	} else {
		// case3：常规匹配模式
		if si < len(str) && pattern[pi] == str[si] || pattern[pi] == '.' {
			return process3(str, pattern, si+1, pi+1) // 常规匹配1次
		} else {
			return false // 无法常规匹配 or case2: str结束而pattern未结束，在常规匹配模式下，无法继续匹配
		}
	}
}

// 可以使用记忆化递归来优化时间复杂度，递归中参数只有si和pi，同时每个递归函数的时间为常数，所以总时间复杂度为O(MN)
// 当然也可以转为动态规划来解决
// 状态就是si，pi
// 选择就是pattern[pi]匹配几个字符
// dp[i][j]: str[i...]是否匹配pattern[j...]（str[...i-1]已经成功匹配pattern[...j-1]）
// dp[i][j] = dp[i][j+2] || dp[i+1][j], if pattern[j+1]=='*' && str[i]== pattern[j]
//          = dp[i][j+2], if pattern[j+1]=='*' && str[i]!= pattern[j]
//          = dp[i+][j+1], if pattern[j+1]!='*' && str[i] == pattern[j]
//          = false, otherwise
// target: dp[0][0]
// base: dp[M][N] = true， dp[M][0...N-1]必须是"x*x*"这种，且非*的位置为true，否则为false
// base: dp[0...M-1][N] = false
// base: dp[M-1][N-1] = true if str[M-1] match pattern[N-1] else false, dp[1...M-2][N-1]=false
func StringPattern4(str, pattern string) bool {
	m, n := len(str), len(pattern)
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	// init
	dp[m][n] = true
	for j := n - 2; j >= 0; j -= 2 {
		if pattern[j] != '*' && pattern[j+1] == '*' {
			dp[m][j] = true
		}
	}
	if pattern[n-1] == '.' || str[m-1] == pattern[n-1] {
		dp[m-1][n-1] = true
	}
	// iterate
	for i := m - 1; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			if pattern[j+1] == '*' {
				if str[i] == pattern[j] || pattern[j] == '.' {
					dp[i][j] = dp[i][j+2] || dp[i+1][j]
				} else {
					dp[i][j] = dp[i][j+2]
				}
			} else {
				dp[i][j] = (str[i] == pattern[j] || pattern[j] == '.') && dp[i+1][j+1]
			}
		}
	}
	return dp[0][0]
}

func main() {
	str, pattern := "abbbc", "ab*d*c"
	fmt.Println(StringPattern(str, pattern))
	fmt.Println(StringPattern2(str, pattern))
	fmt.Println(StringPattern3(str, pattern))
	fmt.Println(StringPattern4(str, pattern))
}
