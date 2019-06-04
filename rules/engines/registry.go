package engines

import (
	"errors"
	"fmt"
	"strings"

	"github.com/stackrox/default-authz-plugin/rules"
)

// EngineCreator creates a rule engine. The string argument specifies options for the engine, such as the location of
// a configuration file.
type EngineCreator func(string) (rules.Engine, error)

var (
	engineTypes = make(map[string]EngineCreator)
)

// RegisterEngineType registers a rule engine creator for the given rule engine type name. This function should only be
// called from `init()` functions.
func RegisterEngineType(typeName string, creator EngineCreator) {
	if _, exists := engineTypes[typeName]; exists {
		panic(fmt.Errorf("rule engine of type %q already registered", typeName))
	}
	engineTypes[typeName] = creator
}

// RegisterStaticEngineType registers a rule engine under the given type name. This can be used for rule engines that
// accept no options.
func RegisterStaticEngineType(typeName string, engine rules.Engine) {
	creator := func(params string) (rules.Engine, error) {
		if params != "" {
			return nil, fmt.Errorf("rule engine of type %q does not take parameters", typeName)
		}
		return engine, nil
	}
	RegisterEngineType(typeName, creator)
}

// GetRuleEngine creates and returns a rule engine from the given specification. The specification must be of the form
// `<name>[:<options>]`, where `<name>` is the name of a rule engine type previously registered via one of the above
// functions.
func GetRuleEngine(nameWithParams string) (rules.Engine, error) {
	parts := strings.SplitN(nameWithParams, ":", 2)
	if len(parts) == 0 {
		return nil, errors.New("empty type name/params")
	}
	engineCreator := engineTypes[parts[0]]
	if engineCreator == nil {
		return nil, fmt.Errorf("invalid rule engine type name %q", parts[0])
	}

	var params string
	if len(parts) > 1 {
		params = parts[1]
	}
	return engineCreator(params)
}
