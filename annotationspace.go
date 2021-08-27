package annotationspace

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

type Annotation int

const (
	colon      = ":"
	whitespace = " "
)

const (
	TODO Annotation = iota
	FIXME
	NOTE
	REFACTOR
)

const (
	todo     = "TODO"
	fixme    = "FIXME"
	note     = "NOTE"
	refactor = "REFACTOR"
)

func (a Annotation) String() string {
	switch a {
	case TODO:
		return todo
	case FIXME:
		return fixme
	case NOTE:
		return note
	case REFACTOR:
		return refactor
	default:
		panic("not declare")
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
						case strings.Contains(c.Text(), REFACTOR.String()+colon) && !strings.Contains(c.Text(), REFACTOR.String()+colon+whitespace):
							pass.Reportf(c.Pos(), "require whitespace after REFACTOR:")
						}
						// check colon
						switch {
						case strings.Contains(c.Text(), TODO.String()) && !strings.Contains(c.Text(), TODO.String()+colon):
							pass.Reportf(c.Pos(), "require colon after TODO")
						case strings.Contains(c.Text(), FIXME.String()) && !strings.Contains(c.Text(), FIXME.String()+colon):
							pass.Reportf(c.Pos(), "require colon after FIXME")
						case strings.Contains(c.Text(), NOTE.String()) && !strings.Contains(c.Text(), NOTE.String()+colon):
							pass.Reportf(c.Pos(), "require colon after NOTE")
						case strings.Contains(c.Text(), REFACTOR.String()) && !strings.Contains(c.Text(), REFACTOR.String()+colon):
							pass.Reportf(c.Pos(), "require colon after REFACTOR")
						}
					}
				}
			}
		}
	})

	return nil, nil
}
