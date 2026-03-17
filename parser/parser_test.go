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
	checkParserErrors(t, p)
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

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser had %d errors", len(errors))
	for _, err := range errors {
		t.Errorf("parser error: %s", err)
	}
	t.FailNow()
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

func TestReturnStatements(t *testing.T) {
	input := `
	return 5;
	return 10;
	return 993322;`
	statementsNum := 3

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != statementsNum {
		t.Fatalf("Program contains %d statements. got: %d",
			statementsNum, len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("Expected a *ast.ReturnStatement, got: %T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral is not 'return'. got: '%s'",
				returnStmt.TokenLiteral())
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program contains 1 statement. got: %d", len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got: %T",
			program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp not *ast.Identifier. got: %T", stmt.Expression)
	}
	if ident.Value != "foobar" {
		t.Errorf("ident.Value is not %s. got:%s", "foobar", ident.Value)
	}
	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral is not %s. got: %s", "foobar", ident.TokenLiteral())
	}
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "55;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program contains 1 statement. got: %d", len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got: %T",
			program.Statements[0])
	}

	integer, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp not *ast.IntegerLiteral. got: %T", stmt.Expression)
	}
	if integer.Value != 55 {
		t.Errorf("literal.Value is not %d. got: %d", 55, integer.Value)
	}
	if integer.TokenLiteral() != "55" {
		t.Errorf("literal.TokenLiteral is not %s. got: %s", "55",
			integer.TokenLiteral())
	}
}
