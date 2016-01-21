## jsonhl 

Go library to colorize JSON output of terminal applications.

## Features

* small and focused
* colorizes JSON output without changing the JSON formatting

## Usage

Using io.Reader and io.Writer:

```go
jsonhl.Highlight(os.Stdin, os.Stdout)
jsonhl.HighlightC(os.Stdin, os.Stdout, jsonhl.DefaultColors)
```

Using strings:

```go
fmt.Println(jsonhl.HighlightString(`{ "hello" : "world" }`))
fmt.Println(jsonhl.HighlightStringC(`{ "hello" : "world" }`, jsonhl.DefaultColors))
```
