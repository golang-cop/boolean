<p align="center"><img src="https://raw.githubusercontent.com/golang-cop/brand/main/social/golang-oop.png" alt="golang-cop/boolean" width="720"></p>

# Boolean

## Usage

> [!NOTE] main.go

```golang
package main

import (
    "fmt"
    Boolean "github.com/golang-oop/boolean/src
)

func main() {
    bool := Boolean.True()
    fmt.Println(bool.Inspect())
    fmt.Println(Boolean.True().Equal(Boolean.False()).Inspect())
}
```
```bash
$ task build
task: [build] go build -o bin/boolean
task: [build] ./bin/boolean
<Boolean:0x14000110012 value=true>
Boolean.True().Equal(Boolean.False()).Inspect():  <Boolean:0x14000110018 value=false>
```

## Useful links

[GitHub Flavored Markdown Spec](https://github.github.com/gfm/)  
[Dynamic Method Invocation](https://medium.com/@ansujain/dynamic-method-invocation-using-reflection-in-go-a-comprehensive-guide-f2457d025964)