package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
)

func main() {
	fs := token.NewFileSet()
	file, err := parser.ParseFile(fs, "template/template.go", nil, 0)
	if err != nil {
		fmt.Println("Can't parse file", err)
	}

	ast.Print(nil, file)

	ast.Walk(MyVisit(Visit), file)

	f, err := os.Create("out/main.go")
	if err != nil {
		fmt.Println("Can't open new file")
	}
	defer f.Close()

	printer.Fprint(f, fs, file)
}

type MyVisit func(node ast.Node) ast.Visitor

func (f MyVisit) Visit(node ast.Node) (w ast.Visitor) {
	return f(node)
}

func Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.CallExpr:
		if v, ok := n.Args[0].(*ast.BasicLit); ok {
			v.Value = "\"Fuga\""
		}
	}
	return MyVisit(Visit)
}
