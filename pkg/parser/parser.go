package parser

import (
	"errors"
	"go-lox/pkg/tokens"
	"log"
)

type Parser struct {
	current int
	input   []tokens.Token
}

func (p *Parser) Parse() (Expr, error) {
	return expression(p)
}

func NewParser(input []tokens.Token) *Parser {
	return &Parser{current: 0, input: input}
}

func expression(p *Parser) (Expr, error) {
	return equality(p)
}

func equality(p *Parser) (Expr, error) {
	expr, err := comparison(p)

	if err != nil {
		return nil, err
	}
	for match(p, tokens.BANG_EQUAL, tokens.EQUAL_EQUAL) {
		operator := p.previous()
		right, err := comparison(p)
		if err != nil {
			return nil, err
		}
		expr = Binary{expr, operator, right}
	}
	return expr, nil
}

func comparison(p *Parser) (Expr, error) {
	expr, err := term(p)
	if err != nil {
		return nil, err
	}
	for match(p, tokens.GREATER, tokens.GREATER_EQUAL, tokens.LESS, tokens.LESS_EQUAL) {
		operator := p.previous()
		right, err := term(p)
		if err != nil {
			return nil, err
		}
		expr = Binary{expr, operator, right}
	}
	return expr, nil
}

func term(p *Parser) (Expr, error) {
	expr, err := factor(p)
	if err != nil {
		return nil, err
	}
	for match(p, tokens.MINUS, tokens.PLUS) {
		operator := p.previous()
		right, err := factor(p)
		if err != nil {
			return nil, err
		}
		expr = Binary{expr, operator, right}
	}
	return expr, nil
}

func factor(p *Parser) (Expr, error) {
	expr, err := unary(p)
	if err != nil {
		return nil, err
	}
	for match(p, tokens.SLASH, tokens.STAR) {
		operator := p.previous()
		right, err := unary(p)
		if err != nil {
			return nil, err
		}
		expr = Binary{expr, operator, right}
	}

	return expr, nil
}

func unary(p *Parser) (Expr, error) {
	if match(p, tokens.BANG, tokens.MINUS) {
		operator := p.previous()
		right, err := unary(p)
		if err != nil {
			return nil, err
		}
		return Unary{operator, right}, err
	}
	return primary(p)
}

func primary(p *Parser) (Expr, error) {
	if match(p, tokens.FALSE) {
		return Literal{value: false}, nil
	}

	if match(p, tokens.TRUE) {
		return Literal{value: true}, nil
	}

	if match(p, tokens.NIL) {
		return Literal{value: nil}, nil
	}

	if match(p, tokens.NUMBER, tokens.STRING) {
		l := p.previous()
		return Literal{value: (&l).GetLiteral()}, nil
	}

	if match(p, tokens.LEFT_PAREN) {
		expr, err := expression(p)
		if err != nil {
			return nil, err
		}

		_, err = consume(p, tokens.RIGHT_PAREN, "Expected ')' afeter expression")
		if err != nil {
			return nil, err
		}

		return Grouping{expr}, err
	}

	return nil, errors.New("No primary match!")
}

func consume(p *Parser, tt tokens.TokenType, msg string) (tokens.Token, error) {
	if p.check(tt) {
		return p.advance(), nil
	}
	log.Panic(p.peek(), msg)

	return tokens.Token{}, errors.New("Not consuming Token")
}

func match(p *Parser, tt ...tokens.TokenType) bool {
	for _, t := range tt {
		if p.check(t) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) check(tt tokens.TokenType) bool {
	if p.isAtEnd() {
		return false
	}

	tk := p.peek()
	return (&tk).GetTokenType() == tt
}

func (p *Parser) isAtEnd() bool {
	tk := p.peek()
	return (&tk).GetTokenType() == tokens.EOF
}

func (p *Parser) advance() tokens.Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *Parser) peek() tokens.Token {
	return p.input[p.current]
}

func (p *Parser) previous() tokens.Token {
	return p.input[p.current-1]
}
