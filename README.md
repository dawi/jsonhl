## jsonhl 

Go library to colorize json output of terminal applications.

## Features

* small and focused
* colorizes json output without changing the json formatting

## Usage

Using io.Reader and io.Writer:

```
jsonhl.Highlight(os.Stdin, os.Stdout)
jsonhl.HighlightC(os.Stdin, os.Stdout, jsonhl.DefaultColors)
```

Using strings:

```
fmt.Println(jsonhl.HighlightString(`{ "hello" : "world" }`))
fmt.Println(jsonhl.HighlightStringC(`{ "hello" : "world" }`, jsonhl.DefaultColors))
```
