package main

import (
	"RobCompiler/tokenizer"
	_ "embed"
	"fmt"
	"os"
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
