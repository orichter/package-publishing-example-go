#!/bin/bash
#export TERM=ansi
export TERM=xterm-color
export PROJECT_ROOT=${PROJECT_ROOT:-"/home/circleci/project"}
EXIT_STATUS=0
#shellcheck disable=SC1091
#shellcheck disable=SC1090
source "${PROJECT_ROOT}"/services/deploy/release-utils.source

#shellcheck disable=SC1091
#shellcheck disable=SC1090
source "${PROJECT_ROOT}"/release-config.source
if [[ "${SUPPRESS_NPM}" == "true" ]]; then
  WARN "Suppressing npm debloyments. See manage-package-deployments.sh"
  exit
fi

export VERSION="${DEPLOY_TO_TAG}"
INFO "Deploying npm packages using Release Params:"
cat "${PROJECT_ROOT}"/release-config.source
INFO "Version: ${VERSION}"
echo

function main {
  # DEPLOY_TO_TAG="0.1.4-0-1-rc1"
  # DEPLOY_FROM_TAG: "4.0.1-alpha.1"

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

    deploy-to-stage-internal "${NAMESPACE}" "${VERSION}"
    deploy-to-stage "${NAMESPACE}" "${VERSION}"
    deploy-to-github-prod "${NAMESPACE}" "${VERSION}"
    #deploy-to-prod "${NAMESPACE}" "${VERSION}"

  done

  PASS "Successful Deployments can be found at:"
  cat "${PROJECT_ROOT}"/verify/npm-release-verify/successful-deployments.txt
  if test -f "${PROJECT_ROOT}/npm-release-verify/failed-deployments.txt"; then
    ERROR "Failed Deployments to:"
    cat "${PROJECT_ROOT}"/verify/npm-release-verify/failed-deployments.txt
  fi

}

function deploy-to-stage-internal {
  NAMESPACE=$1
  VERSION=$2
  PACKAGE_NAME=${NAMESPACE}-js-client
  pushd "${PROJECT_ROOT}"/verify/npm-release-verify || exit 1

  cp -rf package-"${NAMESPACE}" stage-package-internal-"${NAMESPACE}"

  pushd stage-package-internal-"${NAMESPACE}" || exit 1
  PUBLISH_FROM=@nutanix-core/"${PACKAGE_NAME}"
  PUBLISH_TO=@nutanix-release-engineering/"${PACKAGE_NAME}"

  sed -i "s|${PUBLISH_FROM}|${PUBLISH_TO}|g" package.json

  #PUBLISH_FROM_URL=https://npm.pkg.github.com
  #PUBLISH_TO_URL=https://registry.npmjs.org
  #sed -i "s|${PUBLISH_FROM_URL}|${PUBLISH_TO_URL}|g" package.json

  PUBLISH_FROM_REPO=git://github.com/nutanix-core/ntnx-api-javascript-sdk-external.git
  PACKAGE_URL=https://github.com/nutanix-release-engineering/experiments-example-github-package-npm
  PUBLISH_TO_REPO=git+"${PACKAGE_URL}".git
  # It is unclear if /packages/1509289 is static, so it may need to be updated or removed.
  PACKAGE_URL="${PACKAGE_URL}"/packages/1509289

  sed -i "s|${PUBLISH_FROM_REPO}|${PUBLISH_TO_REPO}|g" package.json

  INFO "Internal Stage Package Config for ${PACKAGE_NAME}"
  cat package.json

  npm version "${VERSION}"
  echo "Publishing ${PACKAGE_NAME} to stage-package-internal package.json:"
  cat package.json

  echo "@orichter:registry=https://npm.pkg.github.com/orichter" > "${HOME}"/.npmrc
  echo "//npm.pkg.github.com/:_authToken=${PASSWORD_PUBLISH_NPM_GITHUB}" >> "${HOME}"/.npmrc
  INFO "https://github.com/nutanix-release-engineering/experiments-example-github-package-npm/packages/"

  if npm publish --access public ; then
    PASS "NPM Package ${PACKAGE_NAME} Successfully Deployed to ${PACKAGE_URL}"
    PASS "${PACKAGE_URL}" >> "${PROJECT_ROOT}"/verify/npm-release-verify/successful-deployments.txt
  else
    ERROR "Failed to Deploy Python Package ${PACKAGE_NAME} to ${PACKAGE_URL}"
    ERROR "${PACKAGE_URL}" >> "${PROJECT_ROOT}"/verify/npm-release-verify/failed-deployments.txt
    debug
    export EXIT_STATUS=1
  fi

  popd || exit 1

  popd || exit 1

}

function deploy-to-stage {
  NAMESPACE=$1
  VERSION=$2
  PACKAGE_NAME=${NAMESPACE}-js-client

  pushd "${PROJECT_ROOT}"/verify/npm-release-verify || exit 1

  cp -rf package-"${NAMESPACE}" stage-package-"${NAMESPACE}"

  pushd stage-package-"${NAMESPACE}" || exit 1
  PUBLISH_FROM=@nutanix-core/"${PACKAGE_NAME}"
  PUBLISH_TO=@nutanix-scratch/"${PACKAGE_NAME}"
  PACKAGE_URL=https://www.npmjs.com/package/"${PUBLISH_TO}"
  sed -i "s|${PUBLISH_FROM}|${PUBLISH_TO}|g" package.json

  PUBLISH_FROM_URL=https://npm.pkg.github.com
  PUBLISH_TO_URL=https://registry.npmjs.org
  sed -i "s|${PUBLISH_FROM_URL}|${PUBLISH_TO_URL}|g" package.json

  PUBLISH_FROM_REPO=git://github.com/nutanix-core/ntnx-api-javascript-sdk-external.git
  #PACKAGE_URL=https://github.com/nutanix-release-engineering/experiments-example-github-package-npm
  #PUBLISH_TO_REPO=git+"${PACKAGE_URL}".git
  PUBLISH_TO_REPO=git+https://github.com/orichter/package-publishing-examples.git

  sed -i "s|${PUBLISH_FROM_REPO}|${PUBLISH_TO_REPO}|g" package.json

  INFO "Stage Package Config for ${PACKAGE_NAME}"
  cat package.json

  npm version "${VERSION}"
  echo "Publishing ${PACKAGE_NAME} stage-package package.json:"
  cat package.json

  echo "//registry.npmjs.org/:_authToken=${PASSWORD_PUBLISH_NPM}" > "${HOME}"/.npmrc

  if npm publish --access public ; then
    PASS "NPM Package ${PACKAGE_NAME} Successfully Deployed to ${PACKAGE_URL}"
    PASS "${PACKAGE_URL}" >> "${PROJECT_ROOT}"/verify/npm-release-verify/successful-deployments.txt
  else
    ERROR "Failed to Deploy Python Package ${PACKAGE_NAME} to ${PACKAGE_URL}"
    ERROR "${PACKAGE_URL}" >> "${PROJECT_ROOT}"/verify/npm-release-verify/failed-deployments.txt
    debug
    export EXIT_STATUS=1
  fi

  popd || exit 1

  popd || exit 1

}

function deploy-to-github-prod {
  NAMESPACE=$1
  VERSION=$2
  PACKAGE_NAME=${NAMESPACE}-js-client
  pushd "${PROJECT_ROOT}"/verify/npm-release-verify || exit 1

  cp -rf package-"${NAMESPACE}" prod-github-package-"${NAMESPACE}"

  pushd prod-github-package-"${NAMESPACE}" || exit 1
  PUBLISH_FROM=@nutanix-core/"${PACKAGE_NAME}"
  # HACK: Currently publishing to orichter for testing. Need to change to nutanix scope
  PUBLISH_TO=@orichter/"${PACKAGE_NAME}"

  sed -i "s|${PUBLISH_FROM}|${PUBLISH_TO}|g" package.json

  #PUBLISH_FROM_URL=https://npm.pkg.github.com
  #PUBLISH_TO_URL=https://registry.npmjs.org
  #sed -i "s|${PUBLISH_FROM_URL}|${PUBLISH_TO_URL}|g" package.json

  PUBLISH_FROM_REPO=git://github.com/nutanix-core/ntnx-api-javascript-sdk-external.git
  PACKAGE_URL=https://github.com/orichter/package-publishing-examples
  PUBLISH_TO_REPO=git+"${PACKAGE_URL}".git
  # It is unclear if /packages/1509289 is static, so it may need to be updated or removed.
  #PACKAGE_URL="${PACKAGE_URL}"/packages/TBD

  sed -i "s|${PUBLISH_FROM_REPO}|${PUBLISH_TO_REPO}|g" package.json

  INFO "Github Prod Package Config for ${PACKAGE_NAME}"
  cat package.json

  npm version "${VERSION}"
  echo "Publishing ${PACKAGE_NAME} prod-github-package package.json:"
  cat package.json

  echo "@orichter:registry=https://npm.pkg.github.com/orichter" > "${HOME}"/.npmrc
  echo "//npm.pkg.github.com/:_authToken=${PASSWORD_PUBLISH_NPM_GITHUB}" >> "${HOME}"/.npmrc
  INFO "${PACKAGE_URL}"

  if npm publish --access public ; then
    PASS "NPM Package ${PACKAGE_NAME} Successfully Deployed to ${PACKAGE_URL}"
    PASS "${PACKAGE_URL}" >> "${PROJECT_ROOT}"/verify/npm-release-verify/successful-deployments.txt
  else
    ERROR "Failed to Deploy Python Package ${PACKAGE_NAME} to ${PACKAGE_URL}"
    ERROR "${PACKAGE_URL}" >> "${PROJECT_ROOT}"/verify/npm-release-verify/failed-deployments.txt
    debug
    export EXIT_STATUS=1
  fi

  popd || exit 1

  popd || exit 1

}

function deploy-to-prod {
  NAMESPACE=$1
  VERSION=$2
  PACKAGE_NAME=${NAMESPACE}-js-client

  pushd "${PROJECT_ROOT}"/verify/npm-release-verify || exit 1

  cp -rf package-"${NAMESPACE}" prod-package-"${NAMESPACE}"

  pushd prod-package-"${NAMESPACE}" || exit 1
  PUBLISH_FROM=@nutanix-core/"${PACKAGE_NAME}"
  PUBLISH_TO=@nutanix-api/"${PACKAGE_NAME}"
  PACKAGE_URL=https://www.npmjs.com/package/"${PUBLISH_TO}"
  sed -i "s|${PUBLISH_FROM}|${PUBLISH_TO}|g" package.json

  PUBLISH_FROM_URL=https://npm.pkg.github.com
  PUBLISH_TO_URL=https://registry.npmjs.org
  sed -i "s|${PUBLISH_FROM_URL}|${PUBLISH_TO_URL}|g" package.json

  PUBLISH_FROM_REPO=git://github.com/nutanix-core/ntnx-api-javascript-sdk-external.git
  PUBLISH_TO_REPO=git+https://github.com/orichter/package-publishing-examples.git
  sed -i "s|${PUBLISH_FROM_REPO}|${PUBLISH_TO_REPO}|g" package.json

  INFO "Npmjs.org Prod Package Config for ${PACKAGE_NAME}"
  cat package.json

  npm version "${VERSION}"
  echo "Publishing ${PACKAGE_NAME} prod-package package.json:"
  cat package.json

  echo "//registry.npmjs.org/:_authToken=${PASSWORD_PUBLISH_NPM}" > "${HOME}"/.npmrc

  echo "HACK: Currently not publishing  ${PACKAGE_NAME} to Prod"
  # HACK: npm publish to prod is currently disabled.
  # Uncomment the line below and comment the line above to publish to prod.
  #if npm publish --access public ; then
  #if npm publish ; then
  # HACK: Temporarily disable prod publish to review package.json
  if true ; then
    PASS "NPM Package ${PACKAGE_NAME} Successfully Deployed to ${PACKAGE_URL}"
    PASS "${PACKAGE_URL}" >> "${PROJECT_ROOT}"/verify/npm-release-verify/successful-deployments.txt
  else
    ERROR "Failed to Deploy Python Package ${PACKAGE_NAME} to ${PACKAGE_URL}"
    ERROR "${PACKAGE_URL}" >> "${PROJECT_ROOT}"/verify/npm-release-verify/failed-deployments.txt
    debug
    export EXIT_STATUS=1
  fi

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
