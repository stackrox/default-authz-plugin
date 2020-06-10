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

	"golang.org/x/crypto/ssh/terminal"

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

	if terminal.IsTerminal(int(os.Stdin.Fd())) {
		fmt.Fprintln(os.Stderr, "============================================")
		fmt.Fprintln(os.Stderr, "StackRox GVal rules evaluator")
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "This tool evaluates GVal rules from a file against arbitrary JSON objects.")
		fmt.Fprintln(os.Stderr, "Just type in the JSON objects you want to evaluate the rules against. An example")
		fmt.Fprintln(os.Stderr, "JSON object would be:")
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, `{"principal":{"authProvider":{"type":"api-token"},"attributes":{"name":["test-token"]}},"scope":{"verb":"view"}}`)
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "Please refer to the authorization plugin documentation for information on the shape")
		fmt.Fprintln(os.Stderr, "of request payloads.")
		fmt.Fprintln(os.Stderr, "If you want to evaluate the rules against JSON objects stored in a file, use your shell's")
		fmt.Fprintln(os.Stderr, "input redirection feature (< FILENAME).")
		fmt.Fprintln(os.Stderr, "============================================")
	}

	dec := json.NewDecoder(os.Stdin)
	for {
		var rawVal map[string]interface{}
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
