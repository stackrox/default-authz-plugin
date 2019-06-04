package gval

import (
	"context"
	"errors"
	"fmt"

	"github.com/PaesslerAG/gval"
	"github.com/stackrox/sample-authz-plugin/pkg/jsonutil"
	"github.com/stackrox/sample-authz-plugin/pkg/payload"
	"github.com/stackrox/sample-authz-plugin/rules"
	"github.com/stackrox/sample-authz-plugin/rules/engines"
)

type engine struct {
	expressions []gval.Evaluable
}

type singleRequest struct {
	Principal *payload.Principal   `json:"principal"`
	Scope     *payload.AccessScope `json:"scope"`
}

func (e engine) Authorized(principal *payload.Principal, scope *payload.AccessScope) (bool, error) {
	req := singleRequest{
		Principal: principal,
		Scope:     scope,
	}

	rawReq, err := jsonutil.ToRaw(req)
	if err != nil {
		return false, fmt.Errorf("converting request object: %v", err)
	}

	for _, expr := range e.expressions {
		if allowed, err := expr.EvalBool(context.Background(), rawReq); err != nil {
			return false, err
		} else if allowed {
			return true, nil
		}
	}
	return false, nil
}

func createGvalEngine(options string) (rules.Engine, error) {
	if options == "" {
		return nil, errors.New("gval engine requires a rules file as an option (or a single rule prefixed with `@`)")
	}

	var exprs []gval.Evaluable
	var err error
	if options[0] == '@' {
		exprs = make([]gval.Evaluable, 1)
		exprs[0], err = exprLanguage.NewEvaluable(options[1:])
	} else {
		exprs, err = loadRules(options, exprLanguage)
	}

	if err != nil {
		return nil, err
	}

	return engine{
		expressions: exprs,
	}, nil
}

func init() {
	engines.RegisterEngineType("gval", createGvalEngine)
}
