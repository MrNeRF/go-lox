package main

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
	for t.current <= len(t.input) {
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
		break
	case ")":
		t.addToken(RIGHT_BRACE)
		break
	case "{":
		t.addToken(LEFT_BRACE)
		break
	case "}":
		t.addToken(RIGHT_BRACE)
	}

}

// single character tokens
func (t *Tokenizer) addToken(tokenType TokenType) {
	t.tokenList = append(t.tokenList, Token{tokenType: tokenType, lexeme: t.input[t.current : t.current+1], line: t.line, literal: nil})
}

// advance for single characters tokens.
func (t *Tokenizer) advance() string {
	return t.input[t.current : t.current+1]
}
