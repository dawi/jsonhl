package jsonhl

import (
	"bytes"
	"github.com/dawi/jsont"
	"io"
	"strings"
)

func HighlightString(jsonString string) (string, error) {
	return HighlightStringC(jsonString, DefaultColors)
}

func HighlightStringC(jsonString string, colors Colors) (string, error) {
	reader := strings.NewReader(jsonString)
	writer := &bytes.Buffer{}
	err := HighlightC(reader, writer, colors)
	return writer.String(), err
}

func HighlightBytes(jsonString []byte, colors Colors) ([]byte, error) {
	return HighlightBytesC(jsonString, DefaultColors)
}

func HighlightBytesC(jsonString []byte, colors Colors) ([]byte, error) {
	reader := bytes.NewBuffer(jsonString)
	writer := &bytes.Buffer{}
	err := HighlightC(reader, writer, colors)
	return writer.Bytes(), err
}

func Highlight(reader io.Reader, writer io.Writer) error {
	return HighlightC(reader, writer, DefaultColors)
}

func HighlightC(reader io.Reader, writer io.Writer, colors Colors) error {
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
