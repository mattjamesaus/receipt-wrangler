#!/usr/bin/env bash
# Delete GHCR tags by name. Used after creating a multi-arch manifest so the
# temporary architecture-specific tags (e.g. latest-amd64) do not linger.
set -euo pipefail

if [[ $# -lt 1 ]]; then
  echo "usage: $0 <tag> [tag...]" >&2
  exit 1
fi

IMAGE="${IMAGE:?IMAGE must be set (e.g. ghcr.io/owner/repo)}"
OWNER_TYPE="${OWNER_TYPE:-User}"

path_without_registry="${IMAGE#ghcr.io/}"
owner="${path_without_registry%%/*}"
package="${path_without_registry#*/}"
package_enc="$(jq -rn --arg p "$package" '$p | @uri')"

if [[ "$OWNER_TYPE" == "Organization" ]]; then
  api_root="orgs/${owner}"
else
  api_root="users/${owner}"
fi

delete_tag() {
  local tag="$1"
  local version_id
  local ids

  if [[ ! "$tag" =~ ^[A-Za-z0-9._/-]+$ ]]; then
    echo "refusing unsafe GHCR tag name: ${tag}" >&2
    exit 1
  fi

  # Read the full listing (avoid `head` + pipefail SIGPIPE) and take the first id.
  ids="$(
    gh api --paginate \
      --jq ".[] | select(.metadata.container.tags[]? == \"${tag}\") | .id" \
      "${api_root}/packages/container/${package_enc}/versions"
  )"
  version_id="${ids%%$'\n'*}"
  if [[ -n "$version_id" ]]; then
    echo "Deleting GHCR tag ${tag} (version ${version_id})"
    gh api --method DELETE "${api_root}/packages/container/${package_enc}/versions/${version_id}"
  else
    echo "No GHCR version found for tag ${tag}; skipping"
  fi
}

for tag in "$@"; do
  delete_tag "$tag"
done
