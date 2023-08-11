package main

import (
	"github.com/akupila/todolint"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(todolint.Analyzer())
}
