#!/bin/bash
set -eo pipefail

# if [ -z "${RELEASE_VERSION}" ]; then
#     echo "The environment variable RELEASE_VERSION needs to be set. Exiting script."
#     exit 1
# fi

go get -u github.com/tcnksm/ghr

REPOSITORY_NAME=$(basename `git rev-parse --show-toplevel`)

ci-scripts/get_release_version.sh

# ghr -t ${API_TOKEN} -r ${REPOSITORY_NAME} -c ${BRANCH} -delete ${RELEASE_VERSION} ./artifacts/
ghr -t ${API_TOKEN} -r ${REPOSITORY_NAME} -c ${BRANCH} ${RELEASE_VERSION} ./artifacts/
