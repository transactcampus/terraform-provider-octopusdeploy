#!/bin/bash
set -eo pipefail

# Get and export release version
. ci-scripts/helpers/get_release_version.sh $1

# Download the package
go get -u github.com/tcnksm/ghr

# Get repository name
REPOSITORY_NAME=$(basename `git rev-parse --show-toplevel`)

echo "Creating new release ${REPOSITORY_NAME}/${BRANCH}/${RELEASE_VERSION}"

# Create release
ghr -t ${API_TOKEN} -r ${REPOSITORY_NAME} -c ${BRANCH} ${RELEASE_VERSION} ./artifacts/