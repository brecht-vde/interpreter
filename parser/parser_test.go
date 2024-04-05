package parser

import (
	"interpreter/ast"
	"interpreter/lexer"
	"testing"
)

func TestLetStatement(t *testing.T) {
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
		t.Fatal("program returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("expected 3 statements but got %d", len(program.Statements))
	}

	test := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range test {
		statement := program.Statements[i]

		if !testLetStatement(t, statement, tt.expectedIdentifier) {
			return
		}
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.errors

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser got %d errors", len(errors))

	for _, msg := range errors {
		t.Errorf("parser error: %v", msg)
	}

	t.FailNow()
}

func testLetStatement(t *testing.T, s ast.Statement, i string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("expected 'let', got '%v'", s.TokenLiteral())
		return false
	}

	letStatement, ok := s.(*ast.LetStatement)

	if !ok {
		t.Errorf("expected *ast.LetStament, got %v", s)
		return false
	}

	if letStatement.Name.Value != i {
		t.Errorf("expected %v, got %v", i, letStatement.Name.Value)
		return false
	}

	if letStatement.Name.TokenLiteral() != i {
		t.Errorf("expected %v, got %v", i, letStatement.Name)
		return false
	}

	return true
}
