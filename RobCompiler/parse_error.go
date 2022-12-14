package main

import (
	"RobCompiler/tokenizer"

	"github.com/fatih/color"
)

var red = color.New(color.FgRed).SprintFunc()
var gray = color.New(color.FgHiBlack).SprintFunc()

type ExpectedStatement struct {
	where tokenizer.SourceRange
}

func (e ExpectedStatement) Error() string {
	message := tokenizer.ErrorTemplate("Expected Keyword", e.where, "I was looking for something like `module`, `import`, `type` or an identifier for a function or a variable. But instead i found this. ")
	return message
}

type FailedParseEarly struct {
}

func (f FailedParseEarly) Error() string {
	return red("============== Compiler Failed Early =======================") + "\nFor some reason a parsing func return nil too early - this is indicative of a goof on the compiler writer. Please go yell at your local compiler nerd"
}

type UnknownStatement struct {
	where tokenizer.SourceRange
}

func (u UnknownStatement) Error() string {
	message := tokenizer.ErrorTemplate("Unknown Statement", u.where, "I was told i would get a keyword but I do not recognize this. (This is probably a tokenizing error - contact your nearest compiler writer)")
	return message
}

type NoModule struct{}

func (n NoModule) Error() string {
	message := "I expected a module definition on the first line but did not find one or maybe the one written was incorrect\n"
	message += "I need a line telling me the name of the module and what functions or types it exposes."
	message += "It should be something like `module Name exposing A, B, C"
	//TODO add more info
	return tokenizer.ErrorTemplate("No Module Specification", tokenizer.SourceRange{}, message)
}

type ModuleNeedsName struct {
	where tokenizer.SourceRange
}

func (m ModuleNeedsName) Error() string {
	note := "I expected a name after a module declaration.\nIt should look something like `module Name`"
	message := tokenizer.ErrorTemplate("Module Requires a Name", m.where, note)
	return message
}

type ExpectedIdentifer struct {
	where tokenizer.SourceRange
	why   string
}

func (e ExpectedIdentifer) Error() string {
	return tokenizer.ErrorTemplate("Expected Identifier", e.where, e.why)
}
