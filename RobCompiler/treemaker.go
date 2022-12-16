package main

import (
	"RobCompiler/tokenizer"
	"fmt"
)

type AST struct {
	toplevel []TopASTNode
}

func Treeify(tokens []tokenizer.Token) (*AST, error) {
	source := &TokenSource{
		tokens:     tokens,
		tokenIndex: 0,
		program: &AST{
			toplevel: []TopASTNode{},
		},
		errors: []error{},
	}

	var state stateFn = lookForOuterStatement
	for !source.Finished() && state != nil {
		fmt.Printf("source.Peek(): %v\n", source.Peek())
		state = state(source)
	}

	if !source.Finished() {
		fmt.Println("====== Finished Early")
		source.EmitError(FailedParseEarly{})
	}

	if len(source.errors) < 1 {
		//we had no errors - we can have a little unsafety as a treat
		source.errors = nil
	}

	return source.program, source.errors
}

type stateFn func(ts *TokenSource) stateFn

// the outer level "looking around" function.
// a program can be thought of blocks
// blocks of module defintion, type definitions, function definition
func lookForOuterStatement(ts *TokenSource) stateFn {
	tok := ts.Peek()

	if tok.Type == tokenizer.NewlineToken {
		ts.Take()
		return lookForOuterStatement
	}

	if ts.PrevToken().Type != tokenizer.NewlineToken {
		ts.EmitError(CustomError{
			name:        "Looking in the Wrong place",
			where:       tok.Range,
			description: "I am the lookForStatement function but i am in a place where i am not looking at the start of the line. Usually I look for type, module, import, or declarations so im not sure what to do",
		})
		return nil
	}

	if tok.Type == tokenizer.CommentToken {
		return parseComment
	}

	if tok.Type == tokenizer.IdentifierToken {
		return handleDeclaration
	}

	//If it was something totally unexpected
	if tok.Type != tokenizer.StatementToken {
		ts.EmitError(ExpectedTopLevel{tok.Range})
		ts.Take()
		return lookForOuterStatement
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

func handleDeclaration(ts *TokenSource) stateFn {
	nameToken := ts.Take()
	name := nameToken.Range.String()
	fmt.Print("found ", name)
	//look for a : or an identifier
	// : -> type annotation
	// identifier (hello, var, _) ->  declaration
	for {
		next := ts.Peek()
		if next.Type == tokenizer.SpaceToken || next.Type == tokenizer.CommentToken {
			ts.Take()
			continue
		}

		if next.Type == tokenizer.ColonToken {
			ts.Take() //take the `:`
			fmt.Println("want to parse type annotation")

			return makeTypeAnnotationParser(name)
		}
		if next.Type == tokenizer.AssignmentToken {
			ts.Take() //take the =
			fmt.Println("want to parse a  declaration (function or constant)")
			return nil
		}

		//else we dont know what we're looking at
		ts.Take()
		//panic("add error saying we saw an identifier but now don't know what to do with it")
		ts.EmitError(UnknownIdentifierUsage{
			where: next.Range,
		})
		fmt.Println("Dont know what im doing with this identifier")
		break
	}

	return lookForOuterStatement
}
func makeTypeAnnotationParser(name string) stateFn {
	f := func(ts *TokenSource) stateFn {

		//already took :
		//looking for Type, typeVariable, or {record: typeVariable}

		typeSignature = TypeSignature{}
		for {
			ts.takeUntilNotSpace()
			t := ts.Peek()
			var foundType Type
			//looking for Type, typeVariable, or {record: typeVariable}/{record: Type}, (Type, Type2)
			if t.Type == tokenizer.OpenParenToken {
				foundType = parseTupleType()
			} else if t.Type == tokenizer.OpenCurlyToken {
				foundType = parseRecordType()
			} else if t.Type == tokenizer.IdentifierToken {
				foundType = NamedType{}
			} else {
				//found something invalid
				ts.EmitError()
			}

			//looking for ->
			break
		}

		return lookForOuterStatement
	}
	return f
}

type TypeSignature struct {
	types []Type
}

func (t TypeSignature) ReturnType() Type {
	if len(t.types) < 1 {
		return NoType{}
	}
	return t.types[len(t.types)-1]
}

type typeType int

const (
	UnknownTypeID typeType = iota
	NamedTypeID
	TypeVariableID
	RecordTypeID
	TupleTypeID
)

type Type interface {
	Type() typeType
	Where() tokenizer.SourceRange
}

// denotes an error - a pure function that doesnt return anything just makes heat
type NoType struct{}

func (n NoType) Type() typeType {
	return UnknownTypeID
}
func (n NoType) Where() tokenizer.SourceRange {
	return tokenizer.SourceRange{}
}

type NamedType struct {
	name  string
	where tokenizer.SourceRange
}

func (n NamedType) Type() typeType {
	return NamedTypeID
}
func (n NamedType) Where() tokenizer.SourceRange {
	return n.where
}

func (ts *TokenSource) takeUntilNotSpace() {
	for {
		t := ts.Peek()
		if t.Type == tokenizer.SpaceToken {
			ts.Take()
		} else {
			break
		}
	}
}

func parseComment(ts *TokenSource) stateFn {
	t := ts.Take()
	ts.AddNode(&CommentNode{
		text:  t.Range.String(),
		where: t.Range,
	})
	return lookForOuterStatement
}

func parseModuleDeclaration(ts *TokenSource) stateFn {
	parts := ts.takeUntilNewline()
	where := mergeRangesOfTokens(parts)

	if len(parts) < 1 {
		ts.EmitError(ModuleNeedsName{where})
		return lookForOuterStatement
	}

	shouldBeNameToken := findFirstNonSpace(parts[1:])

	if shouldBeNameToken.Type != tokenizer.IdentifierToken {
		ts.EmitError(ModuleNeedsName{
			where: where,
		})
		fmt.Println("error emitted")
		return lookForOuterStatement
	}

	var name string = shouldBeNameToken.Range.String()

	ts.AddNode(&ModuleDeclaration{
		name:  name,
		where: where,
	})

	return lookForOuterStatement
}

func parseImport(ts *TokenSource) stateFn {
	parts := ts.takeUntilNewline()
	where := mergeRangesOfTokens(parts)

	if len(parts) < 1 {
		ts.EmitError(ImportNeedsName{where})
		return lookForOuterStatement
	}

	shouldBeNameToken := findFirstNonSpace(parts[1:])

	if shouldBeNameToken.Type != tokenizer.IdentifierToken {
		ts.EmitError(ImportNeedsName{
			where: where,
		})
		fmt.Println("error emitted")
		return lookForOuterStatement
	}

	var name string = shouldBeNameToken.Range.String()

	imp := ImportDeclaration{
		name: name,
	}

	ts.AddNode(&imp)

	return lookForOuterStatement
}

func findFirstNonSpace(toks []tokenizer.Token) tokenizer.Token {
	for i, tok := range toks {
		if tok.Type == tokenizer.SpaceToken || tok.Type == tokenizer.CommentToken {
			continue
		}
		return toks[i]
	}
	return tokenizer.Token{Type: tokenizer.UnknownToken}
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

func parseTypeDefinition(ts *TokenSource) stateFn {
	return nil
}

type TokenSource struct {
	tokens     []tokenizer.Token
	tokenIndex int

	program *AST
	errors  tokenizer.ComboError
}

func (ts *TokenSource) AddNode(node TopASTNode) {
	ts.program.toplevel = append(ts.program.toplevel, node)
}

func (ts *TokenSource) EmitError(err error) {
	ts.errors = append(ts.errors, err)
}
func (ts *TokenSource) PrevToken() tokenizer.Token {
	if ts.tokenIndex == 0 {
		return tokenizer.Token{
			Range: tokenizer.SourceRange{},
			Type:  tokenizer.NewlineToken,
		}
	}
	return ts.tokens[ts.tokenIndex-1]
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
		t := ts.Peek()
		toks = append(toks, t)
		if t.Type == tokenizer.NewlineToken {
			break
		}
		ts.Take()
	}
	return toks
}

func (ts *TokenSource) Finished() bool {
	return ts.tokenIndex >= len(ts.tokens)
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
