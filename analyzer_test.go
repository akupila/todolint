package todolint_test

import (
	"testing"

	"github.com/akupila/todolint"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestRun(t *testing.T) {
	testdata := analysistest.TestData()
	analyzer := todolint.Analyzer()
	analysistest.RunWithSuggestedFixes(t, testdata, analyzer)
}
