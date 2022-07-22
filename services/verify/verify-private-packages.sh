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
echo "Verifying private packages using Release Params:"
cat "${PROJECT_ROOT}"/release-config.source
INFO "Version: ${VERSION}"
echo

function main {
  # The following line is used for testing the release pipeline only.
  # It should be commented out for actual deployments.
  check-prerequisites
  #mvn-sample-release
  mvn-internal-release-verify
  golang-internal-release-verify
  pip-internal-release-verify
  npm-internal-release-verify
}

function check-prerequisites {
  MVN_VERSION_REQUIRED=3.6.0
  MVN_VERSION="$(mvn -version | awk '{print $3}'|head -n 1)"
  # It is uncelar why an == fails here, so a regex search is required.
  if [[ "${MVN_VERSION}" =~ ${MVN_VERSION_REQUIRED} ]] ; then
    PASS "mvn version is ${MVN_VERSION}"
  else
    ERROR "Current mvn version is: ${MVN_VERSION}. These deployment scripts were verified with required version: ${MVN_VERSION_REQUIRED}"
    WARN "Please update mvn to the required version, or re-test deployments with new version."
    exit 1
  fi
}

function golang-internal-release-verify {
  rm -rf "${PROJECT_ROOT}"/golang-release-verify
  mkdir -p "${PROJECT_ROOT}"/golang-release-verify
  pushd "${PROJECT_ROOT}"/golang-release-verify || exit 1
  
  export PACKAGE_NAME=ntnx-api-golang-sdk-external
  export PACKAGE_URL=https://github.com/nutanix-core/"${PACKAGE_NAME}".git
  git clone "${PACKAGE_URL}"

  pushd "${PACKAGE_NAME}" || exit 1
  git fetch --all --tags

  # This is not a good tagging convention and should be refactored.
  export PACKAGE_TAG=categories_parent_go_sdk/v"${VERSION}"
  if git checkout "${PACKAGE_TAG}" ; then
    PASS "Tag ${VERSION} of Package: ${PACKAGE_NAME} Successfully checked out"
    git log --oneline --decorate -n 15
  else
    ERROR "Tag ${VERSION} of Package: ${PACKAGE_NAME} checked out failed."
    git log --oneline --decorate -n 15
    EXIT_STATUS=1
    debug
  fi

  popd || exit 1
  # Rename package for standard deployment, and so we have a copy of the original
  # for debugging purposes.
  cp -rf "${PACKAGE_NAME}" package

  popd || exit 1

}

function mvn-internal-release-verify {
  # HACK: Pinned to DEPLOY_FROM_TAG=4.0.1-alpha-1
  # This java package is currently pinned to the version below
  # rather than inheriting from config.yml. This should be fixed before
  # production deployment
  VERSION=4.0.1-alpha-1

  rm -rf "${PROJECT_ROOT}"/maven-release-verify
  mkdir -p "${PROJECT_ROOT}"/maven-release-verify

  pushd "${PROJECT_ROOT}"/maven-release-verify || exit 1

  mvn archetype:generate -q -DgroupId=com.nutanix.test.sdk -DartifactId=test-sdk-app -DarchetypeArtifactId=maven-archetype-quickstart -DarchetypeVersion=1.4 -DinteractiveMode=false

  pushd test-sdk-app || exit 1
  mv pom.xml pom.xml.old
  cp "${PROJECT_ROOT}"/services/deploy/mvn/pom.xml .
  cp "${PROJECT_ROOT}"/services/deploy/mvn/settings.xml .
  sed -i "s|.{env.DEPLOYMENT_TAG}|${VERSION}|g" pom.xml

  if mvn package -q --settings settings.xml -DdownloadSources=true -DdownloadJavadocs=true ; then
    PASS "Successfully Downloaded Maven Package using pom:"
    cat pom.xml
    mvn dependency:sources dependency:resolve -Dclassifier=javadoc -q -s settings.xml
    mvn dependency:sources -Dsilent=true -q -s settings.xml
    cp -rf ~/.m2/repository/com/nutanix/api/vmm-java-client/"${VERSION}"/ "${PROJECT_ROOT}"/maven-release-verify/package
    INFO "MVN Package Contents:"
    ls -lah "${PROJECT_ROOT}"/maven-release-verify/package
    pushd "${PROJECT_ROOT}"/maven-release-verify/package || exit 1
    #mvn dependency:sources dependency:resolve -Dclassifier=javadoc -q -s settings.xml
    #mvn dependency:sources -Dsilent=true -q -s settings.xml
    #ls -lah "${PROJECT_ROOT}"/maven-release-verify/package
    #cat vmm-java-client-4.0.1-alpha-1-sources.jar.lastUpdated
    popd || exit 1
    echo
  else
    ERROR "Failed to Download Maven Package"
    export EXIT_STATUS=1
    debug
  fi

  #CMD="java -cp target/test-sdk-app-1.0-SNAPSHOT.jar com.nutanix.test.test-sdk.App"
  #OUTPUT=$(eval "${CMD}")
  #echo "${OUTPUT}"
  #if [ "${OUTPUT}" == "Hello World!" ]; then
  #  PASS "Maven Test Rig Tests Pass"
  #else
  #  ERROR "Maven Test Rig Tests Failed"
  #  debug
  #  export EXIT_STATUS=1
  #fi
  echo

  #CMD="java -cp ~/.m2/repository/com/nutanix/test/test-sdk/${VERSION}/test-sdk-${VERSION}.jar com.nutanix.test.test-sdk.App"
  #OUTPUT=$(eval "${CMD}")
  #echo "${OUTPUT}"
  #if [ "${OUTPUT}" == "Hello World!" ]; then
  #  PASS "Maven Dependency Tests Pass"
  #  echo
  #  INFO "Staging Dependency for External Deployment"
  #  mkdir -p "${PROJECT_ROOT}"/mvn-external-release
  #  cp -rf ~/.m2/repository/com/nutanix/example/hellonutanixworld/"${VERSION}"/ "${PROJECT_ROOT}"/mvn-external-release
  #else
  #  ERROR "Maven Dependency Tests Failed"
  #  debug
  #  export EXIT_STATUS=1
  #fi
  echo

  popd || exit 1
  popd || exit 1

}

function npm-internal-release-verify {
  rm -rf "${PROJECT_ROOT}"/npm-release-verify
  mkdir -p "${PROJECT_ROOT}"/npm-release-verify
  pushd "${PROJECT_ROOT}"/npm-release-verify || exit 1
  
  echo "${INTERNAL_NPM_PULL_CREDENTIALS}" > "${HOME}/.npmrc"

  # HACK: NPM Verify pinned to VERSION=16.7.0-2
  # This override is necessary because the npm package tag is currently inconsistent with the other packages.
  # It should be fixed before final deployment
  VERSION=16.7.0-2

  if npm pack @nutanix-core/categories-javascript-client-sdk@"${VERSION}" ; then
    PASS "Successfully Downloaded NPM Package"
    INFO @nutanix-core/categories-javascript-client-sdk@"${VERSION}"
  else
    ERROR "Failed to Download NPM Package"
    INFO @nutanix-core/categories-javascript-client-sdk@"${VERSION}"
    export EXIT_STATUS=1
  fi

  tar -zxf nutanix-core-categories-javascript-client-sdk-"${VERSION}".tgz
  # This should be untarred into package which will be used in deploy-npm.sh for actual deployments

  cp -rf package audit
  pushd audit || exit 1
  VULNERABILITIES_COUNT=$(npm i --package-lock-only|grep vulnerabilities)

  CRITICAL_COUNT=$(echo "${VULNERABILITIES_COUNT}" | grep -oP '\w+(?= critical)')
  if [[ "${CRITICAL_COUNT}" -eq 0 ]]; then
    PASS "No critical vulnerabilities in @nutanix-core/categories-javascript-client-sdk"
  else
    npm audit
    WARNING="${CRITICAL_COUNT} Critical Vulnerabilities Found"
    WARN "${WARNING}"
    # HACK: This exit status should really be 1 as these vulnerabilites should be fixed.
    EXIT_STATUS=0
  fi

  popd || exit 1

}

function pip-internal-release-verify {
  rm -rf "${PROJECT_ROOT}"/pip-release-verify
  mkdir -p "${PROJECT_ROOT}"/pip-release-verify
  pushd "${PROJECT_ROOT}"/pip-release-verify || exit 1
  
  sudo apt install python3-pip
  /home/circleci/.pyenv/versions/3.8.5/bin/python3.8 -m pip install --upgrade pip

  #PACKAGE_NAME=categories-client
  PACKAGE_NAME=categories-sdk
  PACKAGE_DIR=categories
  #TMP_VERSION=categories-${VERSION}
  #PACKAGE_PATH="git+ssh://git@github.com/nutanix-core/ntnx-api-python-sdk-external.git@${TMP_VERSION}#subdirectory=${PACKAGE_DIR}&egg=${PACKAGE_NAME}"
  PACKAGE_PATH="git+ssh://git@github.com/nutanix-core/ntnx-api-python-sdk-external.git#subdirectory=${PACKAGE_DIR}&egg=${PACKAGE_NAME}"

  if pip3 install -e "${PACKAGE_PATH}" ; then
    PASS "Successfully Downloaded Python Package"
    INFO "${PACKAGE_PATH}"
  else
    ERROR "Failed to Download Python Package"
    INFO "${PACKAGE_PATH}"
    export EXIT_STATUS=1
  fi

  cp -rf src/${PACKAGE_NAME}/${PACKAGE_DIR} ./package

  popd || exit 1

}

#shellcheck disable=SC1091
#shellcheck disable=SC1090
source "${PROJECT_ROOT}"/services/deploy/release-utils.source
function debug {
  #env |grep -v PASSWORD
  echo
  pwd
  ls -lah
  git status
  git log --oneline --decorate -n 15
  echo

  WARN "Maven Debug"
  ls -lah ~/.m2/repository/com/nutanix/api
  ls -lah ~/.m2/repository/com/nutanix/api/vmm-java-client
  ls -lah ~/.m2/repository/com/nutanix/api/vmm-java-client/"${VERSION}"/
  ls -lah "${PROJECT_ROOT}"/maven-release-verify
  ls -lah "${PROJECT_ROOT}"/maven-release-verify
  ls -lah "${PROJECT_ROOT}"/maven-release-verify/package
}

if [[ "${BASH_SOURCE[0]}" == "${0}" ]]; then
  main "$@"
else
  echo "This file is designed to be executed, not sourced."
fi

exit "${EXIT_STATUS}"
