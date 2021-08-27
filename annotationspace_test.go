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

func Test_annotation_String(t *testing.T) {
	tests := []struct {
		name string
		a    annotationspace.Annotation
		want string
	}{
		{
			"TODO",
			annotationspace.TODO,
			"TODO",
		},
		{
			"FIXME",
			annotationspace.FIXME,
			"FIXME",
		},
		{
			"NOTE",
			annotationspace.NOTE,
			"NOTE",
		},
		{
			"REFACTOR",
			annotationspace.REFACTOR,
			"REFACTOR",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.String(); got != tt.want {
				t.Errorf("annotation.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
