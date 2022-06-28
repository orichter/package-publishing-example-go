#!/bin/bash
#export TERM=ansi
export TERM=xterm-color
PROJECT_ROOT=${PROJECT_ROOT:-"/home/circleci/project"}
EXIT_STATUS=0
#shellcheck disable=SC1091
#shellcheck disable=SC1090
source "${PROJECT_ROOT}"/services/deploy/release-utils.source

#shellcheck disable=SC1091
#shellcheck disable=SC1090
source "${PROJECT_ROOT}"/release-config.source
INFO "Deploying npm packages using Release Params:"
cat "${PROJECT_ROOT}"/release-config.source
echo

export TEST_DEPLOYMENT_TAG=0.16."${DEPLOYMENT_TAG//\./-}"
export VERSION="${TEST_DEPLOYMENT_TAG}"
#export VERSION="${DEPLOYMENT_TAG}"

function main {
  deploy-to-stage-internal
  #deploy-to-stage
  #deploy-to-github-prod
  #deploy-to-prod
}

function deploy-to-stage-internal {
  pushd "${PROJECT_ROOT}"/npm-release-verify || exit 1

  cp -rf package stage-package-internal

  pushd stage-package-internal || exit 1
  PUBLISH_FROM=@nutanix-core/categories-javascript-client-sdk
  PUBLISH_TO=@nutanix-release-engineering/release-canadidate-javascript-sdk
  sed -i "s|${PUBLISH_FROM}|${PUBLISH_TO}|g" package.json

  #PUBLISH_FROM_URL=https://npm.pkg.github.com
  #PUBLISH_TO_URL=https://registry.npmjs.org
  #sed -i "s|${PUBLISH_FROM_URL}|${PUBLISH_TO_URL}|g" package.json

  PUBLISH_FROM_REPO=git://github.com/nutanix-core/ntnx-api-javascript-sdk-external.git
  PUBLISH_TO_REPO=git+https://github.com/nutanix-release-engineering/experiments-example-github-package-npm.git
  sed -i "s|${PUBLISH_FROM_REPO}|${PUBLISH_TO_REPO}|g" package.json

  INFO "Internal Stage Package Config"
  cat package.json

  npm version "${VERSION}"
  echo "Publishing stage-package-internal package.json:"
  cat package.json

  echo "@orichter:registry=https://npm.pkg.github.com/orichter" > "${HOME}"/.npmrc
  echo "//npm.pkg.github.com/:_authToken=${PASSWORD_PUBLISH_NPM_GITHUB}" >> "${HOME}"/.npmrc
  npm publish || exit 1
  popd || exit 1

  popd || exit 1

}

function deploy-to-stage {
  pushd "${PROJECT_ROOT}"/npm-release-verify || exit 1

  cp -rf package stage-package

  pushd stage-package || exit 1
  PUBLISH_FROM=@nutanix-core/categories-javascript-client-sdk
  PUBLISH_TO=@nutanix-scratch/release-canadidate-javascript-sdk
  sed -i "s|${PUBLISH_FROM}|${PUBLISH_TO}|g" package.json

  PUBLISH_FROM_URL=https://npm.pkg.github.com
  PUBLISH_TO_URL=https://registry.npmjs.org
  sed -i "s|${PUBLISH_FROM_URL}|${PUBLISH_TO_URL}|g" package.json

  PUBLISH_FROM_REPO=git://github.com/nutanix-core/ntnx-api-javascript-sdk-external.git
  PUBLISH_TO_REPO=git+https://github.com/orichter/package-publishing-examples.git
  sed -i "s|${PUBLISH_FROM_REPO}|${PUBLISH_TO_REPO}|g" package.json

  INFO "Stage Package Config"
  cat package.json

  npm version "${VERSION}"
  echo "Publishing stage-package package.json:"
  cat package.json

  echo "//registry.npmjs.org/:_authToken=${PASSWORD_PUBLISH_NPM}" > "${HOME}"/.npmrc
  npm publish --access public || exit 1
  popd || exit 1

  popd || exit 1

}

function deploy-to-github-prod {
  pushd "${PROJECT_ROOT}"/npm-release-verify || exit 1

  cp -rf package prod-github-package

  pushd prod-github-package || exit 1
  PUBLISH_FROM=@nutanix-core/categories-javascript-client-sdk
  PUBLISH_TO=@orichter/release-candidate-javascript-sdk
  sed -i "s|${PUBLISH_FROM}|${PUBLISH_TO}|g" package.json

  #PUBLISH_FROM_URL=https://npm.pkg.github.com
  #PUBLISH_TO_URL=https://registry.npmjs.org
  #sed -i "s|${PUBLISH_FROM_URL}|${PUBLISH_TO_URL}|g" package.json

  PUBLISH_FROM_REPO=git://github.com/nutanix-core/ntnx-api-javascript-sdk-external.git
  PUBLISH_TO_REPO=git+https://github.com/orichter/package-publishing-examples.git
  sed -i "s|${PUBLISH_FROM_REPO}|${PUBLISH_TO_REPO}|g" package.json

  INFO "Github Prod Package Config"
  cat package.json

  npm version "${VERSION}"
  echo "Publishing prod-github-package package.json:"
  cat package.json

  echo "@orichter:registry=https://npm.pkg.github.com/orichter" > "${HOME}"/.npmrc
  echo "//npm.pkg.github.com/:_authToken=${PASSWORD_PUBLISH_NPM_GITHUB}" >> "${HOME}"/.npmrc
  npm publish --access public || exit 1

  popd || exit 1

}

function deploy-to-prod {
  pushd "${PROJECT_ROOT}"/npm-release-verify || exit 1

  cp -rf package prod-package

  pushd prod-package || exit 1
  PUBLISH_FROM=@nutanix-core/categories-javascript-client-sdk
  PUBLISH_TO=@nutanix-api/javascript-client-sdk
  sed -i "s|${PUBLISH_FROM}|${PUBLISH_TO}|g" package.json

  PUBLISH_FROM_URL=https://npm.pkg.github.com
  PUBLISH_TO_URL=https://registry.npmjs.org
  sed -i "s|${PUBLISH_FROM_URL}|${PUBLISH_TO_URL}|g" package.json

  PUBLISH_FROM_REPO=git://github.com/nutanix-core/ntnx-api-javascript-sdk-external.git
  PUBLISH_TO_REPO=git+https://github.com/orichter/package-publishing-examples.git
  sed -i "s|${PUBLISH_FROM_REPO}|${PUBLISH_TO_REPO}|g" package.json

  INFO "Npmjs.org Prod Package Config"
  cat package.json

  npm version "${VERSION}"
  echo "Publishing prod-package package.json:"
  cat package.json

  echo "//registry.npmjs.org/:_authToken=${PASSWORD_PUBLISH_NPM}" > "${HOME}"/.npmrc

  echo "Currently not publishing to Prod"
  # Uncomment the line below and comment the line above to publish to prod.
  #npm publish --access public || exit 1

  popd || exit 1

}

function debug {
  env |grep -v PASSWORD
  pwd
  ls -lah
  ls -lah "${PROJECT_ROOT}"/services
  ls -lah "${PROJECT_ROOT}"/services/deploy
}

if [[ "${BASH_SOURCE[0]}" == "${0}" ]]; then
  main "$@"
  exit "${EXIT_STATUS}"
else
  WARN "This file is designed to be executed, not sourced."
fi
