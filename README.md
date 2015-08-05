## jsonhl 

Go library to colorize json output of terminal applications.

## Features

* small and focused go library
* adds color to your json output
* without changing the json formatting

## Usage

Using io.Reader and io.Writer:

```
jsonhl.Highlight(os.Stdin, os.Stdout)
```

Using strings:

```
fmt.Println(jsonhl.HighlightString(` { "hello" : "world" } `))
```

## Todos

* tests
* configurable colors
