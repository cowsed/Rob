package tokenizer

import (
	"fmt"
	"strings"
)

type SourceRange struct {
	Start, End  int
	SourceRunes *[]rune
}

func (sr SourceRange) Length() int {
	return sr.End - sr.Start
}
func (sr SourceRange) GetStartOfLine() int {
	for i := sr.Start - 1; i > 0; i-- {
		if (*sr.SourceRunes)[i] == '\n' {
			return i + 1
		}
	}
	return 0
}

func (sr SourceRange) GetLine() string {

	lineStart := sr.Start
	if (*sr.SourceRunes)[sr.Start] == '\n' {
		lineStart--
	}
	//find current line
	for i := lineStart; i >= 0; i-- {
		lineStart = i
		if (*sr.SourceRunes)[i] == '\n' {
			lineStart++
			break
		}
	}

	lineEnd := sr.End - 1
	for i := sr.End - 1; i < len(*sr.SourceRunes); i++ {
		lineEnd = i
		fmt.Println("finding end", string((*sr.SourceRunes)[i]))
		if (*sr.SourceRunes)[i] == '\n' {
			break
		}
	}
	currentLine := (*sr.SourceRunes)[lineStart:lineEnd]

	return string(currentLine)
}
func (sr SourceRange) LineNumber() int {
	return strings.Count(string((*sr.SourceRunes)[:sr.Start]), "\n") + 1
}

func (sr SourceRange) String() string {
	return string((*sr.SourceRunes)[sr.Start:sr.End])
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

func (ts *TokenizingSource) currentRange() SourceRange {
	return SourceRange{
		Start:       ts.end_of_last_token,
		End:         ts.source_index,
		SourceRunes: &ts.source_runes,
	}
}

func (ts *TokenizingSource) EmitToken(tokenType TokenType) {
	ts.tokens = append(ts.tokens, Token{
		Range: ts.currentRange(),
		Type:  tokenType,
	})

	ts.end_of_last_token = ts.source_index

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
	return ts.source_index >= len(ts.source_runes)
}

func (t Token) String() string {
	str_range := t.Range.String()
	if str_range == "\n" {
		str_range = "\\n"
	} else if str_range == "\t" {
		str_range = "\t"
	}
	return fmt.Sprintf("{%s: `%s`}", t.Type.String(), str_range)
}

func Tokenize(source_text string) ([]Token, error) {
	source := TokenizingSource{
		source:            source_text,
		source_identifier: "filename.rob",
	}
	sourceToFilename[&source.source_runes] = source.source_identifier
	source.Setup()

	var tokenizer Tokenizer = &SearchForSomething{
		tokSource: &source,
	}

	for !source.Finished() {
		tokenizer = tokenizer.tokenize(nil)
		if tokenizer == nil {
			tokenizer = &SearchForSomething{
				tokSource: &source,
			}
		}
	}
	source.EmitToken(EOFToken)
	var err error
	if len(source.errors) == 0 {
		err = nil
	} else {
		err = ComboError(source.errors)
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

	startOf2LengthOperator := isFirstLetterKeyOf(next, twoLengthTokenTypes)
	oneLengthOperator := isKeyOf(next, oneLengthTokenTypes)

	if oneLengthOperator || startOf2LengthOperator {
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

	//We've got no idea what this is, thats an error

	sfs.tokSource.Take()
	sfs.tokSource.EmitError(&UnknownCharacter{
		where: sfs.tokSource.currentRange(),
	})
	sfs.tokSource.EmitToken(UnknownToken)

	return sfs
}

type MakeStringLiteral struct {
	tokSource *TokenizingSource
}

func (ms *MakeStringLiteral) tokenize(previus Tokenizer) Tokenizer {

	ms.tokSource.Take() // first "
	for {
		r := ms.tokSource.Take()
		if r == '"' {
			break
		}

		if r == '\n' {
			ms.tokSource.EmitError(UnclosedStringLiteral{
				where: ms.tokSource.currentRange(),
			})
			ms.tokSource.Return()
			break
		}
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
	for op := range twoLengthTokenTypes {
		if r == []rune(op)[0] {
			return true
		}
	}
	return false
}

func (mi *MakeIdentifer) tokenize(previus Tokenizer) Tokenizer {
	for !willSeparate(mi.tokSource.Take()) {
	}
	mi.tokSource.Return()

	text := mi.tokSource.currentRange().String()
	if isOuterKeyWord(text) {
		mi.tokSource.EmitToken(StatementToken)
	} else {
		mi.tokSource.EmitToken(IdentifierToken)
	}
	return nil
}

func isOuterKeyWord(t string) bool {
	outerKeywordSet := map[string]bool{
		"type":   true,
		"alias":  true,
		"module": true,
		"import": true,
	}
	return outerKeywordSet[t]
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
