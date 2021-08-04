package anotationspace_test

import (
	"testing"

	"github.com/s-beats/anotationspace"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, anotationspace.Analyzer, "a")
}
