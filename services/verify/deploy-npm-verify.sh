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
echo "Verifying npm package deployment using Release Params:"
cat ./release-config.source

#shellcheck disable=SC1091
source ./services/deploy/deploy-npm.sh # This line is unnecessary unless using the following variables
#export TEST_DEPLOYMENT_TAG=0.6."${DEPLOYMENT_TAG//\./-}"
#export VERSION="${TEST_DEPLOYMENT_TAG}"
#export VERSION="${DEPLOYMENT_TAG}"


function main {
  mkdir -p "${PROJECT_ROOT}"/npm-release-verify
  pushd "${PROJECT_ROOT}"/npm-release-verify || exit 1

  deploy-to-stage-internal-verify
  #deploy-to-stage-verify
  #deploy-to-github-prod-verify
  #deploy-to-prod-verify
  popd || exit 1
}

function deploy-to-stage-internal-verify {
  set-github-npm-credentials

  mkdir -p ./stage-github-internal-verify
  pushd ./stage-github-internal-verify || exit 1

  export PACKAGE=@nutanix-release-engineering/release-canadidate-javascript-sdk@"${VERSION}"
  if npm install "${PACKAGE}"; then
    PASS "NPM Package ${PACKAGE} Successfully Installed from Github Internal"
  else
    ERROR "Failed to Install NPM Package ${PACKAGE} from Github Internal"
    debug
    export EXIT_STATUS=1
    exit 1
  fi
  popd || exit 1

}

function deploy-to-stage-verify {
  set-npmjs-npm-credentials

  mkdir -p ./stage-verify
  pushd ./stage-verify || exit 1
  export PACKAGE=@nutanix-scratch/release-canadidate-javascript-sdk@"${VERSION}"
  if npm install "${PACKAGE}"; then
    PASS "NPM Package ${PACKAGE} Successfully Installed from @nutanix-scratch"
  else
    ERROR "Failed to Install NPM Package ${PACKAGE} from @nutanix-scratch"
    debug
    export EXIT_STATUS=1
    exit 1
  fi
  popd || exit 1

}

function deploy-to-github-prod-verify {
  set-github-npm-credentials

  mkdir -p ./prod-github-verify
  pushd ./prod-github-verify || exit 1

  echo "${EXTERNAL_GITHUB_NPM_PULL_CREDENTIALS}" > "${HOME}/.npmrc"

  export PACKAGE=@orichter/release-canadidate-javascript-sdk@"${VERSION}"
  if npm install "${PACKAGE}"; then
    PASS "NPM Package ${PACKAGE} Successfully Installed from Github"
  else
    ERROR "Failed to Install NPM Package ${PACKAGE} from Github"
    debug
    export EXIT_STATUS=1
    exit 1
  fi
  popd || exit 1

}

function deploy-to-prod-verify {
  set-npmjs-npm-credentials
  mkdir -p ./prod-verify
  pushd ./prod-verify || exit 1
  export PACKAGE=@nutanix-api/javascript-sdk@"${VERSION}"
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
  #//npm.pkg.github.com/:_authToken='"${PASSWORD_PUBLISH_NPM_GITHUB}"'
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
