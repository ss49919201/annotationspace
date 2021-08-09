package annotationspace

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

type annotation int

const (
	colon      = ":"
	whitespace = " "
)

const (
	TODO annotation = iota
	FIXME
	NOTE
	REFUCTOR
)

func (a annotation) String() string {
	switch a {
	case TODO:
		return "TODO"
	case FIXME:
		return "FIXME"
	case NOTE:
		return "NOTE"
	case REFUCTOR:
		return "REFUCTOR"
	default:
		panic("not decleare")
	}
}

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
						switch {
						case strings.Contains(c.Text(), TODO.String()+colon) && !strings.Contains(c.Text(), TODO.String()+colon+whitespace):
							pass.Reportf(c.Pos(), "require whitespace after TODO:")
						case strings.Contains(c.Text(), FIXME.String()+colon) && !strings.Contains(c.Text(), FIXME.String()+colon+whitespace):
							pass.Reportf(c.Pos(), "require whitespace after FIXME:")
						case strings.Contains(c.Text(), NOTE.String()+colon) && !strings.Contains(c.Text(), NOTE.String()+colon+whitespace):
							pass.Reportf(c.Pos(), "require whitespace after NOTE:")
						case strings.Contains(c.Text(), REFUCTOR.String()+colon) && !strings.Contains(c.Text(), REFUCTOR.String()+colon+whitespace):
							pass.Reportf(c.Pos(), "require whitespace after REFUCTOR:")
						}
						// check colon
						switch {
						case strings.Contains(c.Text(), TODO.String()) && !strings.Contains(c.Text(), TODO.String()+colon):
							pass.Reportf(c.Pos(), "require colon after TODO")
						case strings.Contains(c.Text(), FIXME.String()) && !strings.Contains(c.Text(), FIXME.String()+colon):
							pass.Reportf(c.Pos(), "require colon after FIXME")
						case strings.Contains(c.Text(), NOTE.String()) && !strings.Contains(c.Text(), NOTE.String()+colon):
							pass.Reportf(c.Pos(), "require colon after NOTE")
						case strings.Contains(c.Text(), REFUCTOR.String()) && !strings.Contains(c.Text(), REFUCTOR.String()+colon):
							pass.Reportf(c.Pos(), "require colon after REFUCTOR")
						}
					}
				}
			}
		}
	})

	return nil, nil
}
