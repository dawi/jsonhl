# jsonhl

Json syntax highlighting for terminal applications.

## Fetaures

* small library
* that adds color to your json output 
* without changing its formatting

## Usage

```
fmt.Println(jsonhl.HighlightJson(` { "hello" : "world" } `))
```

## Todo

* tests
* configurable colors
* optional reformatting (maybe)
