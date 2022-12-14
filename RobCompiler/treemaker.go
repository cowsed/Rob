package main

import (
	"RobCompiler/tokenizer"
	"fmt"
)

type Program struct {
	moduleName string
	exposing   *Exposes

	imports []Import
}

func Treeify(tokens []tokenizer.Token) (*Program, error) {
	source := &TokenSource{
		tokens:     tokens,
		tokenIndex: 0,
		program: &Program{
			moduleName: "",
			exposing:   nil,
			imports:    []Import{},
		},
		errors: []error{},
	}

	var state stateFn = lookForStatement
	for !source.Finished() && state != nil {
		fmt.Printf("source.Peek(): %v\n", source.Peek())
		state = state(source)
	}

	if !source.Finished() {
		fmt.Println("====== Finished Early")
		source.EmitError(FailedParseEarly{})
	}

	if source.program.exposing == nil {
		source.EmitError(NoModule{})
	}

	fmt.Println("======= ERRORS ===============", len(source.errors))
	fmt.Printf("%T\n", source.errors[0])
	fmt.Printf("%T\n", source.errors[1])
	fmt.Printf("%T\n", source.errors[2])

	if len(source.errors) < 1 {
		//we had no errors - we can have a little unsafety as a treat
		fmt.Println("NO ERRORS")
		source.errors = nil
	}

	return source.program, source.errors
}

type stateFn func(ts *TokenSource) stateFn

// the outer level "looking around" function.
// a program can be thought of blocks
// blocks of module defintion, type definitions, function definition
func lookForStatement(ts *TokenSource) stateFn {
	tok := ts.Peek()
	//If it was something totally unexpected
	if tok.Type != tokenizer.StatementToken {
		ts.EmitError(ExpectedStatement{tok.Range})
		ts.Take()
		return lookForStatement
	}

	switch tok.Range.String() {
	case "type":
		return parseTypeDefinition
	case "module":
		return parseModuleDeclaration
	case "import":
		return parseImport
	default:
		ts.EmitError(UnknownStatement{})
		ts.takeUntilNewline()
	}

	return nil
}

func parseModuleDeclaration(ts *TokenSource) stateFn {
	parts := ts.takeUntilNewline()
	fmt.Println("took module stuff")
	fmt.Println(parts)
	where := mergeRangesOfTokens(parts)

	if len(parts) < 1 {
		ts.EmitError(ModuleNeedsName{where})
		return lookForStatement
	}
	fmt.Println("Looking for identifier")
	//find first non Whitespace character after 'module' - thats the name
	var name string
	for i := 1; i < len(parts); i++ {
		if parts[i].Type == tokenizer.SpaceToken {
			continue
		}
		if parts[i].Type != tokenizer.IdentifierToken {
			fmt.Println("was in fact a ", parts[i].Type.String())
			ts.EmitError(ModuleNeedsName{
				where: where,
			})
			fmt.Println("error emitted")
			return lookForStatement
		} else {
			name = parts[i].Range.String()
		}
	}

	ts.program.exposing = &Exposes{
		name:     name,
		exposing: []string{},
	}
	fmt.Println("Made exposing")

	return lookForStatement
}
func mergeRangesOfTokens(toks []tokenizer.Token) tokenizer.SourceRange {
	if len(toks) < 1 {
		return tokenizer.SourceRange{}
	}
	source := toks[0].Range.SourceRunes
	merged := tokenizer.SourceRange{
		Start:       toks[0].Range.Start,
		End:         toks[0].Range.End,
		SourceRunes: source,
	}
	for _, tok := range toks {
		if source != tok.Range.SourceRunes {
			panic("I can't merge two ranges from different files. This should never happen really - its a compiler issue")
		}
		merged.Start = min(merged.Start, tok.Range.Start)
		merged.End = max(merged.End, tok.Range.End)

	}
	return merged
}

func parseImport(ts *TokenSource) stateFn {
	return nil
}

func parseTypeDefinition(ts *TokenSource) stateFn {
	return nil
}

type TokenSource struct {
	tokens     []tokenizer.Token
	tokenIndex int

	program *Program
	errors  tokenizer.ComboError
}

func (ts *TokenSource) EmitError(err error) {
	ts.errors = append(ts.errors, err)
}

func (ts *TokenSource) Peek() tokenizer.Token {
	return ts.tokens[ts.tokenIndex]
}

func (ts *TokenSource) Take() tokenizer.Token {
	ts.tokenIndex++
	if ts.tokenIndex >= len(ts.tokens) {
		return ts.tokens[len(ts.tokens)-1]
	}
	return ts.tokens[ts.tokenIndex-1]
}

func (ts *TokenSource) takeUntilNewline() []tokenizer.Token {
	toks := []tokenizer.Token{}
	for {
		t := ts.Take()
		toks = append(toks, t)
		if t.Type == tokenizer.NewlineToken {
			break
		}
	}
	return toks
}

func (ts *TokenSource) Finished() bool {
	return ts.tokenIndex >= len(ts.tokens)
}

type Import struct {
	name string
}

type Exposes struct {
	name     string
	exposing []string
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
