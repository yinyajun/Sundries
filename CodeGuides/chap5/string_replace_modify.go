package main

import "fmt"

func ReplaceString(a []byte) {
	// space num
	numSpace := 0
	index := 0 // 左半区的大小
	//for i := range a {
	//	if a[i] != 0 {
	//		index++
	//		if a[i] == ' ' {
	//			numSpace++
	//		}
	//	}
	//}

	// 提前结束循环
	for ; index < len(a) && a[index] != 0; index++ {
		if a[index] == ' ' {
			numSpace++
		}
	}
	newIndex := index + 2*numSpace

	for i, j := index-1, newIndex-1; i >= 0; {
		if a[i] == ' ' {
			a[j] = '0'
			a[j-1] = '2'
			a[j-2] = '%'
			i--
			j -= 3
		} else {
			a[j] = a[i]
			i--
			j--
		}
	}
	fmt.Println(string(a))
}

func Modify(a []byte) {
	i := len(a) - 1
	j := len(a) - 1
	for i >= 0 {
		if a[i] != '*' {
			a[j] = a[i]
			j--
		}
		i--
	}
	for k := j; k >= 0; k-- {
		a[k] = '*'
	}
	fmt.Println(string(a))
}

func main() {
	a := []byte{'a', ' ', 'b', ' ', ' ', 'c', 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	ReplaceString(a)
	Modify([]byte("12**345"))
}
