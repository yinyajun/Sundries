package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"github.com/gammazero/deque"
)

// 正则文法，递归下降文法，通过文法来定义优先级，越往下优先级越高
// additive -> multiplicative ((+|-) multiplicative) *
// multiplicative -> primary ((*|/) primary)*
// primary -> IntLiteral | (additive)

// 48 * ((70-65)-43)+8*1

// ----------------------------------
type parser2 struct {
	tokens *bufio.Reader
}

func (p *parser2) tokensPeek() byte {
	tok, err := p.tokens.Peek(1)
	if err == nil {
		return tok[0]
	}
	return '0'
}

func (p *parser2) tokensRead() byte {
	b, _ := p.tokens.ReadByte()
	return b
}

func (p *parser2) cond(expr bool, a, b int) int {
	if expr {
		return a
	}
	return b
}

// additive -> multiplicative ((+|-) multiplicative) *
func (p *parser2) additive(tokens *bufio.Reader) (int, error) {
	cur, err1 := p.multiplicative(tokens)
	if err1 == nil {
		for {
			token := p.tokensPeek()
			if token != '0' && (isAdd(token) || isMinus(token)) {
				token = p.tokensRead()
				b, err2 := p.multiplicative(tokens)
				if err2 == nil {
					cur = p.cond(isAdd(token), cur+b, cur-b)
				} else {
					panic("expected right part")
				}
			} else {
				break
			}
		}
	}
	return cur, err1
}

// mul := prim ((*|/) prim)*
func (p *parser2) multiplicative(tokens *bufio.Reader) (int, error) {
	cur, err1 := p.primary(tokens)
	if err1 == nil {
		for {
			token := p.tokensPeek()
			if token != '0' && (isStar(token) || isSlash(token)) {
				token = p.tokensRead()
				b, err2 := p.primary(tokens)
				if err2 == nil {
					cur = p.cond(isStar(token), cur*b, cur/b)
				} else {
					panic("expected right part")
				}
			} else {
				break
			}
		}
	}
	return cur, err1
}

// prim := int | (additive)
func (p *parser2) primary(tokens *bufio.Reader) (int, error) {
	token := p.tokensPeek()
	if token != '0' {
		if isNumber(token) {
			cur := 0
			for isNumber(token) {
				p.tokensRead()
				cur = 10*cur + Byte2Num(token)
				token = p.tokensPeek()
			}
			// !isNumber(token)
			return cur, nil
		} else if isLeftBracket(token) {
			p.tokensRead()
			c, err := p.additive(tokens)
			if err == nil {
				token = p.tokensPeek()
				if isRightBracket(token) {
					p.tokensRead()
					return c, nil
				} else {
					panic("expected right bracket")
				}
			} else {
				panic("expected additive in brackets")
			}
		}
	}
	// token == '0'
	return -1, errors.New("additive error")
}

func (p *parser2) Evaluate() {
	res, _ := p.additive(p.tokens)
	fmt.Println(res)
}

// ----------------------------------
// 这里实现没有使用peek预读，使用右结合来规避左递归
type parser struct{}

func (p *parser) cond(expr bool, a, b int) int {
	if expr {
		return a
	}
	return b
}

// 右递归方式，有结合性错误
// add := mul | mul + add
func (p *parser) additive(tokens *bufio.Reader) (int, error) {
	a, err1 := p.multiplicative(tokens)
	token, _ := tokens.ReadByte()
	if err1 == nil && (isAdd(token) || isMinus(token)) {
		b, err2 := p.additive(tokens)
		if err2 == nil {
			return p.cond(isAdd(token), a+b, a-b), nil
		} else {
			panic("expecting right part")
		}
	}
	tokens.UnreadByte()
	return a, err1
}

func (p *parser) multiplicative(tokens *bufio.Reader) (int, error) {
	a, err1 := p.primary(tokens)
	token, _ := tokens.ReadByte()
	if err1 == nil && (isStar(token) || isSlash(token)) {
		b, err2 := p.multiplicative(tokens)
		if err2 == nil {
			return p.cond(isStar(token), a*b, a/b), nil
		} else {
			panic("expecting right part")
		}
	}
	tokens.UnreadByte()
	return a, err1
}

func (p *parser) primary(tokens *bufio.Reader) (int, error) {
	token, _ := tokens.ReadByte()
	if token != '0' {
		if isNumber(token) {
			cur := 0
			for isNumber(token) {
				cur = 10*cur + Byte2Num(token)
				token, _ = tokens.ReadByte()
			}
			// !isNumber(token)
			tokens.UnreadByte() // 重要
			return cur, nil
		}
		if isLeftBracket(token) {
			c, err := p.additive(tokens)
			if err == nil {
				token, _ = tokens.ReadByte()
				if isRightBracket(token) {
					return c, nil
				} else {
					panic("expected right bracket")
				}
			} else {
				panic("expected additive in brackets")
			}
		}
	}
	tokens.UnreadByte()
	return -1, errors.New("additive error")
}

func (p *parser) Evaluate(tokens *bufio.Reader) {
	res, _ := p.additive(tokens)
	fmt.Println(res)

}

func Byte2Num(a byte) int {
	return int(a - '0')
}

func isStar(a byte) bool {
	return a == '*'
}

func isSlash(a byte) bool {
	return a == '/'
}

func isLeftBracket(a byte) bool {
	return a == '('
}

func isRightBracket(a byte) bool {
	return a == ')'
}

func isAdd(a byte) bool {
	return a == '+'
}

func isMinus(a byte) bool {
	return a == '-'
}

func isNumber(a byte) bool {
	return a > '0' && a < '9'
}

// ----------------------------------
// 针对上述方法的简化版本
// 遍历过程中，将加法表达式写入到一个deque中，然后从左向右计算，这样就没有结合性问题
// 遍历过程中，遇到数先缓存，直到遇到法则符号，先将数写入到deque中，并且查看之前的符号是否是乘法（如果是，直接计算结果存入），最后将现在的法则符号写入
// 如果遇到(，那么进入递归，直到遇到）返回，返回括号内的值以及遍历到的位置

// 这里的递归非常有意思，在遍历a的过程中，发生的递归, 将a拆分了多个层级去遍历。（对于本身非递归结构的数据去递归）
// 递归过程类似于下面的图示
// .....(						  .........
// 		 .......(		 ........)
// 				 .......)

// 根据上图，可以发现，遇到左括号进入下一个递归函数，需要开始index，并且遇到要返回结束的index, 当然也需要返回括号内
// 由于这里是线性递归栈，可以类比链表的递归
func evaluate2(a []byte, i int) (val, j int) {
	pre := 0 // 数字的结束在遇到下一个符号，用此缓存遇到下一个符号之前的数字
	deq := new(deque.Deque)
	for i < len(a) && a[i] != ')' {
		if isNumber(a[i]) {
			pre = pre*10 + Byte2Num(a[i])
			i++
		} else if a[i] == '(' { // 正常情况下，（前是计算符号，不是一个数
			a, b := evaluate2(a, i+1)
			pre = a
			i = b + 1
		} else { // +-*/
			AddNum(deq, pre)
			deq.PushBack(a[i])
			pre = 0
			i++
		}
	}
	// base case: i ==len(a) || a[i] = ')'
	AddNum(deq, pre)
	val = getNum(deq)
	j = i
	return
}

func evaluate(a string, i int) (val int, j int) {
	pre := 0 // 记录数值的缓存
	deq := new(deque.Deque)

	for i < len(a) {
		if a[i] == ')' {
			AddNum(deq, pre)
			val = getNum(deq) // 将deque计算的结果作为val
			j = i
			return
		}
		if isNumber(a[i]) {
			pre = pre*10 + Byte2Num(a[i])
			i++
		} else if a[i] != '(' { // 缓存计算数，遇到下一个符号
			AddNum(deq, pre)   // 将计算数写入deque
			deq.PushBack(a[i]) // 将符号写入deque
			pre = 0
			i++
		} else { // 遇到（
			a, b := evaluate(a, i+1)
			pre = a
			i = b + 1
		}
	}
	// 此时cur可能不为0
	AddNum(deq, pre)
	val = getNum(deq) // 将deque计算的结果作为val
	j = i
	return
}

// 将计算数num写入deque
func AddNum(deq *deque.Deque, num int) {
	if deq.Len() != 0 {
		symbol := deq.PopBack().(byte)
		if isAdd(symbol) || isMinus(symbol) { // 加减法，啥也不做
			deq.PushBack(symbol)
		} else { // 乘除法，做计算
			pre := deq.PopBack().(int)
			if isStar(symbol) {
				num = pre * num
			} else {
				num = pre / num
			}
		}
	}
	deq.PushBack(num)
}

func getNum(deq *deque.Deque) int {
	res := 0
	addFlag := true
	for deq.Len() > 0 {
		cur := deq.PopFront()
		if v, ok := cur.(byte); ok && isAdd(v) {
			addFlag = true
		} else if v, ok := cur.(byte); ok && isMinus(v) {
			addFlag = false
		} else { // number
			if addFlag {
				res += cur.(int)
			} else {
				res -= cur.(int)
			}
		}
	}
	return res
}

func main() {
	a := []byte("2*(2+5)/2")
	reader := bufio.NewReader(bytes.NewReader(a))
	p := &parser2{reader}
	p.Evaluate()
	fmt.Println(evaluate(string(a), 0))
	fmt.Println(evaluate2(a, 0))
}
