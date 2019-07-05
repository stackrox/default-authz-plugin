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

// AuthProviderInfo encapsulates information about the StackRox Authentication Provider.
type AuthProviderInfo struct {
	// Type is the type of the Authentication Provider, such as "saml" or "oidc".
	Type string `json:"type,omitempty"`
	// ID is a unique ID for the Authentication Provider.
	ID string `json:"id,omitempty"`
	// Name is the user-defined name of the Authentication Provider.
	Name string `json:"name,omitempty"`
}

// Principal is an entity to be authorized by the Authorization Plugin.
type Principal struct {
	// AuthProvider contains information about the StackRox Authentication Provider.
	AuthProvider AuthProviderInfo `json:"authProvider,omitempty"`

	// Attributes are the attributes of the entity.
	Attributes map[string]interface{} `json:"attributes,omitempty"`
}
