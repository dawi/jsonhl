package jsonhl

import (
	"bufio"
	"bytes"
	"testing"
)

func TestReadToken_string_eof(t *testing.T) {
	reader := bufio.NewReader(bytes.NewBufferString(`"string"`))
	if out, _ := readToken(reader); out != `"string"` {
		t.Errorf("ERROR: '%s'", out)
	}
}

func TestReadToken_string_whitespace(t *testing.T) {
	reader := bufio.NewReader(bytes.NewBufferString(`"string" `))
	if out, _ := readToken(reader); out != `"string"` {
		t.Errorf("ERROR: '%s'", out)
	}
}

func TestReadToken_string_with_whitespace_eof(t *testing.T) {
	reader := bufio.NewReader(bytes.NewBufferString(`"aaa \r\n\t bbb"`))
	if out, _ := readToken(reader); out != `"aaa \r\n\t bbb"` {
		t.Errorf("ERROR: '%s'", out)
	}
}

func TestReadToken_string_with_whitespace_whitespace(t *testing.T) {
	reader := bufio.NewReader(bytes.NewBufferString(`"aaa \r\n\t bbb" `))
	if out, _ := readToken(reader); out != `"aaa \r\n\t bbb"` {
		t.Errorf("ERROR: '%s'", out)
	}
}

func TestReadToken_boolean_eof(t *testing.T) {
	reader := bufio.NewReader(bytes.NewBufferString("boolean"))
	if out, _ := readToken(reader); out != "boolean" {
		t.Errorf("ERROR: '%s'", out)
	}
}

func TestReadToken_boolean_whitespace(t *testing.T) {
	reader := bufio.NewReader(bytes.NewBufferString("boolean "))
	if out, _ := readToken(reader); out != "boolean" {
		t.Errorf("ERROR: '%s'", out)
	}
}

func TestReadToken_number_eof(t *testing.T) {
	reader := bufio.NewReader(bytes.NewBufferString("11.99"))
	if out, _ := readToken(reader); out != "11.99" {
		t.Errorf("ERROR: '%s'", out)
	}
}

func TestReadToken_number_whitespace(t *testing.T) {
	reader := bufio.NewReader(bytes.NewBufferString("11.99 "))
	if out, _ := readToken(reader); out != "11.99" {
		t.Errorf("ERROR: '%s'", out)
	}
}
