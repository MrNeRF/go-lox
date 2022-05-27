package tokens

import (
	"log"
	"strconv"
)

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
func (t *Tokenizer) ScanTokens() {
	for !t.isAtEnd() {
		t.start = t.current
		t.scanToken()
	}
	t.addToken(&Token{tokenType: EOF, line: t.line, literal: nil})
}

func (t *Tokenizer) GetTokenList() []Token {
	return t.tokenList
}

func (t *Tokenizer) scanToken() {
	c := t.advance()
	switch c {
	case "(":
		t.addToken(&Token{tokenType: LEFT_PAREN, lexeme: t.input[t.start:t.current], line: t.line, literal: nil})
	case ")":
		t.addToken(&Token{tokenType: RIGHT_PAREN, lexeme: t.input[t.start:t.current], line: t.line, literal: nil})
	case "{":
		t.addToken(&Token{tokenType: LEFT_BRACE, lexeme: t.input[t.start:t.current], line: t.line, literal: nil})
	case "}":
		t.addToken(&Token{tokenType: RIGHT_BRACE, lexeme: t.input[t.start:t.current], line: t.line, literal: nil})
	case ",":
		t.addToken(&Token{tokenType: COMMA, lexeme: t.input[t.start:t.current], line: t.line, literal: nil})
	case ".":
		t.addToken(&Token{tokenType: DOT, lexeme: t.input[t.start:t.current], line: t.line, literal: nil})
	case "-":
		t.addToken(&Token{tokenType: MINUS, lexeme: t.input[t.start:t.current], line: t.line, literal: nil})
	case "+":
		t.addToken(&Token{tokenType: PLUS, lexeme: t.input[t.start:t.current], line: t.line, literal: nil})
	case ";":
		t.addToken(&Token{tokenType: SEMICOLON, lexeme: t.input[t.start:t.current], line: t.line, literal: nil})
	case "*":
		t.addToken(&Token{tokenType: STAR, lexeme: t.input[t.start:t.current], line: t.line, literal: nil})
	case " ":
	case "\r":
	case "\t":
	case "\n":
		t.line++
	case "\"":
		t.addstringLiteralToken()
	case "/":
		t.addSlashToken()
	case "!", "=", "<", ">":
		t.addOperatorToken()
	default:
		if isDigit(c) {
			t.addnumberToken()
		} else {
			log.Fatal("Unexpected Character: '", c, "'", " at line ", t.line)
		}
	}
}

func (t *Tokenizer) addToken(token *Token) {
	t.tokenList = append(t.tokenList, *token)
}

// advance for single characters tokens.
func (t *Tokenizer) advance() string {
	var c string = t.input[t.current : t.current+1]
	t.current++
	return c
}

// lookahead next character
func (t *Tokenizer) peek() string {
	if t.isAtEnd() {
		return "\000"
	}
	c := t.input[t.current : t.current+1]
	return c
}

// we have a BIG problem here if the number is at EOF.
func (t *Tokenizer) peekNext() string {
	if t.current+1 >= len(t.input) {
		return "\000"
	}
	return t.input[t.current+1 : t.current+2]
}

func (t *Tokenizer) isAtEnd() bool {
	return t.current >= len(t.input)
}

func (t *Tokenizer) addstringLiteralToken() {
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
	t.addToken(&Token{tokenType: STRING, lexeme: t.input[t.start:t.current], line: t.line, literal: t.input[t.start+1 : t.current-1]})
}

func (t *Tokenizer) addOperatorToken() {
	c := t.input[t.start:t.current]
	switch c {
	case "!":
		if t.match("=") {
			t.addToken(&Token{tokenType: BANG_EQUAL, lexeme: "!=", line: t.line, literal: nil})
		} else {
			t.addToken(&Token{tokenType: BANG, lexeme: "!", line: t.line, literal: nil})
		}
	case "=":
		if t.match("=") {
			t.addToken(&Token{tokenType: EQUAL_EQUAL, lexeme: "==", line: t.line, literal: nil})
		} else {
			t.addToken(&Token{tokenType: EQUAL, lexeme: "=", line: t.line, literal: nil})
		}
	case "<":
		if t.match("=") {
			t.addToken(&Token{tokenType: LESS_EQUAL, lexeme: "<=", line: t.line, literal: nil})
		} else {
			t.addToken(&Token{tokenType: LESS, lexeme: "<", line: t.line, literal: nil})
		}
	case ">":
		if t.match("=") {
			t.addToken(&Token{tokenType: GREATER_EQUAL, lexeme: ">=", line: t.line, literal: nil})
		} else {
			t.addToken(&Token{tokenType: GREATER, lexeme: ">", line: t.line, literal: nil})
		}
	}
}

func (t *Tokenizer) match(m string) bool {
	if t.isAtEnd() {
		return false
	}
	if t.input[t.current:t.current+1] != m {
		return false
	}
	t.current++
	return true
}

func isDigit(s string) bool {
	return s[0] >= '0' && s[0] <= '9'
}

func (t *Tokenizer) addnumberToken() {
	for isDigit(t.peek()) {
		t.advance()
	}
	if t.peek() == "." && isDigit(t.peekNext()) {
		t.advance()
	}
	for isDigit(t.peek()) {
		t.advance()
	}
	floatnum, err := strconv.ParseFloat(t.input[t.start:t.current], 64)
	if err == nil {
		t.addToken(&Token{tokenType: NUMBER, lexeme: t.input[t.start:t.current], line: t.line, literal: floatnum})
	}
}

func (t *Tokenizer) addSlashToken() {
	//TODO
}
