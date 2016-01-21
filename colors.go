package jsonhl

import "github.com/dawi/jsont"

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

var resetColor = "\x1b[0m"
