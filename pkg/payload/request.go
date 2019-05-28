package payload

import (
	"encoding/json"
	"fmt"
	"io"
)

// AuthorizationRequest encapsulates an authorization request to the authorization plugin.
type AuthorizationRequest struct {
	// Principal is the principal for which to check authorization.
	Principal Principal `json:"principal,omitempty"`

	// Scopes are the scopes to be authorized.
	RequestedScopes []AccessScope `json:"requestedScopes,omitempty"`
}

// ParseAndValidateRequest parses an authorization request from the given source, and ensures that all requested scopes are
// well-formed.
func ParseAndValidateRequest(src io.Reader) (*AuthorizationRequest, error) {
	var req AuthorizationRequest
	if err := json.NewDecoder(src).Decode(&req); err != nil {
		return nil, err
	}

	for _, scope := range req.RequestedScopes {
		if err := ValidateScope(&scope); err != nil {
			return nil, fmt.Errorf("validating requested scope: %v", err)
		}
	}

	return &req, nil
}
