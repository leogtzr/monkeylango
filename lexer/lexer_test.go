package main

import (
	"fmt"
	"testing"

	"github.com/leogtzr/monkeylango/token"
)

func TestNextToken(t *testing.T) {
	t.Parallel()

	input := `=+(){},;`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	fmt.Println(input)
	fmt.Println(tests)
}
