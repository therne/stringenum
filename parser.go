package stringenum

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"strings"
)

type ParsingOptions struct {
	Types []string
}

type ParsedFile struct {
	PackageName string
	Enums       []EnumDesc
}

type EnumDesc struct {
	Type   string
	Values map[string]string
}

func Parse(srcDir string, opt ParsingOptions) (parsedFiles map[string]ParsedFile, err error) {
	parsedFiles = make(map[string]ParsedFile)

	files, err := ioutil.ReadDir(srcDir)
	if err != nil {
		return
	}
	toks := token.NewFileSet()
	for _, f := range files {
		if f.IsDir() || !strings.HasSuffix(f.Name(), ".go") || strings.HasSuffix(f.Name(), "_test.go") {
			continue
		}
		src, err := parser.ParseFile(toks, f.Name(), nil, 0)
		if err != nil {
			return nil, err
		}
		v := newAstVisitor(opt)
		ast.Walk(v, src)
		if v.Error != nil {
			return nil, v.Error
		}
		result := v.Result()
		if len(result.Enums) > 0 {
			parsedFiles[f.Name()] = result
		}
	}
	return
}

type astVisitor struct {
	Error error

	packageName   string
	intermediates map[string]*EnumDesc
	options       ParsingOptions
}

func newAstVisitor(opt ParsingOptions) *astVisitor {
	return &astVisitor{
		intermediates: make(map[string]*EnumDesc),
		options:       opt,
	}
}

func (a *astVisitor) Visit(node ast.Node) (w ast.Visitor) {
	switch n := node.(type) {
	case *ast.File:
		a.packageName = fmt.Sprint(n.Name)
		return a
	case *ast.GenDecl:
		if n.Tok.String() == "type" {
			// check that type is string
			typeSpec := n.Specs[0].(*ast.TypeSpec)
			typ, originalType := fmt.Sprint(typeSpec.Name), fmt.Sprint(typeSpec.Type)

			isTarget := false
			for _, target := range a.options.Types {
				if typ == target {
					isTarget = true
					break
				}
			}
			if isTarget && originalType != "string" {
				a.Error = fmt.Errorf("expected type %s to be string, but was %s", typ, originalType)
				return nil
			}
			return a
		} else if n.Tok.String() == "var" {
			return nil
		}
		for _, entry := range n.Specs {
			decl, ok := entry.(*ast.ValueSpec)
			if !ok || len(decl.Values) != 1 {
				continue
			}
			name := fmt.Sprint(decl.Names[0])

			if typ, ok := decl.Type.(*ast.Ident); ok {
				if literal, ok := decl.Values[0].(*ast.BasicLit); ok && literal.Kind.String() == "STRING" {
					// case 1) const Type EnumType = "value"
					a.addResult(typ.Name, name, literal.Value)
				}

			} else if cast, ok := decl.Values[0].(*ast.CallExpr); ok && len(cast.Args) == 1 {
				if literal, ok := cast.Args[0].(*ast.BasicLit); ok && literal.Kind.String() == "STRING" {
					// case 2) const Type = EnumType("value")
					a.addResult(fmt.Sprint(cast.Fun), name, literal.Value)
				}
			}
		}
		return nil
	}
	return nil
}

func (a *astVisitor) addResult(typ, name, value string) {
	enum := a.intermediates[typ]
	if enum == nil {
		enum = &EnumDesc{
			Type:   typ,
			Values: make(map[string]string),
		}
		a.intermediates[typ] = enum
	}
	enum.Values[name] = value
}

func (a *astVisitor) Result() (res ParsedFile) {
	res.PackageName = a.packageName
	for _, enumName := range a.options.Types {
		if enum, ok := a.intermediates[enumName]; ok {
			res.Enums = append(res.Enums, *enum)
		}
	}
	return res
}
