package parser

import (
	"fmt"
	"strconv"
	"unicode"
)

type ParseError struct {
	Message string
	Pos     int
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("%s, %d", e.Message, e.Pos)
}

type Parser struct {
	expr   string
	pos    int
	length int
}

func EvalExpr(expr string) (float64, error) {
	p := &Parser{
		expr:   expr,
		pos:    0,
		length: len(expr),
	}
	return p.parseExpression()
}

func (p *Parser) parseExpression() (float64, error) {
	return p.parseAddSub()
}

func (p *Parser) parseAddSub() (float64, error) {
	left, err := p.parseMulDiv()
	if err != nil {
		return 0, err
	}
	for {
		op := p.peek()
		if op != '+' && op != '-' {
			break
		}
		p.consume()
		right, err := p.parseMulDiv()
		if err != nil {
			return 0, err
		}
		switch op {
		case '+':
			left += right
		case '-':
			left -= right
		}
	}
	return left, nil
}

func (p *Parser) parseMulDiv() (float64, error) {
	left, err := p.parseUnary()
	if err != nil {
		return 0, err
	}
	for {
		op := p.peek()
		if op != '*' && op != '/' {
			break
		}
		p.consume()
		right, err := p.parseUnary()
		if err != nil {
			return 0, err
		}
		switch op {
		case '*':
			left *= right
		case '/':
			if right == 0 {
				return 0, &ParseError{"division by zero", p.pos}
			}
			left /= right
		}
	}
	return left, nil
}

func (p *Parser) parseUnary() (float64, error) {
	op := p.peek()
	if op == '+' || op == '-' {
		p.consume()
		val, err := p.parseUnary()
		if err != nil {
			return 0, err
		}
		if op == '-' {
			val = -val
		}
		return val, nil
	}
	return p.parsePrimary()
}

func (p *Parser) parsePrimary() (float64, error) {
	current := p.peek()
	if current == '(' {
		p.consume()
		expr, err := p.parseExpression()
		if err != nil {
			return 0, err
		}
		if p.peek() != ')' {
			return 0, &ParseError{"expected ')'", p.pos}
		}
		p.consume()
		return expr, nil
	}
	return p.parseNumber()
}

func (p *Parser) parseNumber() (float64, error) {
	start := p.pos
	for p.pos < p.length && (unicode.IsDigit(rune(p.expr[p.pos])) || p.expr[p.pos] == '.') {
		p.pos++
	}
	if start == p.pos {
		return 0, &ParseError{"expected number", p.pos}
	}
	num, err := strconv.ParseFloat(p.expr[start:p.pos], 64)
	if err != nil {
		return 0, &ParseError{"invalid num", start}
	}
	return num, nil
}

func (p *Parser) peek() byte {
	if p.pos >= p.length {
		return 0
	}
	return p.expr[p.pos]
}

func (p *Parser) consume() {
	if p.pos < p.length {
		p.pos++
	}
} 