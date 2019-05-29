package gval

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/PaesslerAG/gval"
)

var (
	ignoreRE = regexp.MustCompile(`^\s*(#.*)?$`)
)

// loadRules loads rules (specified as gval expressions) from the given file, one expression per line. Empty lines,
// or lines including comments only, are ignored.
func loadRules(filename string, language gval.Language) ([]gval.Evaluable, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var exprs []gval.Evaluable
	lineNo := 0
	for scanner.Scan() {
		lineNo++
		line := scanner.Text()
		if ignoreRE.MatchString(line) {
			continue
		}

		expr, err := language.NewEvaluable(line)
		if err != nil {
			return nil, fmt.Errorf("parsing expression in %s:%d: %v", filename, lineNo, err)
		}
		exprs = append(exprs, expr)
	}
	return exprs, nil
}
