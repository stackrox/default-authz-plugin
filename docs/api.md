# Authorization Plugin API

This document describes the API of the default authorization plugin.
It also serves as a guideline for users implementing their custom
authorization plugins, which must provide the same API (except for
the name of the endpoint).

## REST Endpoint

The default authorization plugin serves its API under the URL `/authorize`. The name
of this URL is configurable in the StackRox Kubernetes Security Platform and can
be different for other authorization plugins. The endpoint must accept HTTP POST
requests with a JSON-encoded request body.

### Request Format

The request body must be a JSON object of the following form:
```
{
  "principal": {
    "authProvider": {
      "type": ... // The type of auth provider, such as oidc, saml2, or api-token.
      "name": ... // The name of the auth provider.
      "id": ... // A UID for the auth provider.
    },
    "attributes": {
      // Arbitrary attributes describing a principal (keys are always strings). Dependent on the auth provider.
      // Typically, values are lists of strings.
    }
  },
  "requestedScopes": [
  	// Each JSON object in this array is referred to as an "access scope"
  	// or "scope". A scope determines the resources a principal is trying
  	// to access, and how they are being accessed ("view" or "edit").
    {
      "verb": ... // The type of operation, "view" or "edit".
      "noun": ... // The name of the resource to be accessed, such as "Alert".
      "attributes": {
        // Attributes describing the noun. May be omitted, which means that
        // access to all objects of a given resource type irrespective of attributes
        // is requested.
        "cluster": {
          // Attributes describing the cluster, if access is constrained to a single cluster.
          "name": ... // The name of the cluster.
          "id": ... // The UID of the cluster.
        },
        "namespace": ... // The name of the namespace, if access is constrained to a single namespace.
      }
    },
    ... // additional scopes
  ]
}
```

Generally, any part of a scope may be empty, which the authorization plugin
is to interpret as access to all scopes with the specified field values, and arbitrary values
for all empty fields. Hence, the empty scope `{}` would represent access to all operations.
In practice, the StackRox Kubernetes Security Platform will always specify
the `verb` and the `noun`, though this might change in future versions.
Additionally, if a more specific part of a scope is specified, all less specific
fields must be specified as well. The order of specificity is (starting with the
most unspecific):
1. Verb
2. Noun
3. Cluster
4. Namespace

Hence, a scope is malformed if a non-empty namespace is specified, but it does not reference
a cluster. On the other hand, it is legal for a scope to specify a non-empty cluster but omit
the namespace.

As noted in the schema above, possible values for verbs are
`view` and `edit`. See the [list of resources](resources.md) for a list of
all possible nouns.

### Response Format

If the request is well-formed, the authorization plugin should respond with a
HTTP 200 (OK) response code. The response body should be a JSON-encoded
object specifying all the scopes to which access was granted:
```
{
  "authorizedScopes": [
    ... // all requested scopes to which access was granted.
  ]
}
```

The StackRox Kubernetes Security Platform will assume that access to
any scope that was requested but is not contained in the response is denied.

In case of an error (such as a malformed request, or an internal error in the authorization
plugin logic), a non-200 response code with an appropriate error message in the
body should be returned. Even though the error message is not surfaced to a user
of StackRox APIs, it will be logged and hence should not contain any sensitive
information.
