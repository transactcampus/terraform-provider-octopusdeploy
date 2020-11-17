#!/bin/bash
set -eo pipefail	set -eo pipefail

# Download the package
go get -u github.com/tcnksm/ghr

# Get and export release version
. ci-scripts/helpers/get_release_version.sh $1

# Create release
ghr -t ${API_TOKEN} -r ${REPOSITORY_NAME} -c ${BRANCH} ${RELEASE_VERSION} ./artifacts/