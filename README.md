stringenum
=========

A go tool to auto-generate serialization / validation methods for enum types aliasing `string`.

#### Features

 * Supports JSON, YAML serialization
 * Implements `Validator` interface
 * **Does not support default values yet.** But it has good fit with 
 third-party default value modules like [creasty/defaults](https://github.com/creasty/defaults)
 because it's basically a string.
 
 
 
## Installation

You need to install `stringenum` to generate enum stub codes.

```
 $ go get github.com/therne/stringenum/...
```

## Usage

On the top of your type definition sources, add `go generate` clause to generate stub codes with `stringenum`.

```diff
+ //go:generate stringenum Kind
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

## Generated Values and Methods


 * `<Type>Values`: A list of all available values in the enum.
 * `<Type>FromString(string)`: Casts string into the enum. An error is returned if given string is not defined on the enum.
 * `IsValid()`: Returns false if the value is not defined on the enum.
 * `Validate()`: Returns an error if the value is not defined on the enum.
 * `MarshalText` / `UnmarshalText`: Implements `encoding.TextMarshaler` / `TextUnmarshaler` interface for JSON / YAML serialization.
 * `String() `: Casts the enum into a `string`. Implements `fmt.Stringer` interface.


## License: MIT