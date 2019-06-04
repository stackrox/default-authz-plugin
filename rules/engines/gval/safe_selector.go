package gval

import (
	"context"

	"github.com/PaesslerAG/gval"
)

var (
	baseParser = gval.Parser{
		Language: gval.Base(),
	}
)

// nullSafeSelector is a variable selector that safely handles null values in object paths, i.e., `foo.bar` where `foo`
// is null.
func nullSafeSelector(path gval.Evaluables) gval.Evaluable {
	return func(c context.Context, v interface{}) (interface{}, error) {
		res, _ := baseParser.Var(path...)(c, v)
		return res, nil
	}
}
