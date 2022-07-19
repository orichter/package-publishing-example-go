#!/bin/bash
#export TERM=ansi
export TERM=xterm-color
PROJECT_ROOT=${PROJECT_ROOT:-"/home/circleci/project"}
EXIT_STATUS=0

export INTERNAL_NPM_PULL_CREDENTIALS='@nutanix-core:registry=https://npm.pkg.github.com/nutanix-core
//npm.pkg.github.com/:_authToken='"${GITHUB_PACKAGE_READ_TOKEN}"'
strict-ssl=false'

#shellcheck disable=SC1091
#shellcheck disable=SC1090
source "${PROJECT_ROOT}"/services/deploy/release-utils.source

#shellcheck disable=SC1091
#shellcheck disable=SC1090
source "${PROJECT_ROOT}"/release-config.source
export VERSION="${DEPLOY_FROM_TAG}"
echo "Verifying Hacks using Release Params:"
cat "${PROJECT_ROOT}"/release-config.source
INFO "Version: ${VERSION}"
echo

function main {
  # The following line is used for testing the release pipeline only.
  # It should be commented out for actual deployments.
  check-hacks
}

function check-hacks {
  # It is uncelar why an == fails here, so a regex search is required.
  pushd "${PROJECT_ROOT}" || exit 1
  if [[ $(grep -R HACK:|wc -l) -eq 2 ]] ; then
    PASS "All Hacks Removed."
  else
    ERROR "There are still hacks in the code."
    grep -R HACK:|grep -v grep
    #debug
    EXIT_STATUS=1
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
}

if [[ "${BASH_SOURCE[0]}" == "${0}" ]]; then
  main "$@"
else
  echo "This file is designed to be executed, not sourced."
fi

exit "${EXIT_STATUS}"
