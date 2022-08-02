package main

import (
	"github.com/fatih/addlint/addcheck"
	"github.com/fatih/addlint/stringcheck"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(
		addcheck.Analyzer,
		stringcheck.Analyzer,
	)
}
