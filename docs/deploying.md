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

**A note on performance:** When enabled, the Authorization Plugin is queried for all
accesses to the StackRox API originating from human or API token users. In order to
ensure a minimal impact on latency, we strongly recommend to run the Authorization Plugin
on the same node where the StackRox Central Pod is running. The aforementioned YAML
template includes an [inter-pod affinity](https://kubernetes.io/docs/concepts/configuration/assign-pod-node/)
configuration for this.

## Accessing the Docker Image

A pre-built Docker image for the Default Authorization Plugin is published via the `stackrox.io`
registry, and can be pulled as `stackrox.io/default-authz-plugin:1.0`. This image is equivalent to
the image built by running `make image` on unmodified sources (see also the documentation on [building](building.md)).

The image published via the `stackrox.io` registry is accessible only to authenticated clients. You can
use the `stackrox.io` credentials you should have received as a customer of the StackRox Kubernetes Security
Platform.

For deployment in a Kubernetes cluster, these credentials need to be stored as an *image pull secret*.
The deployment YAML template contained in this repository assumes this secret is called `stackrox`.
This matches the name of the image pull secret used for pulling the main StackRox Kubernetes Security
Platform Docker images, created by the `setup.sh` script that is part of the StackRox Central deployment
bundle. Hence, if you deploy the Default Authorization Plugin in the same Kubernetes cluster that is running
StackRox Central (recommended), no further action with respect to image pull secret setup is required. If you
deploy it in a different cluster, run the `setup.sh` script from the StackRox Central deployment bundle in the respective
cluster.

If you require custom modifications to the Default Authorization Plugin, or wish to build
the Docker image yourself, we assume you know the relevant steps of deploying from the respective
registry that stores your built and pushed image inside your cluster. If you do not require image pull
secrets for accessing this registry, remove the relevant `imagePullSecrets` parts from the
`authz-plugin.yaml.template` YAML template.

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
- `AUTHZ_PLUGIN_IMAGE` is the image of the Authorization Plugin. This defaults to the image published
  via the `stackrox.io` registry, i.e., `stackrox.io/default-authz-plugin:<version>`. Change this if
  you obtain the image from a different registry.

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
