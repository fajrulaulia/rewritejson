# package rewritejson
Package to generate to json for easy write something


## How to use
Package strconv implements conversions to and from array string representations of JSON.\
add package into you project :\
`go get github.com/fajrulaulia/rewritejson`

import :\
`import c "github.com/fajrulaulia/rewritejson"`

### Standard Definition
> func rewritejson.Response(data []string) []byte

### Implementation in real-life
``` go
package ..
..............
`import c "github.com/fajrulaulia/rewritejson"`
...............
...............
defined:=[]string{"code:number|500", "message:string|This Error"}
output := c.Response(defined)
```

## Author
- fajrul aulia
