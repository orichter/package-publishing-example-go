#!/bin/bash
#export TERM=ansi
#export TERM=xterm-color
export PROJECT_ROOT=${PROJECT_ROOT:-"/home/circleci/project"}
export EXIT_STATUS=0
#shellcheck disable=SC1091
#shellcheck disable=SC1090
source "${PROJECT_ROOT}"/services/deploy/release-utils.source

#shellcheck disable=SC1091
#shellcheck disable=SC1090
source "${PROJECT_ROOT}"/release-config.source
echo "Verifying pip package deployment using Release Params:"
cat "${PROJECT_ROOT}"/release-config.source

# Python 2.x has problems with the - character as a deployment tag
# So we strip it.
PYTHON_DEPLOY_TO_TAG=${DEPLOY_TO_TAG//-/}
PYTHON_DEPLOY_TO_TAG=${PYTHON_DEPLOY_TO_TAG//v/}
#PYTHON_DEPLOY_TO_TAG=${DEPLOY_TO_TAG}
# There may also be problems with the _ character
# If so, uncomment the following.
#PYTHON_DEPLOY_TO_TAG=${DEPLOY_TO_TAG//_/}
export VERSION="${PYTHON_DEPLOY_TO_TAG}"
INFO "Verifying Version: ${VERSION}"

function main {
  mkdir -p "${PROJECT_ROOT}"/verify/pip-release-verify
  pushd "${PROJECT_ROOT}"/verify/pip-release-verify || exit 1

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
  export PUBLISH_FROM_NAME=categories-sdk
  export PACKAGE_NAME=release-candidate-"${PUBLISH_FROM_NAME}"
  export PACKAGE_URL=https://test.pypi.org/simple/
  pip3 config set global.no-cache-dir false

  i=0
  # Due to artifactory caching delays, we expect a single verification failure.
  # So we throw away the output of the first call to avoid error logging.
  deploy-to-test-pypi > /dev/null 2>&1
  until [ $i -gt 99 ]
  do
    ((i=i+1))
    if deploy-to-test-pypi ; then
      echo "Verification Succeeded"
      return
    else
      INFO "Trying again in 10 seconds."
      echo "Number of Retries: $i"
      sleep 10
      #deploy-to-stage-verify
    fi
  done
  debug
  EXIT_STATUS=1
}

function deploy-to-test-pypi {
  #poetry cache clear --all .
  #rm -rf "${HOME}"/.cache/pip
  #mv "${HOME}"/.poetry "${HOME}"/.poetry.old
  #pip3 list
  #pwd
  #ls -lah "${HOME}"
  #ls -lah
  #pip3 install --upgrade pip
  #pip3 uninstall -y "${PUBLISH_FROM_NAME}"
  #pip3 cache dir
  #pip3 cache purge
  #pip3 list
  # Give time for pip publishing indices to update
  #sleep 300
  #INFO "Waiting"
  #sleep 300

  INFO "Verifying install of ${PACKAGE_NAME} Version: ${VERSION} from ${PACKAGE_URL} remote"
  if pip3 install --no-deps --no-cache-dir --upgrade --index-url ${PACKAGE_URL} --extra-index-url https://pypi.org/simple/ "${PACKAGE_NAME}"=="${VERSION}" ; then
    PASS "Python Package ${PACKAGE_NAME} Version: ${VERSION} Successfully Installed from ${PACKAGE_URL}"
  else
    ERROR "Failed to Install Python Package ${PACKAGE_NAME} Version: ${VERSION} from ${PACKAGE_URL}"
    #debug
    return 1
    #export EXIT_STATUS=1
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
  ls -lah "${HOME}"
  ls -lah "${HOME}"/.cache
  ls -lah "${HOME}"/.pyenv
  ls -lah "${HOME}"/.poetry
}

if [[ "${BASH_SOURCE[0]}" == "${0}" ]]; then
  main "$@"
else
  echo "This file is designed to be executed, not sourced."
fi

exit "${EXIT_STATUS}"
