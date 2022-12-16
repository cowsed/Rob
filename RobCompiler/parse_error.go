package main

import (
	"RobCompiler/tokenizer"

	"github.com/fatih/color"
)

var red = color.New(color.FgRed).SprintFunc()
var gray = color.New(color.FgHiBlack).SprintFunc()

type ExpectedTopLevel struct {
	where tokenizer.SourceRange
}

func (e ExpectedTopLevel) Error() string {
	message := tokenizer.ErrorTemplate("Expected Statement or Declaration", e.where, "I was looking for something like `module`, `import`, `type` or an identifier for a function or a variable. But instead i found this. ")
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

type ImportNeedsName struct {
	where tokenizer.SourceRange
}

func (i ImportNeedsName) Error() string {
	note := "I expected a name after an import statement.\nIt should look something like `import ModuleName`"
	message := tokenizer.ErrorTemplate("Module Requires a Name", i.where, note)
	return message
}

type ExpectedIdentifer struct {
	where tokenizer.SourceRange
	why   string
}

func (e ExpectedIdentifer) Error() string {
	return tokenizer.ErrorTemplate("Expected Identifier", e.where, e.why)
}

type ModuleRedeclaration struct {
	where tokenizer.SourceRange
}

func (m ModuleRedeclaration) Error() string {
	return tokenizer.ErrorTemplate("Module Redeclaration", m.where, "Module redeclared here. There should only ever be one module declaration per file and it should be the first line")
}

type UnknownIdentifierUsage struct {
	where tokenizer.SourceRange
}

func (u UnknownIdentifierUsage) Error() string {
	return tokenizer.ErrorTemplate("Unknown Identifier Usage", u.where, "I found an identifer here but im not sure what to do with it. Usually when I see an identifier like this it is for a `Type Annotation` name: Type1 -> Type2 or a declaration name t = t + 1")
}

type UnexpectedTypeInAnnotation struct {
	where tokenizer.SourceRange
}

func (u UnexpectedTypeInAnnotation) Error() string {
	return tokenizer.ErrorTemplate("Malformed Type in Annotation", u.where, "I was expecting a type, a tuple of types, or a record type but I found this instead")
}

type CustomError struct {
	name        string
	where       tokenizer.SourceRange
	description string
}

func (s CustomError) Error() string {
	return tokenizer.ErrorTemplate(s.name, s.where, s.description)
}
