package tokens

import "log"

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
	t.tokenList = append(t.tokenList, Token{tokenType: EOF})
}

func (t *Tokenizer) GetTokenList() []Token {
	return t.tokenList
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
