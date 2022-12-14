package tokenizer

const EOFRune = 0x00

type Token struct {
	Range SourceRange
	Type  TokenType
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
	'^':  RaisedToken,
	'<':  LessThanToken,
	'>':  GreaterThanToken,
	':':  ColonToken,
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
	EOFToken

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

	ColonToken
	AsignmentToken
	PlusToken
	MinusToken
	StarToken
	DivideToken
	RaisedToken

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

	StatementToken
)

func (tt TokenType) String() string {
	return map[TokenType]string{
		UnknownToken:       "UnknownToken",
		EOFToken:           "EOFToken",
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
		RaisedToken:        "RaisedToken",
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
		StatementToken:     "StatementToken",
		ColonToken:         "ColonToken",
	}[tt]
}
