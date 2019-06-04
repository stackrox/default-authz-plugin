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
