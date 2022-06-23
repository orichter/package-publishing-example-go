#!/bin/bash
#export TERM=ansi
#export TERM=xterm-color
PROJECT_ROOT=${PROJECT_ROOT:-"/home/circleci/project"}
#shellcheck disable=SC1091
#shellcheck disable=SC1090
source "${PROJECT_ROOT}"/services/deploy/release-utils.source

export EXTERNAL_GITHUB_NPM_PULL_CREDENTIALS='@orichter:registry=https://npm.pkg.github.com/orichter
//npm.pkg.github.com/:_authToken='"${GITHUB_PACKAGE_READ_TOKEN}"'
strict-ssl=false' > "${HOME}/.npmrc"

#shellcheck disable=SC1091
source ./release-config.source
echo "Verifying npm package deployment using Release Params:"
cat ./release-config.source

function main {
  mkdir -p "${PROJECT_ROOT}"/npm-release-verify
  pushd "${PROJECT_ROOT}"/npm-release-verify || exit 1

  echo "${EXTERNAL_GITHUB_NPM_PULL_CREDENTIALS}" > "${HOME}/.npmrc"

  export TEST_DEPLOYMENT_TAG=0.1."${DEPLOYMENT_TAG//\./-}"
  #export PACKAGE=@orichter/package-publishing-examples@"${DEPLOYMENT_TAG}"
  export PACKAGE=@orichter/package-publishing-examples@"${TEST_DEPLOYMENT_TAG}"
  #if npm install @orichter/package-publishing-examples@"${DEPLOYMENT_TAG}"; then
  if npm install "${PACKAGE}"; then
    PASS "NPM Package ${PACKAGE} Successfully Installed"
  else
    ERROR "Failed to Install NPM Package ${PACKAGE}"
    debug
    exit 1
  fi
  popd || exit 1
}

function debug {
  #env |grep -v PASSWORD
  echo
  pwd
  ls -lah
  git status
  git log --oneline --decorate -n 15
  echo
}

if [[ "${BASH_SOURCE[0]}" == "${0}" ]]; then
  main "$@"
else
  echo "This file is designed to be executed, not sourced."
fi

exit "${EXIT_STATUS}"
