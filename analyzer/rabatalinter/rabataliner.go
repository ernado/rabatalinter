package rabatalinter

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

func processTestFile(pass *analysis.Pass, f *ast.File) {
	// RB0: Find all receivers and check they are equal to `m`.
	const expectedReceiver = "m"
	for _, decl := range f.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok && fn.Recv != nil {
			for _, field := range fn.Recv.List {
				if starExpr, ok := field.Type.(*ast.StarExpr); ok {
					if ident, ok := starExpr.X.(*ast.Ident); ok {
						if ident.Name != expectedReceiver {
							pass.Reportf(field.Pos(), "RB0: receiver should be `%s` instead of `%s`", expectedReceiver, ident.Name)
						}
					}
				}
			}
		}
	}
}

// NewAnalyzer returns Analyzer that makes you use a separate _test package.
func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "rabatalinter",
		Doc:  "Linter for rabataio",
		Run: func(pass *analysis.Pass) (interface{}, error) {
			for _, f := range pass.Files {
				processTestFile(pass, f)
			}

			return nil, nil
		},
	}
}
