package main

// 辗转相除法
// 证明如下：gcd(m, n)
// m/n=k...r
// to prove gcd(m,n) = gcd(n, r)
// gcd(m, n)=c ==> m=ac, n=bc
// r=m-kn = ac-kbc = (a-kb)c
// 可见余数r也能被gcd c整除
// m = kn+r ,n , r 都有公约数c

func gcd(m, n int) int {
	if n == 0 {
		return m
	}
	return gcd(n, m%n)
}

//func main() {
//	fmt.Println(gcd(45, 36))
//}
