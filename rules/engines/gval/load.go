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
