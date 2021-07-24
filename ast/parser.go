package ast

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer/stateful"
)

var (
	lexer = stateful.MustSimple([]stateful.Rule{
		{Name: "String", Pattern: `"(\\"|[^"])*"`, Action: nil},
		{Name: "Number", Pattern: `(\d*\.)?\d+`, Action: nil},
		{Name: "Go", Pattern: `go`, Action: nil},
		{Name: "Initial", Pattern: `initial`, Action: nil},
		{Name: "Ident", Pattern: `[a-zA-Z_]\w*`, Action: nil},
		{Name: "Paren", Pattern: `[\{\}\(\)]`, Action: nil},
		{Name: "Assign", Pattern: `<=`, Action: nil},
		{Name: "Op", Pattern: `[-+=,#]`, Action: nil},
		{Name: "EOL", Pattern: `[\n\r]+`, Action: nil},
		{Name: "whitespace", Pattern: `[ \t]+`, Action: nil},
	})
	parser = participle.MustBuild(&ProgramFile{},
		participle.Lexer(lexer),
		participle.CaseInsensitive("Ident"),
		participle.Unquote("String"),
		participle.Elide("whitespace", "EOL"),
		participle.UseLookahead(2),
	)
)

func getParser() *participle.Parser {
	return parser
}
