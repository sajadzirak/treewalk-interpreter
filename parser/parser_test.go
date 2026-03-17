package parser

import (
	"Guerilla/ast"
	"Guerilla/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
		let x = 5;
		let y = 10;
		let foobar = 858585;
	`
	statementsNum := 3

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != statementsNum {
		t.Fatalf("Program contains %d statements. got: %d",
			statementsNum, len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}
	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, expectedIdent string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("Expected token literal let. got: %s", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("Expected a *ast.LetStatement, got: %T", s)
		return false
	}

	if letStmt.Name.Value != expectedIdent {
		t.Errorf("letStmt.Name.Value is not '%s'. got: '%s'", expectedIdent,
			letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != expectedIdent {
		t.Errorf("letStmt.Name.TokenLiteral is not '%s'. got: '%s'", expectedIdent,
			letStmt.Name.TokenLiteral())
		return false
	}

	return true
}
