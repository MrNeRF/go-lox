package main

import (
	"fmt"
	"log"
)

type TokenType uint8

const (
	LEFT_PAREN    TokenType = 0
	RIGHT_PAREN   TokenType = 1
	LEFT_BRACE    TokenType = 2
	RIGHT_BRACE   TokenType = 3
	COMMA         TokenType = 4
	DOT           TokenType = 5
	MINUS         TokenType = 6
	PLUS          TokenType = 7
	SEMICOLON     TokenType = 8
	SLASH         TokenType = 9
	STAR          TokenType = 10
	BANG          TokenType = 11
	BANG_EQUAL    TokenType = 12
	EQUAL         TokenType = 13
	EQUAL_EQUAL   TokenType = 14
	GREATER       TokenType = 15
	GREATER_EQUAL TokenType = 16
	LESS          TokenType = 17
	LESS_EQUAL    TokenType = 18

	// Literals.
	IDENTIFIER TokenType = 19
	STRING     TokenType = 20
	NUMBER     TokenType = 21

	// Keywords.
	AND    TokenType = 22
	CLASS  TokenType = 23
	ELSE   TokenType = 24
	FALSE  TokenType = 25
	FUN    TokenType = 26
	FOR    TokenType = 27
	IF     TokenType = 28
	NIL    TokenType = 29
	OR     TokenType = 30
	PRINT  TokenType = 31
	RETURN TokenType = 32
	SUPER  TokenType = 33
	THIS   TokenType = 34
	TRUE   TokenType = 35
	VAR    TokenType = 36
	WHILE  TokenType = 37

	//eof
	EOF TokenType = 38
)

type Token struct {
	tokenType TokenType
	lexeme    string
	line      int
	literal   interface{}
}

func (tk Token) String() string {
	s := fmt.Sprintf("{%v %v %v %v}", tk.tokenType, tk.lexeme, tk.line, tk.literal)
	return s
}

func (tk TokenType) String() string {
	switch tk {
	case LEFT_PAREN:
		return "LEFT_PAREN"
	case RIGHT_PAREN:
		return "RIGHT_PAREN"
	case LEFT_BRACE:
		return "LEFT_BRACE"
	case RIGHT_BRACE:
		return "RIGHT_BRACE"
	case COMMA:
		return "COMMA"
	case DOT:
		return "DOT"
	case MINUS:
		return "MINUS"
	case PLUS:
		return "PLUS"
	case SEMICOLON:
		return "SEMICOLON"
	case SLASH:
		return "SLASH"
	case STAR:
		return "STAR"
	case BANG:
		return "BANG"
	case BANG_EQUAL:
		return "BANG_EQUAL"
	case EQUAL:
		return "EQUAL"
	case EQUAL_EQUAL:
		return "EQUAL_EQUAL"
	case GREATER:
		return "GREATER"
	case GREATER_EQUAL:
		return "GREATER_EQUAL"
	case LESS:
		return "LESS"
	case LESS_EQUAL:
		return "LESS_EQUAL"
	case IDENTIFIER:
		return "IDENTIFIER"
	case STRING:
		return "STRING"
	case NUMBER:
		return "NUMBER"
	case AND:
		return "AND"
	case CLASS:
		return "CLASS"
	case ELSE:
		return "ELSE"
	case FALSE:
		return "FALSE"
	case FUN:
		return "FUN"
	case FOR:
		return "FOR"
	case IF:
		return "IF"
	case NIL:
		return "NIL"
	case OR:
		return "OR"
	case PRINT:
		return "PRINT"
	case RETURN:
		return "RETURN"
	case SUPER:
		return "SUPER"
	case THIS:
		return "THIS"
	case TRUE:
		return "TRUE"
	case VAR:
		return "VAR"
	case WHILE:
		return "WHILE"
	default:
		return ""
	}
}

type Tokenizer struct {
	input     string
	tokenList []Token

	start   int
	current int
	line    int
}

func NewTokenizer(input string) *Tokenizer {
	return &Tokenizer{input: input, start: 0, current: 0, line: 1}
}

// len only works here if characters are a single byte
// which is true for ASCII characters only.
func (t *Tokenizer) scanTokens() {
	for !t.isAtEnd() {
		t.start = t.current
		t.scanToken()
	}
	t.tokenList = append(t.tokenList, Token{tokenType: EOF})
}

func (t *Tokenizer) scanToken() {
	c := t.advance()
	switch c {
	case "(":
		t.addToken(LEFT_PAREN)
	case ")":
		t.addToken(RIGHT_BRACE)
	case "{":
		t.addToken(LEFT_BRACE)
	case "}":
		t.addToken(RIGHT_BRACE)
	case ",":
		t.addToken(COMMA)
	case ".":
		t.addToken(DOT)
	case "-":
		t.addToken(MINUS)
	case "+":
		t.addToken(PLUS)
	case ";":
		t.addToken(SEMICOLON)
	case "*":
		t.addToken(STAR)
	case " ":
	case "\r":
	case "\t":
	case "\n":
		t.line++
	case "\"":
		t.stringLiteral()
	default:
		log.Fatal("Unexpected Character: '", c, "'", " at line ", t.line)
	}
}

// single character tokens
func (t *Tokenizer) addToken(tokenType TokenType) {
	t.tokenList = append(t.tokenList, Token{tokenType: tokenType, lexeme: t.input[t.current : t.current+1], line: t.line, literal: nil})
}

// advance for single characters tokens.
func (t *Tokenizer) advance() string {
	var c string = t.input[t.current : t.current+1]
	t.current++
	return c
}

// lookahead next character
func (t *Tokenizer) peek() string {
	c := t.input[t.current : t.current+1]
	return c
}

func (t *Tokenizer) isAtEnd() bool {
	return t.current >= len(t.input)
}

func (t *Tokenizer) stringLiteral() {
	for t.peek() != "\"" && !t.isAtEnd() {
		if t.peek() == "\n" {
			t.line++
		}
		t.advance()
	}

	if t.isAtEnd() {
		log.Fatal("Unterminated string at line: ", t.line)
		return
	}

	t.advance()

	// Trim the surrounding quotes.
	lexeme := t.input[t.start:t.current]
	value := t.input[t.start+1 : t.current-1]
	t.tokenList = append(t.tokenList, Token{tokenType: STRING, lexeme: lexeme, line: t.line, literal: value})
}
