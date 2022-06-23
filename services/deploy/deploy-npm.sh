#!/bin/bash
#export TERM=ansi
export TERM=xterm-color
PROJECT_ROOT=${PROJECT_ROOT:-"/home/circleci/project"}
EXIT_STATUS=0
#shellcheck disable=SC1091
#shellcheck disable=SC1090
source "${PROJECT_ROOT}"/services/deploy/release-utils.source

function debug {
  env |grep -v PASSWORD
  pwd
  ls -lah
  ls -lah "${PROJECT_ROOT}"/services
  ls -lah "${PROJECT_ROOT}"/services/deploy
}

function main {
  pushd "${PROJECT_ROOT}"/npm-release-verify || exit 1

  #shellcheck disable=SC1091
  #shellcheck disable=SC1090
  source "${PROJECT_ROOT}"/release-config.source
  INFO "Deploying npm package using Release Params:"
  cat "${PROJECT_ROOT}"/release-config.source
  echo

  pushd prod-github-package || exit 1
  export TEST_DEPLOYMENT_TAG=0.0."${DEPLOYMENT_TAG//\./-}"
  #npm version "${DEPLOYMENT_TAG}"
  npm version "${TEST_DEPLOYMENT_TAG}"
  echo "Publishing package.json:"
  cat package.json

  echo "@orichter:registry=https://npm.pkg.github.com/orichter" >> "${HOME}"/.npmrc
  echo "//npm.pkg.github.com/:_authToken=${PASSWORD_PUBLISH_NPM_GITHUB}" >> "${HOME}"/.npmrc
  npm publish --access public || exit 1
  popd || exit 1

  #echo "//registry.npmjs.org/:_authToken=${PASSWORD_PUBLISH_NPM}" > "${HOME}"/.npmrc
  #npm publish --access public || exit 1

  popd || exit 1

}

if [[ "${BASH_SOURCE[0]}" == "${0}" ]]; then
  main "$@"
else
  echo "This file is designed to be executed, not sourced."
fi

exit "${EXIT_STATUS}"
