package jsonhl

import (
	"bufio"
	"bytes"
	"errors"
	"io"
)

// tokenize a json string into json tokens
func tokenize(reader io.Reader) ([]string, error) {
	r := bufio.NewReader(reader)
	var tokens []string
	for {
		if token, err := readToken(r); err != nil || token == "" {
			return tokens, err
		} else {
			tokens = append(tokens, token)
		}
	}
}

// type definition of a read token funtion
type readTokenFunction func(reader *bufio.Reader) (string, error)

// array of functions that are tried in order to read the next token
var readTokenFunctions = []readTokenFunction{readCharToken, readStringToken, readWhitespaceToken, readNonStringToken}

// reads the next json token from reader
func readToken(reader *bufio.Reader) (string, error) {
	for _, readTokenFunction := range readTokenFunctions {
		if token, err := readTokenFunction(reader); err != nil || token != "" {
			return handleEOF(token, err)
		}
	}
	return "", errors.New("Error while parsing tokens. This error should never happen.")
}

// reads a single character token like { } [ ] : ,
func readCharToken(reader *bufio.Reader) (string, error) {
	for {
		if b, err := reader.ReadByte(); err != nil {
			return "", err
		} else if b == '{' || b == '}' || b == '[' || b == ']' || b == ':' || b == ',' {
			return string(b), nil
		} else {
			reader.UnreadByte()
			return "", nil
		}
	}
}

// reads a string token like "abc"
func readStringToken(reader *bufio.Reader) (string, error) {

	var escape = false

	var buffer bytes.Buffer

	b, err := reader.ReadByte()
	if err != nil {
		return "", err
	}

	if b != '"' {
		reader.UnreadByte()
		return "", nil
	}

	buffer.WriteByte(b)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			return "", err
		}
		if escape {
			escape = false
			buffer.WriteByte(b)
		} else if b == '\\' {
			escape = true
			buffer.WriteByte(b)
		} else if b == '"' {
			escape = false
			buffer.WriteByte(b)
			return buffer.String(), nil
		} else {
			escape = false
			buffer.WriteByte(b)
		}
	}
}

// reads a whitespace token like " \r \n \t "
func readWhitespaceToken(reader *bufio.Reader) (string, error) {
	var buffer bytes.Buffer
	for {
		if b, err := reader.ReadByte(); err != nil {
			return buffer.String(), err
		} else if isWhitespace(b) {
			buffer.WriteByte(b)
		} else {
			reader.UnreadByte()
			return buffer.String(), nil
		}
	}
}

// reads a non string token like true, false, 123.45
func readNonStringToken(reader *bufio.Reader) (string, error) {
	var buffer bytes.Buffer
	for {
		if b, err := reader.ReadByte(); err != nil {
			return buffer.String(), err
		} else if isWhitespace(b) || isJsonChar(b) {
			reader.UnreadByte()
			return buffer.String(), nil
		} else {
			buffer.WriteByte(b)
		}
	}
}

// true if byte represents a whitespace
func isWhitespace(b byte) bool {
	return b == ' ' || b == '\r' || b == '\n' || b == '\t'
}

// true if byte represents a json character
func isJsonChar(b byte) bool {
	return b == ':' || b == ',' || b == '{' || b == '}' || b == '[' || b == ']'
}

// ignores EOF errors
func handleEOF(token string, err error) (string, error) {
	if err == nil || err == io.EOF {
		return token, nil
	}
	return token, err
}
