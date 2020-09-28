package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/therne/stringenum"
)

func main() {
	var (
		parsingOpts  stringenum.ParsingOptions
		outputSuffix string
	)
	pflag.StringVarP(&outputSuffix, "output-file-suffix", "o", "enums", "postfix appended in generated output sources (e.g. srcName_enums.go) ")
	pflag.Parse()

	if pflag.NArg() < 1 {
		log.Fatalln("usage: stringenum <types...>")
	}
	parsingOpts.Types = pflag.Args()

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalln("unable to locate pwd")
	}
	parsedFiles, err := stringenum.Parse(pwd, parsingOpts)
	if err != nil {
		log.Fatalln(err.Error())
	}
	found := make(map[string]bool)
	for fileName, parsedFile := range parsedFiles {
		code, err := stringenum.GenerateCode(parsedFile)
		if err != nil {
			fmt.Printf("error generating code for %s: %v\n", fileName, err)
			return
		}

		outputFileName := fmt.Sprintf("%s_%s.go", fileName[:len(fileName)-3], outputSuffix)
		if err := ioutil.WriteFile(outputFileName, []byte(code), 0644); err != nil {
			fmt.Printf("error writing generated source %s: %v\n", outputFileName, err)
			return
		}
		fmt.Println("generated:", filepath.Join(pwd, outputFileName))
		for _, enum := range parsedFile.Enums {
			found[enum.Type] = true
		}
	}
	if len(parsingOpts.Types) != len(found) {
		for _, name := range parsingOpts.Types {
			if !found[name] {
				log.Fatalln("enum", name, "not found in source files.")
			}
		}
	}
}
