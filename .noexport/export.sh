#!/usr/bin/env bash

# Export source archives.
# The optional first argument specifies the parent output directory.

set -euo pipefail

root="$(git rev-parse --show-toplevel)"
[[ -n "$root" ]] || {
	echo >&2 "Could not determine git root directory."
	exit 1
}

out_dir="${1:-$PWD}"

tag="$(make -s -C "$root" tag)"

exclude_patterns=()

exclude_file="${root}/.noexport/EXCLUDE"
if [[ -f "$exclude_file" ]]; then
	IFS=$'\n' read -d '' -r -a exclude_patterns < <(egrep -v '^\s*(#.*)?$' "$exclude_file") || true
fi

exclude_regex="$(printf '(%s)|' "${exclude_patterns[@]}")"
# The following serves two purposes: it ensures that there is no trailing `|` in the expression,
# and it prevents matching all lines in case of an empty EXCLUDE file.
exclude_regex+='$^'

tmpdir="$(mktemp -d)"
export_dir="${tmpdir}/default-authz-plugin"
mkdir "$export_dir"

git -C "${root}" ls-files | egrep -v '^\.noexport/' | egrep -v '(^|/)\.gitignore$' | egrep -v "$exclude_regex" \
	| xargs -- tar -C "${root}" -cf - \
	| tar -C "$export_dir" -xf -

"${root}/.noexport/generate-third-party-notices.sh" "${export_dir}/THIRD_PARTY_NOTICES"

archive_basename="default-authz-plugin-${tag}-src"

mkdir -p "$out_dir"
cd "$out_dir"
out_dir_abs="$(pwd)"

tar -C "$tmpdir" -cvzf "${out_dir_abs}/${archive_basename}.tar.gz" default-authz-plugin
( cd "$tmpdir" ; zip -r "${out_dir_abs}/${archive_basename}.zip" default-authz-plugin )

echo "Wrote source archives to ${out_dir_abs}"
