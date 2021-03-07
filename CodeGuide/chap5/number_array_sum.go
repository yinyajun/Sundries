package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	INIT = iota
	SYMBOL
	NUM
)

type lexer struct {
	numbers []int
	token   strings.Builder
	sum     int
}

type State int

func (m *lexer) initToken(ch uint8) State {
	if IsSymbol(ch) {
		m.token.WriteByte(ch)
		return SYMBOL // from INIT to SYMBOL
	}
	if IsNumber(ch) {
		m.token.WriteByte(ch)
		return NUM // from INIT to NUM
	}
	// invalid
	if m.token.Len() > 0 { // store existed token
		m.numbers = append(m.numbers, string2Int(m.token.String()))
		m.sum += string2Int(m.token.String())
	}
	m.token.Reset()
	return INIT
}

func (m *lexer) lex(expr string) {
	var state State = INIT
	for i := range expr {
		ch := expr[i]
		switch state {
		case INIT:
			state = m.initToken(ch)
			break
		case SYMBOL:
			if IsSymbol(ch) { // state remains SYMBOL
				m.token.WriteByte(ch)
			} else if IsNumber(ch) { // SYMBOL to NUM
				state = NUM
				symLen := m.token.Len()
				m.token.Reset()
				if symLen%2 == 1 { // compress
					m.token.WriteByte('-')
				}
				m.token.WriteByte(ch)
			} else {
				state = m.initToken(ch) // SYMBOL to INIT
			}
			break
		case NUM:
			if IsNumber(ch) { // state remains NUM
				m.token.WriteByte(ch)
			} else {
				if m.token.Len() > 0 { // store existed token
					m.numbers = append(m.numbers, string2Int(m.token.String()))
					m.sum += string2Int(m.token.String())
				}
				m.token.Reset()
				state = m.initToken(ch) // NUM to INIT || NUM to SYMBOL
			}
		}
	}
	if state == NUM && m.token.Len() > 0 { // store existed token
		m.numbers = append(m.numbers, string2Int(m.token.String()))
		m.sum += string2Int(m.token.String())
	}
}

func string2Int(str string) int {
	ret, _ := strconv.Atoi(str)
	return ret
}

func IsSymbol(ch uint8) bool {
	return ch == '-'
}
func IsNumber(ch uint8) bool {
	return ch > '0' && ch < '9'
}

func NumberArraySum(expr string) int {
	m := &lexer{[]int{}, strings.Builder{}, 0}
	m.lex(expr)
	return m.sum
}

func NumberArraySum2(expr string) int {
	cond := func(c bool, a, b int) int {
		if c {
			return a
		}
		return b
	}
	res, num, pos := 0, 0, true
	for i := range expr {
		ch := expr[i]
		if IsNumber(ch) {
			num = num*10 + cond(pos, string2Int(string(ch)), -string2Int(string(ch)))
		} else {
			res += num
			num = 0
			if IsSymbol(ch) {
				if i-1 >= 0 && IsSymbol(expr[i-1]) {
					pos = !pos
				} else {
					pos = false
				}
			} else {
				pos = true
			}
		}
	}
	res += num
	return res
}

func main() {
	expr := "fheu---4-51r"
	fmt.Println(NumberArraySum(expr))
	fmt.Println(NumberArraySum2(expr))
}
