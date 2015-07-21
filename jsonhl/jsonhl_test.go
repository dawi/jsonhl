package jsonhl

import (
	"strings"
	"testing"
)

var colorReplacer = strings.NewReplacer("kc", keyColor, "vc", valueColor, "gc", bracketColor, "rc", resetColor)

func TestHighlightJson(t *testing.T) {

	input := `{"key1":"value1","key2":value2}`

	expected := colorReplacer.Replace(
		`gc{rckc"key1"rcgc:rcvc"value1"rcgc,rckc"key2"rcgc:rcvcvalue2rcgc}rc`)

	if result := HighlightJson(input); result != expected {
		t.Error("ERROR: " + result)
	}
}
