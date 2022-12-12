package main

const EOFRune = 0x00

type Token struct {
	srange    SourceRange
	tokenType TokenType
}

var oneLengthTokenTypes = map[rune]TokenType{

	' ':  SpaceToken,
	'\t': TabToken,
	'\n': NewlineToken,
	'.':  DotToken,
	',':  CommaToken,
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
