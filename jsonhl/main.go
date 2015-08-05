package main

import (
	"github.com/dawi/jsonhl"
	"os"
)

func main() {
	jsonhl.Highlight(os.Stdin, os.Stdout)
}