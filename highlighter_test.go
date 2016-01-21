package jsonhl

import (
	"strings"
	"testing"
	"github.com/dawi/jsont"
)

var colorReplacer = strings.NewReplacer(
	"kc", DefaultColors[jsont.FieldName],
	"vc", DefaultColors[jsont.String],
	"gc", DefaultColors[jsont.ObjectStart],
	"rc", resetColor)

func TestHighlightJson(t *testing.T) {

	validate(
		`{"hello":"world"}`,
		`gc{rckc"hello"rcgc:rcvc"world"rcgc}rc`, t)

	validate(
		`{"key1":"value1","key2":"value2"}`,
		`gc{rckc"key1"rcgc:rcvc"value1"rcgc,rckc"key2"rcgc:rcvc"value2"rcgc}rc`, t)

}

func validate(input string, expected string, t *testing.T) {
	expected = colorReplacer.Replace(expected)
	if result, _ := HighlightString(input); result != expected {
		t.Error("ERROR: " + result)
	}
}
