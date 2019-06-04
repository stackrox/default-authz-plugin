package gval

import "github.com/PaesslerAG/gval"

var (
	// exprLanguage is the expression language we use for rules.
	exprLanguage = gval.Full(gval.VariableSelector(nullSafeSelector))
)
