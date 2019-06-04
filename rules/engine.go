package rules

import "github.com/stackrox/default-authz-plugin/pkg/payload"

// Engine abstracts the functionally of a rule engine, which essentially is a predicate on a principal/scope pair.
type Engine interface {
	// Authorized checks if the given principal is allowed to access the requested scope.
	Authorized(principal *payload.Principal, scope *payload.AccessScope) (bool, error)
}
