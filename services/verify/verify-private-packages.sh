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
function debug {
  #env |grep -v PASSWORD
  echo
  pwd
  ls -lah
  git status
  git log --oneline --decorate -n 15
  echo
}

#shellcheck disable=SC1091
#shellcheck disable=SC1090
source "${PROJECT_ROOT}"/release-config.source
echo "Verifying private packages using Release Params:"
cat "${PROJECT_ROOT}"/release-config.source

function main {
  # The following line is used for testing the release pipeline only.
  # It should be commented out for actual deployments.
  check-prerequisites
  #mvn-sample-release
  npm-internal-release-verify
  golang-internal-release-verify
  #mvn-internal-release-verify
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
  
  git clone https://github.com/nutanix-core/ntnx-api-golang-sdk-external
  pushd ntnx-api-golang-sdk-external || exit 1
  git fetch --all --tags
  if git checkout categories_parent_go_sdk/v"${DEPLOYMENT_TAG}" ; then
    PASS "Tag ${DEPLOYMENT_TAG} Successfully checked out"
    git log --oneline --decorate -n 15
  else
    ERROR "Tag ${DEPLOYMENT_TAG} checked out failed."
    git log --oneline --decorate -n 15
  fi

  popd || exit 1

  popd || exit 1

}

function mvn-internal-release-verify {
  rm -rf "${PROJECT_ROOT}"/maven-release-verify
  mkdir -p "${PROJECT_ROOT}"/maven-release-verify

  pushd "${PROJECT_ROOT}"/maven-release-verify || exit 1

  mvn archetype:generate -q -DgroupId=com.nutanix.test.sdk -DartifactId=test-sdk-app -DarchetypeArtifactId=maven-archetype-quickstart -DarchetypeVersion=1.4 -DinteractiveMode=false

  pushd test-sdk-app || exit 1
  mv pom.xml pom.xml.old
  cp "${PROJECT_ROOT}"/services/deploy/mvn/pom.xml .
  cp "${PROJECT_ROOT}"/services/deploy/mvn/settings.xml .

  mvn package -q --settings settings.xml

  ls -lah ~/.m2/repository/com/nutanix/nutanix-core/ntnx-api/categories/categories-mvc-api-codegen/16.7.0-SNAPSHOT/
  ls -lah ~/.m2/repository/com/nutanix/nutanix-core/ntnx-api/categories/categories-mvc-java-client-sdk/16.7.0-SNAPSHOT/

  EXIT_STATUS=0
  CMD="java -cp target/test-sdk-app-1.0-SNAPSHOT.jar com.nutanix.test.test-sdk.App"
  OUTPUT=$(eval "${CMD}")
  #echo "${OUTPUT}"
  if [ "${OUTPUT}" == "Hello World!" ]; then
    PASS "Maven Test Rig Tests Pass"
  else
    ERROR "Maven Test Rig Tests Failed"
    debug
    export EXIT_STATUS=1
  fi
  echo

  CMD="java -cp ~/.m2/repository/com/nutanix/test/test-sdk/${DEPLOYMENT_TAG}/test-sdk-${DEPLOYMENT_TAG}.jar com.nutanix.test.test-sdk.App"
  OUTPUT=$(eval "${CMD}")
  #echo "${OUTPUT}"
  if [ "${OUTPUT}" == "Hello World!" ]; then
    PASS "Maven Dependency Tests Pass"
    echo
    INFO "Staging Dependency for External Deployment"
    mkdir -p "${PROJECT_ROOT}"/mvn-external-release
    cp -rf ~/.m2/repository/com/nutanix/example/hellonutanixworld/"${DEPLOYMENT_TAG}"/ "${PROJECT_ROOT}"/mvn-external-release
  else
    ERROR "Maven Dependency Tests Failed"
    debug
    export EXIT_STATUS=1
  fi
  echo

  popd || exit 1
  popd || exit 1

}

function npm-internal-release-verify {
  rm -rf "${PROJECT_ROOT}"/npm-release-verify
  mkdir -p "${PROJECT_ROOT}"/npm-release-verify
  pushd "${PROJECT_ROOT}"/npm-release-verify || exit 1
  
  echo "${INTERNAL_NPM_PULL_CREDENTIALS}" > "${HOME}/.npmrc"

  if npm pack @nutanix-core/categories-javascript-client-sdk@"${DEPLOYMENT_TAG}" ; then
    PASS "Successfully Downloaded NPM Module"
    INFO @nutanix-core/categories-javascript-client-sdk@"${DEPLOYMENT_TAG}"
  else
    ERROR "Failed to Download NPM Module"
    INFO @nutanix-core/categories-javascript-client-sdk@"${DEPLOYMENT_TAG}"
    export EXIT_STATUS=1
  fi

  tar -zxf nutanix-core-categories-javascript-client-sdk-"${DEPLOYMENT_TAG}".tgz
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
    EXIT_STATUS=0
  fi


  popd || exit 1

}

if [[ "${BASH_SOURCE[0]}" == "${0}" ]]; then
  main "$@"
else
  echo "This file is designed to be executed, not sourced."
fi

exit "${EXIT_STATUS}"
