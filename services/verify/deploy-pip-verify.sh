#!/bin/bash
#export TERM=ansi
#export TERM=xterm-color
export PROJECT_ROOT=${PROJECT_ROOT:-"/home/circleci/project"}
export EXIT_STATUS=0
#shellcheck disable=SC1091
#shellcheck disable=SC1090
source "${PROJECT_ROOT}"/services/deploy/release-utils.source

#shellcheck disable=SC1091
source ./release-config.source
echo "Verifying pip package deployment using Release Params:"
cat ./release-config.source

# Python 2.x has problems with the - character as a deployment tag
# So we strip it.
PYTHON_DEPLOY_TO_TAG=${DEPLOY_TO_TAG//-/}
#PYTHON_DEPLOY_TO_TAG=${DEPLOY_TO_TAG}
# There may also be problems with the _ character
# If so, uncomment the following.
#PYTHON_DEPLOY_TO_TAG=${DEPLOY_TO_TAG//_/}
export VERSION="${PYTHON_DEPLOY_TO_TAG}"
INFO "Verifying Version: ${VERSION}"

function main {
  mkdir -p "${PROJECT_ROOT}"/pip-release-verify
  pushd "${PROJECT_ROOT}"/pip-release-verify || exit 1

  #deploy-to-stage-internal-verify
  deploy-to-stage-verify
  #deploy-to-github-prod-verify
  #deploy-to-prod-verify
  popd || exit 1
}

function deploy-to-stage-internal-verify {
  echo "deploy-to-stage-internal-verify not yet Implemented"

}

function deploy-to-stage-verify {
  PUBLISH_FROM_NAME=categories-sdk
  PACKAGE_NAME=release-candidate-"${PUBLISH_FROM_NAME}"
  PACKAGE_URL=https://test.pypi.org/simple/

  INFO "Verifying install of ${PACKAGE_NAME} Version: ${VERSION} Successfully Installed from ${PACKAGE_URL} remote"
  if pip install -i ${PACKAGE_URL} "${PACKAGE_NAME}"=="${VERSION}" ; then
    PASS "Python Package ${PACKAGE_NAME} Version: ${VERSION} Successfully Installed from ${PACKAGE_URL}"
  else
    ERROR "Failed to Install Python Package ${PACKAGE_NAME} Version: ${VERSION} from ${PACKAGE_URL}"
    debug
    export EXIT_STATUS=1
  fi

}

function deploy-to-github-prod-verify {
  echo "deploy-to-github-prod-verify not yet Implemented"

}

function deploy-to-prod-verify {
  PUBLISH_FROM_NAME=categories-sdk
  PACKAGE_NAME=release-candidate-"${PUBLISH_FROM_NAME}"
  PACKAGE_URL=https://pypi.org/

  INFO "Verifying install of ${PACKAGE_NAME} Version: ${VERSION} Successfully Installed from ${PACKAGE_URL} remote"
  if pip install "${PACKAGE_NAME}"=="${VERSION}" ; then
    PASS "Python Package ${PACKAGE_NAME} Version: ${VERSION} Successfully Installed from ${PACKAGE_URL}"
  else
    ERROR "Failed to Install Python Package ${PACKAGE_NAME} Version: ${VERSION} from ${PACKAGE_URL}"
    debug
    export EXIT_STATUS=1
  fi

}

function debug {
  #env |grep -v PASSWORD
  echo
  pwd
  ls -lah
  git status
  git log --oneline --decorate -n 15
  echo
  #cat > "${HOME}/.npmrc"
}

if [[ "${BASH_SOURCE[0]}" == "${0}" ]]; then
  main "$@"
else
  echo "This file is designed to be executed, not sourced."
fi

exit "${EXIT_STATUS}"
