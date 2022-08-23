#!/bin/bash
#export TERM=ansi
#export TERM=xterm-color
export PROJECT_ROOT=${PROJECT_ROOT:-"/home/circleci/project"}
export EXIT_STATUS=0
export PACKAGE_NAME=categories-javascript-client-sdk

#shellcheck disable=SC1091
#shellcheck disable=SC1090
source "${PROJECT_ROOT}"/services/deploy/release-utils.source

#shellcheck disable=SC1091
source ./release-config.source
echo "Verifying npm package deployment using Release Params:"
cat ./release-config.source
if [[ "${SUPPRESS_NPM}" == "true" ]]; then
  WARN "Suppressing npm debloyments. See manage-package-deployments.sh"
  exit
fi

export VERSION="${DEPLOY_TO_TAG}"

function main {
  mkdir -p "${PROJECT_ROOT}"/verify/npm-release-verify
  pushd "${PROJECT_ROOT}"/verify/npm-release-verify || exit 1

  NAMESPACES="vmm prism clustermgmt aiops iam storage"
  # Iterate the string variable using for loop
  for NAMESPACE in ${NAMESPACES}; do
    echo "NPM Deploying Namespace: ${NAMESPACE}"
    # HACK: Crude deployment manifest which needs to be reworked.
    if [ "${NAMESPACE}" = "storage" ] ; then
      VERSION="4.0.1-alpha.2"
      VERSION=${DEPLOY_TO_TAG}-${VERSION/./-}
    else
      VERSION="4.0.1-alpha.1"
      VERSION=${DEPLOY_TO_TAG}-${VERSION/./-}
    fi

    deploy-to-stage-internal-verify "${NAMESPACE}" "${VERSION}"
    deploy-to-stage-verify "${NAMESPACE}" "${VERSION}"
    deploy-to-github-prod-verify "${NAMESPACE}" "${VERSION}"
    #deploy-to-prod-verify

  done

  popd || exit 1
}

function deploy-to-stage-internal-verify {
  set-github-internal-npm-credentials
  NAMESPACE=$1
  VERSION=$2
  PACKAGE_NAME=${NAMESPACE}-js-client

  mkdir -p ./stage-github-internal-verify-"${NAMESPACE}"
  pushd ./stage-github-internal-verify-"${NAMESPACE}" || exit 1

  export PACKAGE=@nutanix-release-engineering/"${PACKAGE_NAME}"@"${VERSION}"
  if npm install "${PACKAGE}"; then
    PASS "NPM Package ${PACKAGE} Successfully Installed from Github Internal"
  else
    ERROR "Failed to Install NPM Package ${PACKAGE} from Github Internal"
    debug
    export EXIT_STATUS=1
  fi
  popd || exit 1

}

function deploy-to-stage-verify {
  set-npmjs-npm-credentials
  NAMESPACE=$1
  VERSION=$2
  PACKAGE_NAME=${NAMESPACE}-js-client

  mkdir -p ./stage-verify-"${NAMESPACE}"
  pushd ./stage-verify-"${NAMESPACE}" || exit 1
  export PACKAGE=@nutanix-scratch/"${PACKAGE_NAME}"@"${VERSION}"
  if npm install "${PACKAGE}"; then
    PASS "NPM Package ${PACKAGE} Successfully Installed from @nutanix-scratch"
  else
    ERROR "Failed to Install NPM Package ${PACKAGE} from @nutanix-scratch"
    debug
    export EXIT_STATUS=1
  fi
  popd || exit 1

}

function deploy-to-github-prod-verify {
  set-github-npm-credentials
  NAMESPACE=$1
  VERSION=$2
  PACKAGE_NAME=${NAMESPACE}-js-client

  mkdir -p ./prod-github-verify-"${NAMESPACE}"
  pushd ./prod-github-verify-"${NAMESPACE}" || exit 1

  #echo "${EXTERNAL_GITHUB_NPM_PULL_CREDENTIALS}" > "${HOME}/.npmrc"

  # HACK: Using orichter for testing. Proper domain should be @nutanix
  export PACKAGE=@orichter/"${PACKAGE_NAME}"@"${VERSION}"
  if npm install "${PACKAGE}"; then
    PASS "NPM Package ${PACKAGE} Successfully Installed from Github"
  else
    ERROR "Failed to Install NPM Package ${PACKAGE} from Github"
    debug
    export EXIT_STATUS=1
  fi
  popd || exit 1

}

function deploy-to-prod-verify {
  set-npmjs-npm-credentials


  mkdir -p ./prod-verify-"${NAMESPACE}"
  pushd ./prod-verify-"${NAMESPACE}" || exit 1
  export PACKAGE=@nutanix-api/"${PACKAGE_NAME}"@"${VERSION}"
  if npm install "${PACKAGE}"; then
    PASS "NPM Package ${PACKAGE} Successfully Installed from @nutanix-scratch"
  else
    #ERROR "Failed to Install NPM Package ${PACKAGE} from @nutanix-scratch"
    WARN "Prod Deployment not yet implemented."
    #debug
    #export EXIT_STATUS=1
    exit 1
  fi
  popd || exit 1

}

function set-github-npm-credentials {
  #//npm.pkg.github.com/:_authToken='"${GITHUB_PACKAGE_READ_TOKEN}"'
  #echo '@orichter:registry=https://npm.pkg.github.com/orichter
  echo '@orichter:registry=https://npm.pkg.github.com/orichter
//npm.pkg.github.com/:_authToken='"${GITHUB_PACKAGE_READ_TOKEN}"'
strict-ssl=false' > "${HOME}/.npmrc"
}

function set-github-internal-npm-credentials {
  #//npm.pkg.github.com/:_authToken='"${GITHUB_PACKAGE_READ_TOKEN}"'
  #echo '@orichter:registry=https://npm.pkg.github.com/orichter
  echo '@nutanix-release-engineering:registry=https://npm.pkg.github.com/nutanix-release-engineering
//npm.pkg.github.com/:_authToken='"${GITHUB_PACKAGE_READ_TOKEN}"'
strict-ssl=false' > "${HOME}/.npmrc"
}

function set-npmjs-npm-credentials {
  export EXTERNAL_NPM_STAGE_PULL_CREDENTIALS='@nutanix-scratch:registry=registry=https://registry.npmjs.org/
strict-ssl=false'  > "${HOME}/.npmrc"
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
