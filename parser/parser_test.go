package parser

import (
	"interpreter-go/ast"
	"interpreter-go/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	// 여기서 Statements는 LetStatement임. let 구문이 3개이므로 길이는 3임
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got %d", len(program.Statements))
	}

	expectedIdentifierNames := []string{"x", "y", "foobar"}

	for i, expectedIdentifierName := range expectedIdentifierNames {
		if !testLetStatements(t, program.Statements[i], expectedIdentifierName) {
			return
		}
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %s", msg)
	}
	t.FailNow()
}

func testLetStatements(t *testing.T, s ast.Statement, expectedName string) bool {
	// let을 가진 토큰인지 체크
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got %s", s.TokenLiteral())
		return false
	}

	letStatement, ok := s.(*ast.LetStatement)

	// statement가 LetStatement 타입으로 래핑되도록 제대로 구현되었는지 체크
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	// letStatement의 변수명이 테스트케이스와 일치하는지 체크
	if letStatement.Name.Value != expectedName {
		t.Errorf("letStatement.Name.Value not '%s'. got %s", expectedName, letStatement.Name.Value)
		return false
	}

	// 토큰 리터럴 또한 테스트케이스와 일치하는지 체크
	if letStatement.Name.TokenLiteral() != expectedName {
		t.Errorf("letStatement.Name.TokenLiteral() not '%s'. got %s", expectedName, letStatement.Name.TokenLiteral())
		return false
	}

	return true
}

func TestReturnStatements(t *testing.T) {
	input := `
		return 5;
		return 10;
		return 993322;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	// parser가 파싱한 결과물이 3줄이 맞는지 체크
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	for _, statement := range program.Statements {
		returnStatement, ok := statement.(*ast.ReturnStatement) // ReturnStatement로 형변환
		if !ok {
			t.Errorf("returnStatement not *ast.ReturnStatement. got=%T", statement)
			continue
		}

		// 현재 루프의 statement가 return 문법인지 체크
		if returnStatement.TokenLiteral() != "return" {
			t.Errorf("returnStatement.TokenLiteral not 'return'. got=%q", returnStatement.TokenLiteral())
		}
	}
}
