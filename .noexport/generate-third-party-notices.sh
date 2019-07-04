#!/usr/bin/env bash

set -euo pipefail

root="$(git rev-parse --show-toplevel)"
[[ -n "$root" ]] || {
	echo >&2 "Could not determine git root directory."
	exit 1
}

tpn_dir="${1:-${root}/THIRD_PARTY_NOTICES}"

[[ ! -d "${tpn_dir}" ]] || rm -rf "${tpn_dir}"

( cd "$root" ; ossls audit -c "${root}/.noexport/ossls.yml" -x "$tpn_dir" )

