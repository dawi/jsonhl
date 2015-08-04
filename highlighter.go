package jsonhl

import (
	"bytes"
	"github.com/dawi/jsont"
	"strings"
)

var keyColor = "\x1b[38;5;33m"
var valueColor = "\x1b[36m"
var resetColor = "\x1b[0m"
var bracketColor = "\x1b[90m"

func Highlight(jsonString string) string {

	var resultBuffer bytes.Buffer

	tokenizer := jsont.NewTokenizer(strings.NewReader(jsonString))

	for tokenizer.Next() {

		token := tokenizer.Token()

		switch token.Type {
		case jsont.ObjectStart, jsont.ObjectEnd:
			resultBuffer.WriteString(bracketColor + token.Value + resetColor)
		case jsont.ArrayStart, jsont.ArrayEnd:
			resultBuffer.WriteString(bracketColor + token.Value + resetColor)
		case jsont.Colon, jsont.Comma:
			resultBuffer.WriteString(bracketColor + token.Value + resetColor)
		case jsont.FieldName:
			resultBuffer.WriteString(keyColor + token.Value + resetColor)
		case jsont.String:
			resultBuffer.WriteString(valueColor + token.Value + resetColor)
		case jsont.Null:
			resultBuffer.WriteString(valueColor + token.Value + resetColor)
		case jsont.True, jsont.False:
			resultBuffer.WriteString(valueColor + token.Value + resetColor)
		}
	}

	return resultBuffer.String()
}
