package jsonhl

import (
	"bytes"
	"github.com/dawi/jsont"
	"io"
	"strings"
)

var resetColor = "\x1b[0m"

type Colors map[jsont.TokenType]string

var DefaultColors = Colors{
	jsont.ObjectStart: "\x1b[38;5;242m",
	jsont.ObjectEnd:   "\x1b[38;5;242m",
	jsont.ArrayStart:  "\x1b[38;5;242m",
	jsont.ArrayEnd:    "\x1b[38;5;242m",
	jsont.Colon:       "\x1b[38;5;242m",
	jsont.Comma:       "\x1b[38;5;242m",
	jsont.FieldName:   "\x1b[38;5;33m",
	jsont.True:        "\x1b[38;5;22m",
	jsont.False:       "\x1b[38;5;124m",
	jsont.Null:        "\x1b[38;5;124m",
	jsont.Integer:     "\x1b[38;5;117m",
	jsont.Float:       "\x1b[38;5;117m",
	jsont.String:      "\x1b[38;5;45m",
	jsont.Unknown:     "\x1b[38;5;1m",
}

func HighlightString(jsonString string) (string, error) {
	b := &bytes.Buffer{}
	err := Highlight(strings.NewReader(jsonString), b)
	return string(b.Bytes()), err
}

func HighlightBytes(jsonString []byte) ([]byte, error) {
	b := &bytes.Buffer{}
	err := Highlight(strings.NewReader(string(jsonString)), b)
	return b.Bytes(), err
}

func Highlight(reader io.Reader, writer io.Writer) error {

	tokenizer := jsont.NewTokenizer(reader)

	for tokenizer.Next() {
		token := tokenizer.Token()
		color := DefaultColors[token.Type]
		if _, err := writer.Write([]byte(color + token.Value + resetColor)); err != nil {
			return err
		}
	}

	return tokenizer.Error()
}
