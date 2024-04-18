package lexer

import (
	"interpreter-go/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	// 테스트할 input
	input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
  x + y;
};

let result = add(five, ten);
`

	// 렉서로 파싱하였을 때 예상되는 토큰 리스트
	expectedTokens := []token.Token{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lexer := New(input)

	for i, expectedToken := range expectedTokens {
		lexerToken := lexer.NextToken()
		if lexerToken.Type != expectedToken.Type {
			t.Fatalf("expectedTokens[%d] - tokentype wrong. expected=%q, got=%q", i, expectedToken.Type, lexerToken.Type)
		}
		if lexerToken.Literal != expectedToken.Literal {
			t.Fatalf("expectedTokens[%d] - literal wrong. expected=%q, got=%q", i, expectedToken.Literal, lexerToken.Literal)
		}
	}
}
