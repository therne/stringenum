stringenum
=========

Generates enum codes for types aliased `string`.

 * Supports JSON, YAML serialization
 * Implements `Validator` interface
 * Does not support default values yet, but it has good fit with 
 third-party default value modules like [creasty/defaults](https://github.com/creasty/defaults)
 because it's basically a string.
 
 
 
## Installation

You need to install `stringenum` to generate enum stub codes.

```
 $ go get github.com/
```

## Usage

On the top of your type definition sources, add `go generate` clause to generate stub codes with `stringenum`.

```go
// go:generate stringenum Kind
package mytype

type Kind string

const (
    Apple  = Kind("apple")
    Google = Kind("google")
)
```

Then, run go generate to generate stub codes:

```
 $ go generate ./...
```


## License: MIT