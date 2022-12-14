package tokenizer

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

var red = color.New(color.FgRed).SprintFunc()
var gray = color.New(color.FgHiBlack).SprintFunc()

var sourceToFilename = map[*[]rune]string{}

type ComboError []error

func (me ComboError) Error() string {
	s := ""
	for _, e := range me {
		s += e.Error()
		s += "\n"
	}
	return s
}

type UnknownCharacter struct {
	where SourceRange
}

func (uc *UnknownCharacter) Error() string {
	message := "I honestly don't know what this character is. If it should be allowed as part of a name please contact me. If not, perhaps you meant to enclose it as a \"string\"?\n"

	errStr := ErrorTemplate("Unknown Character", uc.where, message)

	return errStr
}

type UnclosedStringLiteral struct {
	where SourceRange
}

func (usl UnclosedStringLiteral) Error() string {

	message := ErrorTemplate("Unclosed String Literal", usl.where, "")

	return message
}

func ErrorTemplate(errorName string, srange SourceRange, afterMessage string) string {
	filename := sourceToFilename[srange.SourceRunes]
	if filename == "" {
		filename = "--filename unknown--"
	}
	message := red(stringAndBorder("- "+errorName, 60) + "\n\n")
	if srange.SourceRunes != nil {
		line := srange.GetLine()
		lineNum := srange.LineNumber()
		lineNumStr := fmt.Sprintf("%d|", lineNum)
		message += fmt.Sprintf("%s%s", lineNumStr, line) + "\n"

		untilStart := srange.Start - srange.GetStartOfLine()

		message += strings.Repeat(" ", len(lineNumStr)) + strings.Repeat(" ", untilStart) + red(strings.Repeat("^", srange.Length())+"\n")
	}
	message += gray("at " + filename + "\n")

	message += afterMessage + "\n"

	return message

}

func stringAndBorder(s string, width int) string {
	msg := s
	msg += strings.Repeat("-", max(width-len(s), 0))
	return msg
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
