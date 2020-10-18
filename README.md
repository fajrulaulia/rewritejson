# package rewritejson
Package to generate to json for easy write something


## How to use
Package strconv implements conversions to and from array string representations of JSON.

> func rewritejson.Response(data []string) []byte

`import "github.com/fajrulaulia/rewritejson"`

### Standard Definition
this how to convert payload to json
``` go
defined:=[]string{"code:number|500", "message:string|This Error"}
output := core.Response(defined)
```

## Author
- fajrul aulia
