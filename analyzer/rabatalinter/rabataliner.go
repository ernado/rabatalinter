package rabatalinter

import (
	"go/ast"
	"slices"
	"strings"

	"golang.org/x/tools/go/analysis"
)

func processTestFile(pass *analysis.Pass, f *ast.File) {
	// RB0: Find all receivers and check they are in allowlist.
	allowedReceivers := []string{"m", "suite", "mock"}
	fileName := pass.Fset.Position(f.Pos()).Filename

	if strings.HasSuffix(fileName, ".pb.go") {
		return
	}

	ast.Inspect(f, func(n ast.Node) bool {
		// Check if it's a function declaration with a receiver.
		fn, ok := n.(*ast.FuncDecl)
		if !ok || fn.Recv == nil || len(fn.Recv.List) == 0 {
			return true
		}

		// Get the receiver field (usually only one).
		recvField := fn.Recv.List[0]

		// Check if the receiver has a name (it should).
		if len(recvField.Names) == 0 {
			return true
		}

		// Get the receiver name identifier.
		recvName := recvField.Names[0]

		// Compare the receiver name with the expected name.
		if recvName != nil && !slices.Contains(allowedReceivers, recvName.Name) {
			pass.Reportf(recvName.Pos(), "RB0: receiver should be one of %v instead of `%s`", allowedReceivers, recvName.Name)
		}

		return true
	})
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
