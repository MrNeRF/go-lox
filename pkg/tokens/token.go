package tokens

import "fmt"

type Token struct {
	tokenType TokenType
	lexeme    string
	line      int
	literal   interface{}
}

func NewToken(tt TokenType, lx string, ln int, lit interface{}) *Token {
	return &Token{tokenType: tt, lexeme: lx, line: ln, literal: lit}
}

func (tk Token) String() string {
	s := fmt.Sprintf("{%v %v %v %v}", tk.tokenType, tk.lexeme, tk.line, tk.literal)
	return s
}

func (tk *Token) GetTokenType() TokenType {
	return tk.tokenType
}

func (tk *Token) GetLiteral() interface{} {
	return tk.literal
}

func (tk *Token) GetLexeme() string {
	return tk.lexeme
}
