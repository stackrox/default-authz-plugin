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
