<p align="center"><img src="https://raw.githubusercontent.com/go-composites/brand/main/social/go-composites.png" alt="go-composites/boolean" width="720"></p>

# Boolean

[![ci](https://github.com/go-composites/boolean/actions/workflows/ci.yml/badge.svg)](https://github.com/go-composites/boolean/actions/workflows/ci.yml)

The truth-value composite of [go-composites](https://github.com/go-composites).
A `Boolean` wraps a Go `bool` and exposes its logic algebraically: the logical
operators return another `Boolean.Interface`, so expressions compose and never
go `nil`.

## Install

```sh
go get github.com/go-composites/boolean
```

## Usage

```golang
package main

import (
    "fmt"

    Boolean "github.com/go-composites/boolean/src"
)

func main() {
    t := Boolean.True()
    f := Boolean.False()

    fmt.Println(t.Inspect())                 // <Boolean:0x... value=true>
    fmt.Println(t.And(f).IsTrue())           // false
    fmt.Println(t.Or(f).IsTrue())            // true
    fmt.Println(t.Xor(f).IsTrue())           // true
    fmt.Println(t.Not().IsFalse())           // true
    fmt.Println(t.Equal(f).ToGoString())     // "false"
}
```

## API

### Constructors

- `New(b bool) Interface` — from a Go `bool`.
- `True() Interface` / `False() Interface` — the two constants.
- `Null() Interface` — the Null-Object `Boolean`.

### Conversion & predicates (return plain Go values)

| method | returns | notes |
| --- | --- | --- |
| `ToGoBool()` | Go `bool` | the underlying value |
| `ToGoString()` | Go `string` | `"true"` / `"false"` (`"null"` for the Null-Object) |
| `IsTrue()` | Go `bool` | value is `true` |
| `IsFalse()` | Go `bool` | value is `false` |
| `IsNull()` | Go `bool` | `true` only for the Null-Object |

### Logic (each returns a `Boolean.Interface`)

| method | meaning |
| --- | --- |
| `And(other)` | logical AND |
| `Or(other)` | logical OR |
| `Xor(other)` | exclusive OR |
| `Not()` | negation |
| `Equal(other)` | equality of the two truth values |

### Inspection

- `Inspect() Inspect.Interface` — a one-line `<Boolean:0x... value=...>` view.

## Null-Object

`Null()` returns the never-nil Null-Object `Boolean`: it is neither true nor
false (`IsTrue()` and `IsFalse()` are both `false`), `ToGoString()` renders
`"null"`, and every logical operator returns the null `Boolean`. `IsNull()`
reports `true` for it.

## License

BSD-3-Clause © the go-composites/boolean authors.
