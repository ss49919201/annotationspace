package annotationspace

import (
	"go/ast"
	"regexp"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var regexObject = regexp.MustCompile(`^[A-Z]+:[^ ]+`)

var Analyzer = &analysis.Analyzer{
	Name: "annotationspace",
	Doc:  Doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

const Doc = "annotationspace is ..."

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.File)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.File:
			if n.Comments != nil {
				for _, c := range n.Comments {
					if c.List != nil {
						// check whitespace
						if regexObject.MatchString(c.Text()) {
							pass.Reportf(c.Pos(), "require whitespace after annotation comment")
						}
						// TODO: check colon
					}
				}
			}
		}
	})

	return nil, nil
}
