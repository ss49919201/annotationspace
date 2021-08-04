package main

import (
	"github.com/s-beats/anotationspace"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(anotationspace.Analyzer) }
