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

package rules

import "github.com/stackrox/default-authz-plugin/pkg/payload"

// Engine abstracts the functionally of a rule engine, which essentially is a predicate on a principal/scope pair.
type Engine interface {
	// Authorized checks if the given principal is allowed to access the requested scope.
	Authorized(principal *payload.Principal, scope *payload.AccessScope) (bool, error)
}
