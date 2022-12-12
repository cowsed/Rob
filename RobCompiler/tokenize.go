package main

import (
	"fmt"
	"strings"
)

type MegaError []error

func (me MegaError) Error() string {
	s := ""
	for _, e := range me {
		s += e.Error()
		s += "\n"
	}
	return s
}

type SourceRange struct {
	start, end   int
	source_runes *[]rune
}

func (sr SourceRange) String() string {
	return string((*sr.source_runes)[sr.start:sr.end])
}

// A way to represent a source file and peek, take runes from it
type TokenizingSource struct {
	source            string
	source_identifier string
	source_runes      []rune

	source_index      int
	end_of_last_token int
	tokens            []Token
	errors            []error
}

func (ts *TokenizingSource) Setup() {
	ts.tokens = []Token{}
	ts.source_runes = []rune(ts.source)
	ts.errors = []error{}
}

func (ts *TokenizingSource) EmitError(err error) {
	ts.errors = append(ts.errors, err)
}
func (ts *TokenizingSource) EmitToken(tokenType TokenType) {
	ts.tokens = append(ts.tokens, Token{
		srange: SourceRange{
			start:        ts.end_of_last_token,
			end:          ts.source_index,
			source_runes: &ts.source_runes,
		},
		tokenType: tokenType,
	})
	text := ts.source_runes[ts.end_of_last_token:ts.source_index]
	fmt.Printf("Emitted%s: `%s`\n", tokenType.String(), string(text))
	ts.end_of_last_token = ts.source_index

	fmt.Println("---- Tokens: ", ts.tokens)
}

func (ts *TokenizingSource) Take() rune {
	ts.source_index++
	if ts.source_index > len(ts.source_runes) {
		return EOFRune
	}
	return ts.source_runes[ts.source_index-1]
}
func (ts *TokenizingSource) Peek() rune {
	if ts.source_index >= len(ts.source_runes) {
		return EOFRune
	}
	return ts.source_runes[ts.source_index]
}

func (ts *TokenizingSource) Return() {
	ts.source_index--
	if ts.source_index < 0 {
		ts.source_index = 0
	}
}

func (ts *TokenizingSource) Finished() bool {
	return ts.end_of_last_token >= len(ts.source_runes)
}

const EOFRune = 0x00

type TokenType int

const (
	UnknownToken = iota
	SpaceToken
	NewlineToken
	TabToken
	CommaToken
	DotToken
	OpenParenToken
	CloseParenToken
	OpenCurlyToken
	CloseCurlyToken
	OpenSquareToken
	CloseSquareToken

	AsignmentToken
	PlusToken
	MinusToken
	StarToken
	DivideToken
	LessThanToken
	GreaterThanToken

	EqualsToken
	GreaterThanEqToken
	LessThanEqToken
	RightArrowToken

	CommentToken
	IdentifierToken
	NumberToken
	StringToken
)

func (tt TokenType) String() string {
	return map[TokenType]string{
		UnknownToken:       "UnknownToken",
		SpaceToken:         "SpaceToken",
		NewlineToken:       "NewlineToken",
		TabToken:           "TabToken",
		CommaToken:         "CommaToken",
		DotToken:           "DotToken",
		OpenParenToken:     "OpenParenToken",
		CloseParenToken:    "CloseParenToken",
		OpenCurlyToken:     "OpenCurlyToken",
		CloseCurlyToken:    "CloseCurlyToken",
		OpenSquareToken:    "OpenSquareToken",
		CloseSquareToken:   "CloseSquareToken",
		AsignmentToken:     "AsignmentToken",
		PlusToken:          "PlusToken",
		MinusToken:         "MinusToken",
		StarToken:          "StarToken",
		DivideToken:        "DivideToken",
		LessThanToken:      "LessThanToken",
		GreaterThanToken:   "GreaterThanToken",
		EqualsToken:        "EqualsToken",
		GreaterThanEqToken: "GreaterThanEqToken",
		LessThanEqToken:    "LessThanEqToken",
		CommentToken:       "CommentToken",
		IdentifierToken:    "IdentifierToken",
		NumberToken:        "NumberToken",
		StringToken:        "StringToken",
		RightArrowToken:    "RightArrowToken",
	}[tt]
}

type Token struct {
	srange    SourceRange
	tokenType TokenType
}

func (t Token) String() string {
	str_range := t.srange.String()
	if str_range == "\n" {
		str_range = "\\n"
	} else if str_range == "\t" {
		str_range = "\t"
	}
	return fmt.Sprintf("{%s: `%s`}", t.tokenType.String(), str_range)
}

var oneLengthTokenTypes = map[rune]TokenType{

	' ':  SpaceToken,
	'\t': TabToken,
	'\n': NewlineToken,
	',':  DotToken,
	'.':  CommaToken,
	'(':  OpenParenToken,
	')':  CloseParenToken,
	'[':  OpenSquareToken,
	']':  CloseSquareToken,
	'{':  OpenCurlyToken,
	'}':  CloseCurlyToken,
	'=':  AsignmentToken,
	'+':  PlusToken,
	'-':  MinusToken,
	'*':  StarToken,
	'/':  DivideToken,
	'<':  LessThanToken,
	'>':  GreaterThanToken,
}

var twoLengthTokenTypes = map[string]TokenType{
	"==": EqualsToken,
	">=": GreaterThanEqToken,
	"<=": LessThanEqToken,
	"->": RightArrowToken,
	"--": CommentToken,
	"{-": CommentToken,
}

var commentStart = "--"
var multilineCommentStart = "{-"
var multilineCommentEnd = "-}"

var allowedIdentifierStarts = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"
var allowedNumberStarts = ".0123456789"

func Tokenize(source_text string) ([]Token, error) {
	source := TokenizingSource{
		source:            source_text,
		source_identifier: "filename/",
	}
	source.Setup()

	var tokenizer Tokenizer = &SearchForSomething{
		tokSource: &source,
	}

	for !source.Finished() {
		fmt.Println("finished", source.Finished(), source.source_index, len(source.source_runes))
		tokenizer = tokenizer.tokenize(nil)
		if tokenizer == nil {
			fmt.Println("Not looking for anything particular, searching for something")
			tokenizer = &SearchForSomething{
				tokSource: &source,
			}
		}
	}
	var err error
	if len(source.errors) == 0 {
		err = nil
	} else {
		err = MegaError(source.errors)
	}
	return source.tokens, err
}

type Tokenizer interface {
	tokenize(previous Tokenizer) Tokenizer
}

type SearchForSomething struct {
	tokSource *TokenizingSource
}

func (sfs *SearchForSomething) tokenize(previous Tokenizer) Tokenizer {
	next := sfs.tokSource.Peek()
	fmt.Println("next", next)

	startOf2LengthOperator := isFirstLetterKeyOf(next, twoLengthTokenTypes)
	oneLengthOperator := isKeyOf(next, oneLengthTokenTypes)

	if oneLengthOperator || startOf2LengthOperator {
		fmt.Println("was operator", string(next))
		return &MakeOperator{
			tokSource: sfs.tokSource,
		}
	}

	if next == '"' {
		return &MakeStringLiteral{
			tokSource: sfs.tokSource,
		}
	}

	if strings.ContainsRune(allowedNumberStarts, next) {
		return &MakeNumberLiteral{tokSource: sfs.tokSource}
	}

	if strings.ContainsRune(allowedIdentifierStarts, next) {
		return &MakeIdentifer{
			tokSource: sfs.tokSource,
		}
	}

	fmt.Println("Unknown", string(next))

	return sfs
}

type MakeStringLiteral struct {
	tokSource *TokenizingSource
}

func (ms *MakeStringLiteral) tokenize(previus Tokenizer) Tokenizer {
	ms.tokSource.Take() // first "
	for ms.tokSource.Take() != '"' {

	}
	ms.tokSource.EmitToken(StringToken)
	return nil
}

type MakeNumberLiteral struct {
	tokSource *TokenizingSource
}

func (mm *MakeNumberLiteral) tokenize(previus Tokenizer) Tokenizer {
	for !willSeparate(mm.tokSource.Take()) {
	}
	mm.tokSource.Return()
	mm.tokSource.EmitToken(NumberToken)

	return nil
}

type MakeIdentifer struct {
	tokSource *TokenizingSource
}

func willSeparate(r rune) bool {
	for op := range oneLengthTokenTypes {
		if r == op {
			return true
		}
	}
	return false
}

func (mi *MakeIdentifer) tokenize(previus Tokenizer) Tokenizer {
	for !willSeparate(mi.tokSource.Take()) {
	}
	mi.tokSource.Return()
	mi.tokSource.EmitToken(IdentifierToken)
	return nil
}

type MakeOperator struct {
	tokSource *TokenizingSource
}

func (mo *MakeOperator) tokenize(previous Tokenizer) Tokenizer {
	firstRune := mo.tokSource.Take()

	if isFirstLetterKeyOf(firstRune, twoLengthTokenTypes) {
		secondRune := mo.tokSource.Take()
		operator := string(firstRune) + string(secondRune)

		_, makes2LengthOperator := twoLengthTokenTypes[operator]

		if operator == commentStart {
			mo.tokSource.Return()
			mo.tokSource.Return()
			return &MakeComment{tokSource: mo.tokSource}
		} else if operator == multilineCommentStart {
			mo.tokSource.Return()
			mo.tokSource.Return()
			return &MakeMultilineComment{tokSource: mo.tokSource}
		}

		if makes2LengthOperator {
			operatorType := twoLengthTokenTypes[operator]
			mo.tokSource.EmitToken(operatorType)
			return nil
		}
	}

	oneLengthType := oneLengthTokenTypes[firstRune]

	mo.tokSource.EmitToken(oneLengthType)

	return nil
}

type MakeComment struct {
	tokSource *TokenizingSource
}

func (mc *MakeComment) tokenize(previous Tokenizer) Tokenizer {
	for !mc.tokSource.Finished() {
		r := mc.tokSource.Take()
		if r == '\n' {
			mc.tokSource.Return()
			break
		}
	}
	mc.tokSource.EmitToken(CommentToken)
	return previous
}

type MakeMultilineComment struct {
	tokSource *TokenizingSource
}

func (mc *MakeMultilineComment) tokenize(previous Tokenizer) Tokenizer {
	mc.tokSource.Take()
	mc.tokSource.Take()
	depth := 1
	var last rune
	for !mc.tokSource.Finished() && depth > 0 {
		r := mc.tokSource.Take()
		if string(last)+string(r) == multilineCommentStart {
			depth++
		}
		if string(last)+string(r) == multilineCommentEnd {
			depth--
			break
		}
		last = r
	}
	mc.tokSource.EmitToken(CommentToken)
	return nil
}

func isKeyOf(s rune, source map[rune]TokenType) bool {
	_, ok := source[s]
	return ok
}

func isFirstLetterKeyOf(s rune, source map[string]TokenType) bool {
	for k := range source {
		if []rune(k)[0] == s {
			return true
		}
	}
	return false
}
