package main

func swap2Number(a, b *int) {
	*a = *a ^ *b
	*b = *a ^ *b // a^b^b=a
	*a = *a ^ *b // a^a^b=b
}

//func main() {
//	a, b := 2, 3
//	swap2Number(&a, &b)
//	fmt.Println(a, b)
//}
