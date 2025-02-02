#!/bin/sh

set -eu

if [ -x "$(command -v python3)" ]; then
  alias any_python='python3'
elif [ -x "$(command -v python)" ]; then
  alias any_python='python'
elif [ -x "$(command -v python2)" ]; then
  alias any_python='python2'
else
  echo Python 2 or 3 is required to install meshctl
  exit 1
fi

if [ -z "${GLOO_MESH_VERSION:-}" ]; then
  GLOO_MESH_VERSIONS=$(curl -s https://storage.googleapis.com/storage/v1/b/meshctl/o?pageToken=CiB2Mi4xLjAtYmV0YTIvbWVzaGN0bC1saW51eC1hbWQ2NA | any_python -c "import sys; from distutils.version import StrictVersion, LooseVersion; from json import loads as l; obj = l(sys.stdin.read()); items = obj['items'];  all_releases = [release['name'].split('/')[0] for release in items]; releases = []; [releases.append(n) for n in all_releases if n not in releases]; filtered_releases = list(filter(lambda release_string: len(release_string) > 0 and StrictVersion.version_re.match(release_string[1:]) != None, releases)); filtered_releases.sort(key=LooseVersion, reverse=True); print('\n'.join(filtered_releases))")
else
  GLOO_MESH_VERSIONS="${GLOO_MESH_VERSION}"
fi

if [ "$(uname -s)" = "Darwin" ]; then
  OS=darwin
else
  OS=linux
fi

if [  "$(uname -m)" = "aarch64" ]; then
	GOARCH=arm64
elif [ "$(uname -m)" = "arm64" ]; then
	GOARCH=arm64
else
	GOARCH=amd64
fi

for gloo_mesh_version in $GLOO_MESH_VERSIONS; do

tmp=$(mktemp -d /tmp/gloo_mesh.XXXXXX)
filename="meshctl-${OS}-${GOARCH}"
url="https://storage.googleapis.com/meshctl/${gloo_mesh_version}/${filename}"

if curl -f ${url} >/dev/null 2>&1; then
  echo "Attempting to download meshctl version ${gloo_mesh_version}"
else
  continue
fi

(
  cd "$tmp"

  echo "Downloading ${filename}..."

  SHA=$(curl -sL "${url}.sha256" | cut -d' ' -f1)
  curl -sLO "${url}"
  echo "Download complete!, validating checksum..."
  checksum=$(openssl dgst -sha256 "${filename}" | awk '{ print $2 }')
  if [ "$checksum" != "$SHA" ]; then
    echo "Checksum validation failed." >&2
    exit 1
  fi
  echo "Checksum valid."
)

(
  cd "$HOME"
  mkdir -p ".gloo-mesh/bin"
  mv "${tmp}/${filename}" ".gloo-mesh/bin/meshctl"
  chmod +x ".gloo-mesh/bin/meshctl"
)

rm -r "$tmp"

echo "meshctl was successfully installed 🎉"
echo ""
echo "Add the Gloo Mesh CLI to your path with:"
echo "  export PATH=\$HOME/.gloo-mesh/bin:\$PATH"
echo ""
echo "Now run:"
echo "  meshctl install     # install Gloo Mesh management plane"
echo "Please see visit the Gloo Mesh website for more info:  https://www.solo.io/products/gloo-mesh/"
exit 0
done

echo "No versions of meshctl found."
exit 1