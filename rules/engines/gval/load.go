/*
Copyright 2019 StackRox Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package gval

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/PaesslerAG/gval"
)

const (
	contSuffix = ` \`
)

var (
	ignoreRE = regexp.MustCompile(`^\s*(#.*)?$`)
)

// Rule is the description (and parsed Evaluable) of a rule, loaded from a rules file.
type Rule struct {
	Filename            string
	FirstLine, LastLine int

	Rule       string
	Expression gval.Evaluable
}

// Location returns the location in the file where a rule is defined.
func (r *Rule) Location() string {
	var lineRange string
	if r.LastLine == r.FirstLine {
		lineRange = strconv.Itoa(r.FirstLine)
	} else {
		lineRange = fmt.Sprintf("%d-%d", r.FirstLine, r.LastLine)
	}
	return fmt.Sprintf("%s:%s", r.Filename, lineRange)
}

// LoadRules loads rules (specified as gval expressions) from the given file, one expression per line. Empty lines,
// or lines including comments only, are ignored.
func LoadRules(filename string, language gval.Language) ([]Rule, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var rules []Rule
	lineNo := 0

	var startLine int

	linePrefix := ""
	for scanner.Scan() {
		lineNo++
		line := strings.TrimSpace(scanner.Text())
		if ignoreRE.MatchString(line) {
			continue
		}

		if linePrefix == "" {
			startLine = lineNo
		}

		if strings.HasSuffix(line, contSuffix) {
			linePrefix += strings.TrimSuffix(line, contSuffix) + " "
			continue
		}

		exprStr := linePrefix + line
		expr, err := language.NewEvaluable(exprStr)
		if err != nil {
			return nil, fmt.Errorf("parsing expression in %s:%d: %v", filename, lineNo, err)
		}
		rule := Rule{
			Filename:   filename,
			FirstLine:  startLine,
			LastLine:   lineNo,
			Rule:       exprStr,
			Expression: expr,
		}
		rules = append(rules, rule)
		linePrefix = ""
	}

	if linePrefix != "" {
		return nil, errors.New("at end of file: line continuation with no subsequent line")
	}

	return rules, nil
}
