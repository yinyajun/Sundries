package main

import (
	"fmt"
	"github.com/spaolacci/murmur3"
	"math"
)

/*
100亿条黑名单url，每个url占64B
url过滤，容忍万分之一以下的误判，使用空间不要超过30GB
如果使用map，所需要的空间为64B*100*10^8 = 640GB，远超出要求

bloom filter代表一个集合，精确判断一个元素是否在集合中。（精确但不是完全正确）
优势在于：使用很少的空间就能将准确率做到很高

原理：
1. 首先仍然是一个bitmap
2. 1个输入对于bitmap的影响，不仅仅是1位，而是多位
3. 这么设计是因为，仅仅1位的话，太容易发生碰撞
4. 而多位同时碰撞的概率就大大降低
5. 注意这里被影响的多位必须相互独立，所以采用多个独立的hash函数来做到这个程度
6. 将所有数据依次作为输入，影响bitmap后，此时，一个bloom filter完成
7. 检查一个值是否在bloom filter中？通过多个hash func计算出多个index，如果其bitmap对应的位置都不为0，则说明在集合中（可能有假阳）；若任一位置为0，则一定不再集合中

bloom会有false positive，没有false negative

而bloom的参数设计，和输入规模相关

bloom中bitmap的占用m个bit，有k个hash func

n次输入后，这个位置没有被涂黑的概率为：（1-1/m）^(kn)
n次输入后，这个位置被涂黑的概率为：1 - （1-1/m）^(kn)

k个位置都为黑的概率（误判率）： (1 - （1-1/m）^(kn))^k ~ (1 - exp(-kn/m)) ^ k   重要极限

当确定了m，n，此时误判率就仅仅是k的函数，k取何值时，能使误判率最低？
f(k) = (1 - exp(-nk/m)) ^ k
令b=exp(n/m)
f(k) = (1 - b^(-k)) ^ k
求最值： b^(-k) = 1/2
exp(-kn/m) = 1/2
k = ln2*m/n = 0.7 * m /n (根据m,n，确定k)

将k带入误判率公式，可以得到：
m = - n * ln Pe / (ln2 * ln2)   (根据n和p，确定m)
*/

type bloomFilter struct {
	m    uint // bitmap size
	k    uint // hash func num
	size uint
	bm   *bitmap
}

// p：允许的误判率
// n：输入规模
func NewBloomFilter(p float64, n uint) *bloomFilter {
	m := BFOptSize(p, n)
	k := BFOptHashFuncNum(m, n)
	fmt.Println(m, k)
	return &bloomFilter{
		m:  m,
		k:  k,
		bm: NewBitmap(m),
	}
}

func BFOptSize(p float64, n uint) uint {
	return uint(math.Ceil(-float64(n) * math.Log(p) / (math.Log(2) * math.Log(2))))
}

func BFOptHashFuncNum(m, n uint) uint {
	return uint(math.Ceil(math.Log(2) * float64(m) / float64(n)))
}

func (f *bloomFilter) ErrorRate() float64 {
	if f.m > 100 {
		//(1 - exp(-kn/m)) ^ k
		return math.Pow(1-math.Exp(-float64(f.k)*float64(f.Size())/float64(f.m)), float64(f.k))
	}
	//(1 - （1-1/m）^(kn))^k
	return math.Pow(1-math.Pow(1-1./float64(f.m), float64(f.k)*float64(f.Size())), float64(f.k))
}

func (f *bloomFilter) Add(x string) {
	for i := uint(0); i < f.k; i++ {
		idx := uint(murmur3.Sum32WithSeed([]byte(x), uint32(i))) % f.m
		f.bm.Add(idx)
	}
	f.size++
}

func (f *bloomFilter) Check(x string) bool {
	for i := uint(0); i < f.k; i++ {
		idx := uint(murmur3.Sum32WithSeed([]byte(x), uint32(i))) % f.m
		if !f.bm.Find(idx) {
			return false
		}
	}
	return true
}

func (f *bloomFilter) Size() uint { return f.size }

func main() {
	filter := NewBloomFilter(0.01, 100)
	filter.Add("123")
	filter.Add("456")
	filter.Add("789")
	filter.Add("234")
	filter.Add("567")
	filter.Add("9097")
	fmt.Println(filter.ErrorRate())
	fmt.Println(filter.Check("678"))
	fmt.Println(filter.Check("123"))
	fmt.Println(filter.Check("456"))
	fmt.Println(filter.Check("789"))
}
