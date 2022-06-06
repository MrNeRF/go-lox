package interpreter

import "go-lox/pkg/tokens"

type RuntimeError struct {
	msg string
	tk  tokens.Token
}
