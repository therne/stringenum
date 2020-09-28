//go:generate go run ../cmd/stringenum ExampleEnum
package test

type ExampleEnum string

const (
	Enum1 = ExampleEnum("hello")
	Enum2 = ExampleEnum("world")
	Enum3 = ExampleEnum("goos")
)

const Enum4 ExampleEnum = "geese"

const irrelevantConstDecl = 1234

var irrelevantVarDecl = 5678
