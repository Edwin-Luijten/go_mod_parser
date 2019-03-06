# GO mod parser

This module is extracted from https://github.com/golang/vgo/tree/master/vendor/cmd/go/internal/modfile

## Installation
``` go get github.com/edwin-luijten/go_mod_parser@v0.0 ```  

## Usage

```go
package my_package

import (
	"github.com/edwin-luijten/go_mod_parser"
	"io/ioutil"
	"fmt"
	)

func main() {
    contents, _ := ioutil.ReadFile("./go.mod")
	
    p, err := go_mod_parser.Parse("go.mod", contents, nil)
    if err != nil {
       panic(err)
    }
    
    fmt.Println(p.Module.Mod.Path)
}
```

