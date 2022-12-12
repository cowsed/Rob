package main

import (
	_ "embed"
	"fmt"
	"log"
)

//go:embed examples/simple.rob
var source string

func main() {
	log.Println(source)

	log.Println("Tokenizing")
	tokens, err := Tokenize(source)
	fmt.Println(tokens)
	fmt.Println("errors:\n", err)

}
