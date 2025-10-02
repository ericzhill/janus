package main

import (
	"testing"

	"github.com/alecthomas/participle/v2"
)

func TestBasicPlan(t *testing.T) {
	parser, err := participle.Build[CLI](participle.Lexer(cliLexer), participle.Elide("Whitespace"))
	if err != nil {
		t.Fatalf("Failed to build parser: %v", err)
	}
	input := "plan from v1.2.3 to v2.3.4"
	cli, err := parser.ParseString("", input)
	if err != nil {
		t.Fatalf("Failed to parse input: %v", err)
	}
	if cli.Command != "plan" {
		t.Errorf("Expected command 'plan', got '%s'", cli.Command)
	}
	if cli.From.String() != "v1.2.3" {
		t.Errorf("Expected from version 'v1.2.3', got '%s'", cli.From)
	}
	if cli.To.String() != "v2.3.4" {
		t.Errorf("Expected to version 'v2.3.4', got '%s'", cli.To)
	}
}
