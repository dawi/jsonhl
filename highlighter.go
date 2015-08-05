package jsonhl

import (
	"bytes"
	"github.com/dawi/jsont"
	"io"
	"strings"
)

var keyColor = "\x1b[38;5;33m"
var valueColor = "\x1b[36m"
var resetColor = "\x1b[0m"
var bracketColor = "\x1b[90m"

func HighlightString(jsonString string) string {
	b := &bytes.Buffer{}
	Highlight(strings.NewReader(jsonString), b)
	return string(b.Bytes())
}

func Highlight(reader io.Reader, writer io.Writer) {

	tokenizer := jsont.NewTokenizer(reader)

	for tokenizer.Next() {

		token := tokenizer.Token()

		switch token.Type {
		case jsont.ObjectStart, jsont.ObjectEnd:
			writer.Write([]byte(bracketColor + token.Value + resetColor))
		case jsont.ArrayStart, jsont.ArrayEnd:
			writer.Write([]byte(bracketColor + token.Value + resetColor))
		case jsont.Colon, jsont.Comma:
			writer.Write([]byte(bracketColor + token.Value + resetColor))
		case jsont.FieldName:
			writer.Write([]byte(keyColor + token.Value + resetColor))
		case jsont.String:
			writer.Write([]byte(valueColor + token.Value + resetColor))
		case jsont.Null:
			writer.Write([]byte(valueColor + token.Value + resetColor))
		case jsont.True, jsont.False:
			writer.Write([]byte(valueColor + token.Value + resetColor))
		case jsont.Whitespace:
			writer.Write([]byte(valueColor + token.Value + resetColor))
		case jsont.Unknown:
			writer.Write([]byte(valueColor + token.Value + resetColor))
		}
	}
}
