package tokens

import "fmt"

type Token struct {
	tokenType TokenType
	lexeme    string
	line      int
	literal   interface{}
}

func (tk Token) String() string {
	s := fmt.Sprintf("{%v, %v, %v, %v}", tk.tokenType, tk.lexeme, tk.line, tk.literal)
	return s
}
