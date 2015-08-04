package main

import (
	"fmt"
	"github.com/dawi/jsonhl"
)

func main() {
	fmt.Println(jsonhl.Highlight(` { "hello" : "world" } `))
}