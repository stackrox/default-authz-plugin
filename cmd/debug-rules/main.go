package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/stackrox/default-authz-plugin/rules/engines/gval"
)

var (
	fileFlag = flag.String("file", "", "GVal Rules file")
)

func main() {
	if err := mainCmd(); err != nil {
		log.Fatal(err)
	}
}

func mainCmd() error {
	flag.Parse()

	file := *fileFlag

	if file == "" {
		return errors.New("no file specified, use the -file flag to specify a GVal rules file")
	}

	rules, err := gval.LoadRules(file, gval.ExprLanguage)
	if err != nil {
		return fmt.Errorf("failed to load GVal rules from file %q: %w", file, err)
	}

	fmt.Fprintln(os.Stderr, "Loaded", len(rules), "rule(s) from file", file)

	dec := json.NewDecoder(os.Stdin)
	for {
		var rawVal map[string]interface{}
		dec.InputOffset()
		err := dec.Decode(&rawVal)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return fmt.Errorf("reading JSON input: %w", err)
		}

		var matchingRules []gval.Rule
		for _, rule := range rules {
			ok, err := rule.Expression.EvalBool(context.Background(), rawVal)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error evaluating rule %q, defined at %s on input: %v\n", rule.Rule, rule.Location(), err)
				continue
			}
			if ok {
				matchingRules = append(matchingRules, rule)
			}
		}

		if len(matchingRules) == 0 {
			fmt.Fprintln(os.Stderr, "NO MATCH: no rule matched the input value")
		} else {
			for _, rule := range matchingRules {
				fmt.Fprintf(os.Stderr, "Rule %q, defined at %s, matched input\n", rule.Rule, rule.Location())
			}
		}
	}
}
