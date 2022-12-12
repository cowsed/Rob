package main

import "fmt"

var sourceToFilename = map[*[]rune]string{}

type UnclosedStringLiteral struct {
	srange SourceRange
}

func (usl UnclosedStringLiteral) Error() string {
	filename := sourceToFilename[usl.srange.source_runes]
	if filename == "" {
		filename = "--filename unknown--"
	}
	return fmt.Sprintf("unclosed string literal. %s: at %d:%d", filename, usl.srange.start, usl.srange.start)
}
