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
