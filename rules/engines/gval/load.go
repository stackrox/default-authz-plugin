package gval

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/PaesslerAG/gval"
)

const (
	contSuffix = ` \`
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

	linePrefix := ""
	for scanner.Scan() {
		lineNo++
		line := scanner.Text()
		if ignoreRE.MatchString(line) {
			continue
		}

		if strings.HasSuffix(line, contSuffix) {
			linePrefix += strings.TrimSuffix(line, contSuffix) + " "
			continue
		}

		expr, err := language.NewEvaluable(linePrefix + line)
		if err != nil {
			return nil, fmt.Errorf("parsing expression in %s:%d: %v", filename, lineNo, err)
		}
		exprs = append(exprs, expr)
		linePrefix = ""
	}

	if linePrefix != "" {
		return nil, errors.New("at end of file: line continuation with no subsequent line")
	}

	return exprs, nil
}
