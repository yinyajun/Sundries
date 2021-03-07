package main

import "fmt"

// 易错题，需要多次练习

func StatisticString(a string) string {
	cnt := 0
	ret := []byte{}
	var cur byte
	for i := range a {
		if a[i] != cur { // 在遇到新元素的时候更新
			if cnt > 0 {
				if len(ret) > 0 {
					ret = append(ret, '_')
				}
				ret = append(ret, cur)
				ret = append(ret, '_')
				ret = append(ret, Num2Bytes(cnt)...)
			}
			cur = a[i]
			cnt = 0
		}
		cnt++
	}

	if cnt > 0 {
		if len(ret) > 0 {
			ret = append(ret, '_')
		}
		ret = append(ret, cur)
		ret = append(ret, '_')
		ret = append(ret, Num2Bytes(cnt)...)
	}
	return string(ret)
}

func Num2Bytes(num int) []byte {
	//num >0
	ret := []byte{}
	for num != 0 {
		ret = append(ret, byte(num%10)+'0')
		num = num / 10
	}

	ret2 := make([]byte, len(ret))
	for i := range ret2 {
		ret2[i] = ret[len(ret)-1-i]
	}
	return ret2
}

func GetIndexString(stats string, idx int) string {
	var curByte byte
	cnt := 0
	sum := 0
	for i := 0; i < len(stats); {
		curByte = stats[i]
		i += 2
		// find cnt
		for i < len(stats) && stats[i] >= '0' && stats[i] <= '9' {
			cnt = cnt*10 + int(stats[i]-'0')
			i++
		}
		i += 1 //此时i已经到了_了
		sum += cnt
		cnt = 0
		fmt.Println(sum)
		if sum > idx {
			return string(curByte)
		}
	}
	return string('@')
}
func main() {
	s := StatisticString("aabbbbbbbbbbbbbbbbbbbaaaaaaaaaaaaaaaaaaaaaaaaaaaacba")
	fmt.Println(s)
	fmt.Println(GetIndexString(s, 123))
}
