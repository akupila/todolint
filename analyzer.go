package todolint

import (
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/tools/go/analysis"
)

// Analyzer returns a new todolint analyzer.
func Analyzer() *analysis.Analyzer {
	a := &analysis.Analyzer{
		Name: "todolint",
		Doc:  "Checks that TODO comments are formatted properly",
		Run:  run,
	}
	a.Flags.String("keywords", "TODO,FIXME,BUG", "Comment patterns to check")
	a.Flags.String("context", "\\w+", "Regular expression the context must match")
	return a
}

// (context): <summary>
var format = regexp.MustCompile(
	`^(\(([^)]*)\))?` +
		`(:)?` +
		`( )?` +
		`\s*` +
		`([^/]*)`,
)

func run(pass *analysis.Pass) (interface{}, error) {
	contextRe, err := regexp.Compile(pass.Analyzer.Flags.Lookup("context").Value.String())
	if err != nil {
		return nil, fmt.Errorf("invalid context regexp: %w", err)
	}

	keywords := strings.Split(pass.Analyzer.Flags.Lookup("keywords").Value.String(), ",")
	for i, keyword := range keywords {
		keywords[i] = strings.TrimSpace(keyword)
	}

	for _, file := range pass.Files {
		for _, group := range file.Comments {
			for _, c := range group.List {
				if c.Text[0:2] != "//" {
					// Only check single line comments
					continue
				}
				text := c.Text[2:]

				from, to := findWord(text, keywords)
				if to == 0 {
					// No match.
					continue
				}
				keyword := text[from:to]
				if from != 1 {
					pass.Reportf(c.Pos(), "%s should be at the beginning of the line", keyword)
					continue
				}
				text = text[to:]

				// Always matches as all the groups are optional.
				match := format.FindStringSubmatch(text)
				paren := match[1]
				context := match[2]
				colon := match[3]
				space := match[4]
				summary := match[5]

				if paren == "" {
					pass.Reportf(c.Pos(), "%s should include additional context: %s(<context>)", keyword, keyword)
					continue
				}
				if !contextRe.MatchString(context) {
					pass.Reportf(c.Pos(), "%s context does not match regular expression: %s", keyword, contextRe.String())
				}
				if colon == "" || space == "" {
					pass.Report(analysis.Diagnostic{
						Pos:     c.Pos(),
						Message: fmt.Sprintf("%s should follow the format '// %s(context): text'", keyword, keyword),
						SuggestedFixes: []analysis.SuggestedFix{
							{
								Message: "Add colon and space",
								TextEdits: []analysis.TextEdit{
									{
										Pos:     c.Pos(),
										End:     c.End(),
										NewText: []byte(fmt.Sprintf("// %s%s: %s", keyword, paren, summary)),
									},
								},
							},
						},
					})
				}
				if summary == "" {
					pass.Reportf(c.Pos(), "%s should describe what needs to change", keyword)
				}
			}
		}
	}
	return nil, nil
}

func findWord(str string, words []string) (from, to int) {
	for _, w := range words {
		if i := strings.Index(str, w); i >= 0 {
			return i, i + len(w)
		}
	}
	return 0, 0
}
