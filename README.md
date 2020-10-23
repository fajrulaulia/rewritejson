# package rewritejson
Package to generate to json for easy write something


## Installation
use command `go` for get package.
`go get github.com/fajrulaulia/rewritejson`

import : \
`import c "github.com/fajrulaulia/rewritejson"`

## Usage
this is a function structure how to use this package.\
> func rewritejson.Response(data []string) []byte

## Example
``` go
package ......
..............
import c "github.com/fajrulaulia/rewritejson"
...............
...............
defined:=[]string{"code:number|500", "message:string|This Error"}
output := c.Response(defined)
```

## Author
- fajrul aulia
