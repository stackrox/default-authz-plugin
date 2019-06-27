# Writing GVAL rules

The default authorization plugin allows specifying access rules
via the [Gval](https://github.com/PaesslerAG/gval) expression language.
Gval allows to define predicates over structured data (such as
JSON objects). For the use of an authorization plugin with the
StackRox Kubernetes Security Platform, these data objects
represent information about a principal (user) trying to access
a resource, the type of operation to be performed, as well as
information about the resource itself. The combination of
operation, resource type, and resource attributes is also
called an *access scope*. Access to a scope by a principal is
granted if the predicate over scope and principal evaluates
to `true`.

## Supported Operators

For a complete list of supported operators, please refer to the
[official Gval documentation](https://github.com/PaesslerAG/gval#default-language).
The default authorization plugin uses the "default" language of
Gval. The most important operators supported in this language
are:
* equality and negated equality (`==`, `!=`),
* regular expression match and non-match (`=~`, `!~`),
* logical conjunction (and) and disjunction (or) (`&&`, `||`),
* logical negation (`!`),
* access to struct fields via `.` (e.g., `foo.bar`)

## Input Data Format

The input data against which predicates are evaluated can be
thought of as a JSON object of the following form:
```
{
	"principal": {
		... information about the principal ...
	},
	"scope": {
		... information about the scope to be accessed ...
	}
}
```

See the [API description](api.md) for detailed information about
the possible shapes for `scope`s and `principal`s.

An example principal/scope combination to be checked could be:
```json
{
	"principal": {
		"authProvider": {
			"name": "api-token",
			"type": "api-token",
			"id": "0b6c4f0f-70e3-4686-aea4-e9c20fb24584"
		},
		"attributes": {
			"name": ["test-token"]
		}
	},
	"scope": {
		"verb": "view",
		"noun": "Deployment",
		"attributes": {
			"cluster": {
				"name": "remote",
				"id": "0b6c4f0f-70e3-4686-aea4-e9c20fb24584"
			}
		}
	}
}
```
This principal/scope combination represents that the API token
with name `test-token` requests access to view all deployments
in cluster `remote`.

## Writing Rules Files

A Gval rules file contains predicates over principal/scope
combinations, one rule per line. Empty lines, as well as comment
lines (starting with `#`) are ignored. A predicate must be on a single
(logical) line, but multiple "physical" lines can be concatenated to a single
logical line via the continuation sequence ` \ ` at the end of every constituent
physical line but the last one.

Access to a scope is granted for a principal if there is any
predicate in the file that evaluates to `true`. If no predicates
evaluate to `true`, access is rejected.

The following rules file grants view access to the `test-token` API
token to all resources, and edit access to all alerts in
namespaces starting with `test-` in the cluster with name
`remote`, as well as alerts in all namespaces in cluster `local`:

```
# Example Gval rules file

# Empty lines and lines starting with hash sign (`#`) will be ignored.

# Grant view access to all resources to the `test-token` user.
# Note: every predicate must be on a single line, but lines can end with ` \`
# to indicate that the following line should be considered as part of the
# current (logical) line.
principal.authProvider.type == "api-token" && \
	principal.attributes["name"][0] == "test-token" && \
	scope.verb == "view"
	
# Grant edit access to all alerts in namespaces starting with `test-`
# in the cluster with name `remote`, as well as alerts in all
# namespaces in the cluster with name `local` to the same user.
principal.authProvider.type == "api-token" && \
	principal.attributes["name"][0] == "test-token" && \
	scope.verb == "edit" && \
	scope.noun == "Alert" && \
	( \
		scope.attributes.cluster.name == "remote" && \
		scope.attributes.namespace =~ "^test-.*" \
	|| \
		scope.attributes.cluster.name == "local" \
	)
```
