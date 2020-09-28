package test

import (
	"os"
	"reflect"
	"testing"

	"github.com/therne/stringenum"
)

func TestParser(t *testing.T) {
	pwd, _ := os.Getwd()
	parsed, err := stringenum.Parse(pwd, stringenum.ParsingOptions{Types: []string{"ExampleEnum"}})
	if err != nil {
		t.Fatalf("Expected stringenum.Parse to return any error, but returned:\n\t%v", err)
	}
	result, ok := parsed["example.go"]
	if len(parsed) != 1 || !ok {
		t.Fatalf("Expected example.go to be parsed only, but found %v", parsed)
	}
	if result.PackageName != "test" {
		t.Fatalf("Expected parsed package name to be 'test', but was '%v'", result.PackageName)
	}
	if len(result.Enums) != 1 {
		t.Fatalf("Expected parsed enum count to be 1, but was %v", len(result.Enums))
	}
	expectedValues := map[string]string{
		"Enum1": "hello",
		"Enum2": "world",
		"Enum3": "goos",
		"Enum4": "geese",
	}
	if !reflect.DeepEqual(result.Enums[0].Values, expectedValues) {
		t.Fatalf("Wrong parsed enums: found %v", result.Enums[0].Values)
	}
}
