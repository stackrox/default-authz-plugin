package payload

// AuthProviderInfo encapsulates information about the StackRox Authentication Provider.
type AuthProviderInfo struct {
	// Type is the type of the Authentication Provider, such as "saml" or "oidc".
	Type string
	// ID is a unique ID for the Authentication Provider.
	ID string
	// Name is the user-defined name of the Authentication Provider.
	Name string
}

// Principal is an entity to be authorized by the Authorization Plugin.
type Principal struct {
	// AuthProvider contains information about the StackRox Authentication Provider.
	AuthProvider AuthProviderInfo

	// Attributes are the attributes of the entity.
	Attributes map[string]interface{}
}
