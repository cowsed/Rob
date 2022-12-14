package main

import (
	"RobCompiler/tokenizer"
	_ "embed"
	"fmt"
	"os"

	"github.com/fatih/color"
)

//go:embed examples/simple.rob
var source string

func main() {
	fmt.Println("===== filename.rob")
	fmt.Println(source)
	fmt.Println("=====")

	fmt.Println("Tokenizing")
	tokens, tok_err := tokenizer.Tokenize(source)

	f, _ := os.Create("token_log.txt")
	fmt.Fprint(f, tokens)
	f.Close()

	fmt.Println("Parsing")
	program, tree_err := Treeify(tokens)

	fmt.Println("Program:")
	fmt.Printf("%+v\n", program)
	fmt.Println("Tokenizing Errors:")
	fmt.Println(tok_err)
	fmt.Println("Tree making Errors:")
	fmt.Println(tree_err)

}

func ColorizeTokens(toks []tokenizer.Token) string {
	s := ""
	for _, t := range toks {
		var printFunc = fmt.Sprint

		switch t.Type {
		case tokenizer.StatementToken:
			printFunc = color.New(color.FgRed).SprintFunc()

		case tokenizer.IdentifierToken:
			printFunc = color.New(color.FgCyan).SprintFunc()

		case tokenizer.CommentToken:
			printFunc = color.New(color.FgHiBlack).SprintFunc()

		case tokenizer.PlusToken, tokenizer.MinusToken, tokenizer.StarToken, tokenizer.DivideToken, tokenizer.RightArrowToken, tokenizer.AsignmentToken,
			tokenizer.DotToken, tokenizer.ColonToken:
			printFunc = color.New(color.FgHiYellow).SprintFunc()

		case tokenizer.StringToken:
			printFunc = color.New(color.FgYellow).SprintFunc()
		}

		s += printFunc(t.Range.String())
	}
	return s
}
