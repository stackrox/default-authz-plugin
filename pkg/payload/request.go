package payload


// AuthorizationRequest encapsulates an authorization request to the authorization plugin.
type AuthorizationRequest struct {
	// Principal is the principal for which to check authorization.
	Principal string

	// Scopes are the scopes to be authorized.
	Scopes []AccessScope
}
