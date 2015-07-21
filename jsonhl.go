package jsonhl

import (
	"bytes"
)

var keyColor = "\x1b[38;5;33m"
var valueColor = "\x1b[36m"
var resetColor = "\x1b[0m"
var bracketColor = "\x1b[90m"

func HighlightJson(jsonString string) string {

	var resultBuffer bytes.Buffer

	tokens, _ := tokenize(bytes.NewBufferString(jsonString))
	for i := 0; i < len(tokens); i++ {

		token := tokens[i]
		firstByte := token[0]

		if isJsonChar(firstByte) {
			resultBuffer.WriteString(bracketColor + token + resetColor)
		} else if isWhitespace(firstByte) {
			resultBuffer.WriteString(token)
		} else {
			if firstByte == '"' && (tokens[i+1] == ":" || tokens[i+2] == ":") {
				resultBuffer.WriteString(keyColor + token + resetColor)
			} else {
				resultBuffer.WriteString(valueColor + token + resetColor)
			}
		}
	}

	return resultBuffer.String()
}
