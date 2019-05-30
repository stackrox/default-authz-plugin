package payload

// AuthorizationResponse encapsulates a response from the authorization plugin.
type AuthorizationResponse struct {
	// AuthorizedScopes are the scopes for which authorization is granted. All other requested scopes that are not
	// listed here (or subsumed by the scopes returned here) are treated as denied.
	AuthorizedScopes []AccessScope `json:"authorizedScopes,omitempty"`
}
