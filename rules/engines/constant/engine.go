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

package constant

import (
	"github.com/stackrox/default-authz-plugin/pkg/payload"
	"github.com/stackrox/default-authz-plugin/rules/engines"
)

type engine struct {
	response bool
}

func (e engine) Authorized(*payload.Principal, *payload.AccessScope) (bool, error) {
	return e.response, nil
}

func init() {
	engines.RegisterStaticEngineType("allow_all", engine{response: true})
	engines.RegisterStaticEngineType("deny_all", engine{response: false})
}
