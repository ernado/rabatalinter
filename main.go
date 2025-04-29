package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/ernado/rabatalinter/analyzer/rabatalinter"
)

func main() {
	singlechecker.Main(rabatalinter.NewAnalyzer())
}
