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

// AuthorizationResponse encapsulates a response from the authorization plugin.
type AuthorizationResponse struct {
	// AuthorizedScopes are the scopes for which authorization is granted. All other requested scopes that are not
	// listed here (or subsumed by the scopes returned here) are treated as denied.
	AuthorizedScopes []AccessScope `json:"authorizedScopes,omitempty"`
}
