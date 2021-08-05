package main

import (
	"github.com/s-beats/annotationspace"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(annotationspace.Analyzer) }
