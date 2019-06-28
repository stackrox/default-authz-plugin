# Deploying the Default Authorization Plugin

The Default Authorization Plugin can be deployed in a Kubernetes cluster.

A sample deployment template YAML can be found in the file
`examples/deployment/authz-plugin.yaml.template`. This file can be used as
a basis for writing a custom deployment YAML, or in conjunction with the
`examples/deployment/deploy.sh` script, which deploys the Default Authorization
Plugin with pre-baked rules and configuration into the cluster in which the Central
component of the StackRox Kubernetes Security Platform is deployed.

When using the above script, the Authorization Plugin endpoint is accessible in the
cluster via `https://authorization-plugin.stackrox/authorize`. A network policy is
created that restricts access to this endpoint to Central. The TLS certificates

## TLS Certificates

The above deploy script uses pre-created TLS certificates. We **strongly** recommend
to not use these certificates in any environment that you care about, and instead
issue certificates from a trusted CA within your organization. When doing so, make sure
that the `Common Name (CN)` or one of the DNS `Subject Alternative Names (SANs)` matches
the endpoint under which the authorization plugin will be accessible in the cluster 
(`authorization-plugin.stackrox` for the example config).

If you choose to deploy using the certificates shipped with this software, you must set the
TLS configuration to "insecure" in the StackRox Kubernetes Security Platform settings for this
Authorization Plugin.

## Using the Deploy Script with Non-Default Values

It is possible to use the deploy script with values other than the default ones by setting
environment variables prior to running `deploy.sh`. These environment variables are:
- `TLS_CERT_FILE` and `TLS_KEY_FILE` control the location of the TLS certificate/key files.
  These files will be mounted into the container under the directory
  `/run/secrets/stackrox.io/default-authz-plugin/tls-certs/` and be named
  `tls.crt` and `tls.key`.
- `SERVER_CONFIG_FILE` controls the location of the server configuration file.
  This file will be mounted into the container as
  `/etc/stackrox-authz-plugin/config/server-config.json`.
  Note that any server configuration to be used with the YAML template must expose the API
  on port 8443.
- `RULES_FILE` controls the location of the [Gval rules file](writing-gval-rules.md) that defines
  access rules. This file will be mounted into the container as
  `/etc/stackrox-authz-plugin/config/rules.gval`.
- `AUTHZ_PLUGIN_IMAGE` is the image of the Authorization Plugin.

By default, the Default Authorization Plugin is set up to allow anonymous access (the network
policy ensures that only Central can talk to its API port). If this seems too insecure for your
environment, or if you are using a CNI plugin that does not support network policies, you can
choose to enforce either HTTP Basic or Client Certificate-based authentication by setting *either*
of the following environment variables:
- `HTPASSWD_FILE` is the location of an `.htpasswd` file (only bcrypt is supported).
  This file will be mounted into the container as
  `/run/secrets/stackrox.io/default-authz-plugin/auth-basic/.htpasswd`
- `CLIENT_CA_FILE` is the locaton of a PEM file containing a client certificate authority.
  This file will be mounted into the container as
  `/run/secrets/stackrox.io/default-authz-plugin/auth-clientcert/client-ca.crt`.
  
Note that if you do not override the `SERVER_CONFIG_FILE` setting, the `deploy.sh` script will
transparently switch to a server configuration file suitable for the chosen authentication method.
If you explicitly specify a `SERVER_CONFIG_FILE`, this explicit choice will always be used, hence
you need to take care it is configured to be used with the chosen authentication method.
