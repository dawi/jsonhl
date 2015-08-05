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

func HighlightString(jsonString string) (string, error) {
	b := &bytes.Buffer{}
	err := Highlight(strings.NewReader(jsonString), b)
	return string(b.Bytes()), err
}

func Highlight(reader io.Reader, writer io.Writer) error {

	tokenizer := jsont.NewTokenizer(reader)

	for tokenizer.Next() {

		token := tokenizer.Token()

		var color string
		switch token.Type {
		case jsont.ObjectStart, jsont.ObjectEnd:
			color = bracketColor
		case jsont.ArrayStart, jsont.ArrayEnd:
			color = bracketColor
		case jsont.Colon, jsont.Comma:
			color = bracketColor
		case jsont.FieldName:
			color = keyColor
		case jsont.String:
			color = valueColor
		case jsont.Null:
			color = valueColor
		case jsont.True, jsont.False:
			color = valueColor
		case jsont.Whitespace:
			color = valueColor
		case jsont.Unknown:
			color = valueColor
		}

		if _, err := writer.Write([]byte(color + token.Value + resetColor)); err != nil {
			return err
		}
	}

	return tokenizer.Error()
}
