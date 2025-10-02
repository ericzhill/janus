package main

import (
	"log/slog"
	"os"
	"strings"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"janus"
)

var cliLexer = lexer.MustSimple([]lexer.SimpleRule{
	{Name: "Whitespace", Pattern: `\s+`},
	{Name: "Version", Pattern: `v\d+\.\d+\.\d+`},
	{Name: "Ident", Pattern: `[a-zA-Z_][a-zA-Z0-9_]*`},
	{Name: "Punct", Pattern: `[-[\]{}()*+,./:;]`},
})

type CLI struct {
	Command string      `parser:"@('plan' | 'migrate')"`
	From    janus.Version `parser:"'from' @Version"`
	To      janus.Version `parser:"'to' @Version"`
}

func main() {
	parser, err := participle.Build[CLI](participle.Lexer(cliLexer), participle.Elide("Whitespace"))
	if err != nil {
		slog.Error("Error building parser", "error", err.Error())
		os.Exit(1)
	}

	cli := &CLI{}
	input := strings.Join(os.Args[1:], " ")
	cli, err = parser.ParseString("", input)
	if err != nil {
		slog.Error("Error parsing command line", "error", err.Error())
		os.Exit(1)
	}

	slog.Info("Parsed command", "command", cli.Command)
}
