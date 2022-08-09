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
  # HACK: golang currently can't be verified, so it is removed.
  pip-internal-release-verify
  npm-internal-release-verify
  mvn-internal-release-verify
  golang-internal-release-verify
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
  export VERSION="${DEPLOY_FROM_TAG}"
  #HACK: This should be named golang-client for consistency
  export PACKAGE_NAME="${NAMESPACE}"-go-client

  sudo curl -O https://storage.googleapis.com/golang/go1.8.linux-amd64.tar.gz
  sudo tar -xf go1.8.linux-amd64.tar.gz
  sudo mv go /usr/local
  export PATH=$PATH:/usr/local/go/bin
  go version

  rm -rf "${PROJECT_ROOT}"/verify/golang-release-verify
  mkdir -p "${PROJECT_ROOT}"/verify/golang-release-verify
  pushd "${PROJECT_ROOT}"/verify/golang-release-verify || exit 1
  
  export REPO_NAME=ntnx-api-golang-sdk-external
  export PACKAGE_URL=https://github.com/nutanix-core/"${REPO_NAME}".git
  git clone "${PACKAGE_URL}"

  pushd "${REPO_NAME}" || exit 1
  git fetch --all --tags

  # This is not a good tagging convention and should be refactored.
  export PACKAGE_TAG="${PACKAGE_NAME}"/v"${VERSION}"
  if git checkout "${PACKAGE_TAG}" ; then
    PASS "Tag ${VERSION} of Package: ${PACKAGE_NAME} Successfully checked out from ${PACKAGE_URL}"
    git log --oneline --decorate -n 15
  else
    ERROR "Tag ${VERSION} of Package: ${PACKAGE_NAME} checked out failed from ${PACKAGE_URL}."
    git log --oneline --decorate -n 15
    EXIT_STATUS=1
    debug
  fi

  pushd "${PACKAGE_NAME}" || exit 1
  # For later use in the verification step.
  #cp go.mod "${PROJECT_ROOT}"
  #if go install "${PACKAGE_URL}"/"${PACKAGE_NAME}" ; then
  #  PASS "Go Successful install of ${PACKAGE_NAME} from ${PACKAGE_URL}"
  #else
  #  ERROR "Go Failed install of ${PACKAGE_NAME} from ${PACKAGE_URL}"
    #WARN "Go get not currently implemented due to unpublished dependencies"
  #  EXIT_STATUS=1
  #fi
  popd || exit 1

  popd || exit 1
  # Rename package for standard deployment, and so we have a copy of the original
  # for debugging purposes.
  cp -rf "${REPO_NAME}"/"${PACKAGE_NAME}" package
  ls -lah package
  popd || exit 1

}

function mvn-internal-release-verify {
  export VERSION="${DEPLOY_FROM_TAG}"
  #HACK Version is currently pinned.
  VERSION="4.0.0-alpha-1"
  export PACKAGE_NAME="${NAMESPACE}"-java-client
  #HACK: Other package versions should be made consistent with MVN_VERSION
  # If this is done, this hack can be removed.
  export MVN_VERSION=${VERSION/alpha./alpha-}

  rm -rf "${PROJECT_ROOT}"/verify/maven-release-verify
  mkdir -p "${PROJECT_ROOT}"/verify/maven-release-verify

  pushd "${PROJECT_ROOT}"/verify/maven-release-verify || exit 1

  mvn archetype:generate -q -DgroupId=com.nutanix.test.sdk -DartifactId=test-sdk-app -DarchetypeArtifactId=maven-archetype-quickstart -DarchetypeVersion=1.4 -DinteractiveMode=false

  pushd test-sdk-app || exit 1
  mv pom.xml pom.xml.old
  cp "${PROJECT_ROOT}"/services/deploy/mvn/pom.template ./pom.xml
  cp "${PROJECT_ROOT}"/services/deploy/mvn/settings.xml .
  sed -i "s|.{env.DEPLOYMENT_TAG}|${MVN_VERSION}|g" pom.xml
  sed -i "s|-java-client|${PACKAGE_NAME}|g" pom.xml
  #cat pom.xml

  if mvn package -q --settings settings.xml -DdownloadSources=true -DdownloadJavadocs=true ; then
    PASS "Successfully Downloaded Maven Package using pom:"
    cat pom.xml
    mvn dependency:sources dependency:resolve -Dclassifier=javadoc -q -s settings.xml
    mvn dependency:sources -Dsilent=true -q -s settings.xml
    cp -rf ~/.m2/repository/com/nutanix/api/"${PACKAGE_NAME}"/"${MVN_VERSION}"/ "${PROJECT_ROOT}"/verify/maven-release-verify/package
    INFO "MVN Package Contents:"
    ls -lah "${PROJECT_ROOT}"/verify/maven-release-verify/package
    pushd "${PROJECT_ROOT}"/verify/maven-release-verify/package || exit 1
    #mvn dependency:sources dependency:resolve -Dclassifier=javadoc -q -s settings.xml
    #mvn dependency:sources -Dsilent=true -q -s settings.xml
    #ls -lah "${PROJECT_ROOT}"/verify/maven-release-verify/package
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
  export VERSION="${DEPLOY_FROM_TAG}"
  export PACKAGE_NAME="${NAMESPACE}"-js-client
  export SCOPE=nutanix-core

  rm -rf "${PROJECT_ROOT}"/verify/npm-release-verify
  mkdir -p "${PROJECT_ROOT}"/verify/npm-release-verify
  pushd "${PROJECT_ROOT}"/verify/npm-release-verify || exit 1
  
  echo "${INTERNAL_NPM_PULL_CREDENTIALS}" > "${HOME}/.npmrc"

  if npm pack @nutanix-core/"${PACKAGE_NAME}"@"${VERSION}" ; then
    PASS "Successfully Downloaded NPM Package"
    INFO @"${SCOPE}"/"${PACKAGE_NAME}"@"${VERSION}"
  else
    ERROR "Failed to Download NPM Package"
    INFO @"${SCOPE}"/"${PACKAGE_NAME}"@"${VERSION}"
    export EXIT_STATUS=1
  fi

  ls -lah
  tar -zxf "${SCOPE}"-"${PACKAGE_NAME}"-"${VERSION}".tgz
  # This should be untarred into package which will be used in deploy-npm.sh for actual deployments

  cp -rf package audit
  pushd package || exit 1
  # This is necessary so the blackduck scan doesn't complain.
  npm i > /dev/null
  popd || exit 1

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
    #EXIT_STATUS=1
  fi

  popd || exit 1

}

function pip-internal-release-verify {
  export VERSION="${DEPLOY_FROM_TAG}"
  PACKAGE_DIR="${NAMESPACE}"
  export PACKAGE_NAME="ntnx_${NAMESPACE}_py_client"
  export PACKAGE_NAME="${PACKAGE_NAME/_/-}"

  rm -rf "${PROJECT_ROOT}"/verify/pip-release-verify
  mkdir -p "${PROJECT_ROOT}"/verify/pip-release-verify
  pushd "${PROJECT_ROOT}"/verify/pip-release-verify || exit 1
  
  sudo apt install python3-pip
  /home/circleci/.pyenv/versions/3.8.5/bin/python3.8 -m pip install --upgrade pip

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

  #cp -rf src/"${PACKAGE_NAME}"/"${PACKAGE_DIR}" ./package
  ls -lah src/
  echo
  ls -lah src/"${PACKAGE_NAME}"
  echo
  ls -lah src/"${PACKAGE_DIR}"
  echo
  ls -lah src/"${PACKAGE_NAME}"/"${PACKAGE_DIR}"
  echo
  ls -lah src/"${PACKAGE_DIR}"/"${PACKAGE_NAME}"

  popd || exit 1

}

#shellcheck disable=SC1091
#shellcheck disable=SC1090
source "${PROJECT_ROOT}"/verify/services/deploy/release-utils.source
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
  ls -lah ~/.m2/repository/com/nutanix/api/"${PACKAGE_NAME}"
  ls -lah ~/.m2/repository/com/nutanix/api/"${PACKAGE_NAME}"/"${VERSION}"/
  ls -lah "${PROJECT_ROOT}"/verify/maven-release-verify
  ls -lah "${PROJECT_ROOT}"/verify/maven-release-verify
  ls -lah "${PROJECT_ROOT}"/verify/maven-release-verify/package
}

if [[ "${BASH_SOURCE[0]}" == "${0}" ]]; then
  main "$@"
else
  echo "This file is designed to be executed, not sourced."
fi

exit "${EXIT_STATUS}"
