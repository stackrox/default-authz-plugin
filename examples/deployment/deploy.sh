#!/usr/bin/env bash

set -euo pipefail

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

kubectl -n stackrox get secret stackrox >/dev/null 2>&1 || {
	echo >&2 "'stackrox' image pull secrets or namespace do not exist."
	echo >&2 "Please launch StackRox in this cluster before running this script."
	exit 1
}

TLS_CERT_FILE="${TLS_CERT_FILE:-${DIR}/../config/server-tls.crt}"
TLS_KEY_FILE="${TLS_KEY_FILE:-${DIR}/../config/server-tls.key}"

SERVER_CONFIG_FILE="${SERVER_CONFIG_FILE:-${DIR}/config/server-config.json}"
RULES_FILE="${RULES_FILE:-${DIR}/config/rules.gval}"

AUTHZ_PLUGIN_IMAGE="${AUTHZ_PLUGIN_IMAGE:-stackrox/default-authz-plugin:latest}"

kubectl -n stackrox create secret tls authz-plugin-tls \
	--cert "${TLS_CERT_FILE}" \
	--key "${TLS_KEY_FILE}" \
	--dry-run -o yaml | kubectl apply -f -

kubectl -n stackrox create configmap authz-plugin-config \
	--from-file server-config.json="${SERVER_CONFIG_FILE}" \
	--from-file rules.gval="${RULES_FILE}" \
	--dry-run -o yaml | kubectl apply -f -

sed -e 's@${IMAGE}@'"$AUTHZ_PLUGIN_IMAGE"'@g' <"${DIR}/authz-plugin.yaml.template" | kubectl apply -f -
