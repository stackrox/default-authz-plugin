package payload

import "fmt"

// Cluster identifies a Cluster managed by the StackRox Kubernetes Security Platform.
type Cluster struct {
	// ID is the unique ID of the cluster.
	ID string
	// Name is the user-defined name of the cluster.
	Name string
}

// NounAttributes are the (optional) attributes of a noun, i.e., cluster and namespace.
type NounAttributes struct {
	Cluster Cluster `json:",omitempty"`
	Namespace string `json:",omitempty"`
}

// AccessScope defines an access scope to be accessed, consisting of a verb (operation, "read" or "edit"), a noun
// (resource, e.g., "deployment"), and possibly attributes further describing the noun (cluster and namespace).
// A scope may not be fully specified, and any unset attribute is interpreted as encompassing all scopes for all
// possible values of the omitted attribute. For example, `{Verb: "edit", Noun: "deployment"}` is a scope representing
// edit access to all deployments in all clusters, and if a cluster is added in the attributes, this would represent
// edit access to all deployments in all namespaces in the respective cluster. However, if the cluster is omitted,
// namespace must also be omitted, and if the noun is omitted (to check for global read or global edit access), the
// attributes must be omitted, too. A scope that does not satisfy these constraints is invalid.
type AccessScope struct {
	Verb string `json:",omitempty"`
	Noun string `json:",omitempty"`

	Attributes NounAttributes `json:",omitempty"`
}

// ValidateScope checks if an AccessScope is valid, according to the above description.
func ValidateScope(scope *AccessScope) error {
	if scope.Verb == "" && scope.Noun != "" {
		return fmt.Errorf("scope omits verb, but declares a noun (%q)", scope.Noun)
	}
	if scope.Noun == "" && scope.Attributes.Cluster.ID != "" {
		return fmt.Errorf("scope omits noun, but declares a cluster ID attribute (%q)", scope.Attributes.Cluster.ID)
	}
	if (scope.Attributes.Cluster.ID == "") != (scope.Attributes.Cluster.Name == "") {
		return fmt.Errorf("scope must declare either both or none of cluster ID (%q) and cluster name (%q)", scope.Attributes.Cluster.ID, scope.Attributes.Cluster.Name)
	}
	if scope.Attributes.Cluster.ID == "" && scope.Attributes.Namespace != "" {
		return fmt.Errorf("scope omits cluster, but not namespace (%q)", scope.Attributes.Namespace)
	}
	return nil
}
