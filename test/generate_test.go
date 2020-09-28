package test

import (
	"testing"

	"github.com/therne/stringenum"
)

func TestGenerateCode(t *testing.T) {
	def := stringenum.ParsedFile{
		PackageName: "example",
		Enums: []stringenum.EnumDesc{
			{
				Type: "Helvetica",
				Values: map[string]string{
					"Neue": "neue",
					"Bold": "bold",
				},
			},
			{
				Type: "NotoSans",
				Values: map[string]string{
					"Nerf": "nerf",
					"CJK":  "chJpKr",
				},
			},
		},
	}
	if _, err := stringenum.GenerateCode(def); err != nil {
		t.Fatalf("Error occurred on stringenum.GenerateCode: %v", err)
	}
}
