#!/bin/bash
set -eo pipefail	set -eo pipefail


# if [ -z "${RELEASE_VERSION}" ]; then	# if [ -z "${RELEASE_VERSION}" ]; then
    # echo "The environment variable RELEASE_VERSION needs to be set. Exiting script."	#     echo "The environment variable RELEASE_VERSION needs to be set. Exiting script."
    # exit 1	#     exit 1
# fi	# fi


go get -u github.com/tcnksm/ghr


ghr -u ${CIRCLE_PROJECT_USERNAME} -r ${CIRCLE_PROJECT_REPONAME} -c ${CIRCLE_SHA1} -delete ${RELEASE_VERSION} ./artifacts/	REPOSITORY_NAME=$(basename `git rev-parse --show-toplevel`)

. ci-scripts/helpers/get_release_version.sh $1

# ghr -t ${API_TOKEN} -r ${REPOSITORY_NAME} -c ${BRANCH} -delete ${RELEASE_VERSION} ./artifacts/
ghr -t ${API_TOKEN} -r ${REPOSITORY_NAME} -c ${BRANCH} ${RELEASE_VERSION} ./artifacts/