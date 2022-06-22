#!/bin/bash
#export TERM=ansi
export TERM=xterm-color

function debug {
  env |grep -v PASSWORD
  pwd
  ls -lah
  ls -lah ./services
  ls -lah ./services/deploy
}

#shellcheck disable=SC1091
source ./release-config.source
echo "Deploying npm package using Release Params:"
cat ./release-config.source
echo
pushd ./npm/hello-world/ || exit 1
npm version "${DEPLOYMENT_TAG}"
echo "Publishing package.json:"
cat ./package.json
echo "//registry.npmjs.org/:_authToken=${PASSWORD_PUBLISH_NPM}" > "${HOME}"/.npmrc
npm publish --access public || exit 1
popd || exit 1