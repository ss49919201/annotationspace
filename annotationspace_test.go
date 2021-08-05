package annotationspace_test

import (
	"testing"

	"github.com/s-beats/annotationspace"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, annotationspace.Analyzer, "a")
}
