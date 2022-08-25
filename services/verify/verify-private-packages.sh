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

  clean-all
  # DEPLOY_TO_TAG="0.1.4-0-1-rc1"
  # DEPLOY_FROM_TAG: "4.0.1-alpha.1"

  NAMESPACES="vmm prism clustermgmt aiops iam storage"
  # Iterate the string variable using for loop
  for NAMESPACE in ${NAMESPACES}; do
    echo "Verifying Namespace: ${NAMESPACE}"
    # HACK: Crude deployment manifest which needs to be reworked.
    if [ "${NAMESPACE}" = "storage" ] ; then
      VERSION="4.0.1-alpha.2"
      #MVN_VERSION="4.0.1-alpha-1"
      #MVN_VERSION="${DEPLOY_FROM_TAG}"
    else
      VERSION="${DEPLOY_FROM_TAG}"
      #MVN_VERSION="${DEPLOY_FROM_TAG}"
    fi

    if [[ "${SUPPRESS_PIP}" == "true" ]]; then
      WARN "Suppressing pip debloyments. See manage-package-deployments.sh"
    else
      pip-internal-release-verify "${NAMESPACE}" "${VERSION}"
    fi

    if [[ "${SUPPRESS_NPM}" == "true" ]]; then
      WARN "Suppressing npm deployments. See manage-package-deployments.sh"
    else
      npm-internal-release-verify "${NAMESPACE}" "${VERSION}"
    fi

    if [[ "${SUPPRESS_GOLANG}" == "true" ]]; then
      WARN "Suppressing golang deployments. See manage-package-deployments.sh"
    else
      golang-internal-release-verify "${NAMESPACE}" "${VERSION}"
    fi

    # HACK: MVN Deployment Version is currently pinned.
    #VERSION="4.0.0-alpha-1"
    #MVN_VERSION="${MVN_VERSION/4.0.1/4.0.0}"
    if [[ "${SUPPRESS_MVN}" == "true" ]]; then
      WARN "Suppressing mvn deployments. See manage-package-deployments.sh"
    else
      MVN_VERSION=${VERSION/alpha./alpha-}
      WARN "MVN Version rewritten to: ${MVN_VERSION}"
      mvn-internal-release-verify "${NAMESPACE}" "${MVN_VERSION}"
    fi

  done
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
  NAMESPACE=$1
  VERSION=$2
  #export VERSION="${DEPLOY_FROM_TAG}"
  # HACK: This should be named golang-client for consistency
  export PACKAGE_NAME="${NAMESPACE}"-go-client

  sudo curl -O https://storage.googleapis.com/golang/go1.8.linux-amd64.tar.gz
  sudo tar -xf go1.8.linux-amd64.tar.gz
  sudo mv go /usr/local
  export PATH=$PATH:/usr/local/go/bin
  go version

  pushd "${PROJECT_ROOT}"/verify/golang-release-verify || exit 1
  
  export REPO_NAME=ntnx-api-golang-sdk-external
  export PACKAGE_URL=https://github.com/nutanix-core/"${REPO_NAME}".git

  # HACK: We only need to clone on the first call vmm, but this is a klugy way to do so.
  if [ "${NAMESPACE}" = "vmm" ] ; then
    git clone "${PACKAGE_URL}"
  fi

  pushd "${REPO_NAME}" || exit 1
  git fetch --all --tags

  # HACK: This is not a good tagging convention and should be refactored.
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
  rm -rf ./package
  cp -rf ./"${REPO_NAME}" ./package
  ls -lah package
  popd || exit 1

}

function mvn-internal-release-verify {
  NAMESPACE=$1
  VERSION=$2
  export PACKAGE_NAME="${NAMESPACE}"-java-client
  # HACK: Other package versions should be made consistent with MVN_VERSION
  # If this is done, this hack can be removed.
  export MVN_VERSION=${VERSION/alpha./alpha-}

  pushd "${PROJECT_ROOT}"/verify/maven-release-verify || exit 1

  mvn archetype:generate -q -DgroupId=com.nutanix.test.sdk -DartifactId=test-"${NAMESPACE}"-sdk-app -DarchetypeArtifactId=maven-archetype-quickstart -DarchetypeVersion=1.4 -DinteractiveMode=false

  pushd test-"${NAMESPACE}"-sdk-app || exit 1
  mv pom.xml pom.xml.old
  cp "${PROJECT_ROOT}"/services/deploy/mvn/pom.template ./pom.xml
  cp "${PROJECT_ROOT}"/services/deploy/mvn/settings.xml .
  sed -i "s|.{env.DEPLOYMENT_TAG}|${MVN_VERSION}|g" pom.xml
  sed -i "s|-java-client|${PACKAGE_NAME}|g" pom.xml

  if mvn package -q --settings settings.xml -DdownloadSources=true -DdownloadJavadocs=true ; then
    PASS "Successfully Downloaded Maven Package: ${PACKAGE_NAME} using pom:"
    cat pom.xml
    mvn dependency:sources dependency:resolve -Dclassifier=javadoc -q -s settings.xml
    mvn dependency:sources -Dsilent=true -q -s settings.xml
    cp -rf "${HOME}"/.m2/repository/com/nutanix/api/"${PACKAGE_NAME}"/"${MVN_VERSION}"/ "${PROJECT_ROOT}"/verify/maven-release-verify/package-"${NAMESPACE}"
    INFO "MVN Package Contents:"
    ls -lah "${PROJECT_ROOT}"/verify/maven-release-verify/package-"${NAMESPACE}"
    pushd "${PROJECT_ROOT}"/verify/maven-release-verify/package-"${NAMESPACE}" || exit 1
    #mvn dependency:sources dependency:resolve -Dclassifier=javadoc -q -s settings.xml
    #mvn dependency:sources -Dsilent=true -q -s settings.xml
    #ls -lah "${PROJECT_ROOT}"/verify/maven-release-verify/package
    #cat vmm-java-client-4.0.1-alpha-1-sources.jar.lastUpdated
    popd || exit 1
    echo
  else
    ERROR "Failed to Download Maven Package ${PACKAGE_NAME}"
    export EXIT_STATUS=1
    debug-maven
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
  NAMESPACE=$1
  VERSION=$2
  export PACKAGE_NAME="${NAMESPACE}"-js-client
  export SCOPE=nutanix-core

  pushd "${PROJECT_ROOT}"/verify/npm-release-verify || exit 1
  
  echo "${INTERNAL_NPM_PULL_CREDENTIALS}" > "${HOME}/.npmrc"

  if npm pack @nutanix-core/"${PACKAGE_NAME}"@"${VERSION}" ; then
    PASS "Successfully Downloaded NPM Package: ${PACKAGE_NAME}"
    INFO @"${SCOPE}"/"${PACKAGE_NAME}"@"${VERSION}"
  else
    ERROR "Failed to Download NPM Package"
    INFO @"${SCOPE}"/"${PACKAGE_NAME}"@"${VERSION}"
    export EXIT_STATUS=1
  fi

  #ls -lah
  tar -zxf "${SCOPE}"-"${PACKAGE_NAME}"-"${VERSION}".tgz
  # This should be untarred into package which will be used in deploy-npm.sh for actual deployments

  mv -f package package-"${NAMESPACE}"
  cp -rf package-"${NAMESPACE}" audit-"${NAMESPACE}"
  pushd package-"${NAMESPACE}" || exit 1
  # This is necessary so the blackduck scan doesn't complain.
  npm i > /dev/null
  popd || exit 1

  pushd audit-"${NAMESPACE}" || exit 1
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
  NAMESPACE=$1
  VERSION=$2
  PACKAGE_DIR="${NAMESPACE}"
  export PACKAGE_NAME="ntnx_${NAMESPACE}_py_client"
  export PACKAGE_NAME="${PACKAGE_NAME/_/-}"

  pushd "${PROJECT_ROOT}"/verify/pip-release-verify || exit 1
  
  sudo apt install python3-pip > /dev/null
  /home/circleci/.pyenv/versions/3.8.5/bin/python3.8 -m pip install --upgrade pip > /dev/null

  #PACKAGE_PATH="git+ssh://git@github.com/nutanix-core/ntnx-api-python-sdk-external.git@${TMP_VERSION}#subdirectory=${PACKAGE_DIR}&egg=${PACKAGE_NAME}"
  PACKAGE_PATH="git+ssh://git@github.com/nutanix-core/ntnx-api-python-sdk-external.git#subdirectory=${PACKAGE_DIR}&egg=${PACKAGE_NAME}"

  if pip3 install -e "${PACKAGE_PATH}" ; then
    PASS "Successfully Downloaded Python Package: ${PACKAGE_NAME}"
    INFO "${PACKAGE_PATH}"
    #ls -lah
    #cp -rf package package-"${NAMESPACE}"
  else
    ERROR "Failed to Download Python Package: ${PACKAGE_NAME}"
    INFO "${PACKAGE_PATH}"
    debug-pip
    export EXIT_STATUS=1
  fi

  # HACK: Python Package Directory is not named consistently with the directory
  cp -rf src/"${PACKAGE_NAME//_/-}"/"${PACKAGE_DIR}" ./package-"${NAMESPACE}"

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

function clean-all {
  rm -rf "${PROJECT_ROOT}"/verify/pip-release-verify
  mkdir -p "${PROJECT_ROOT}"/verify/pip-release-verify
  rm -rf "${PROJECT_ROOT}"/verify/npm-release-verify
  mkdir -p "${PROJECT_ROOT}"/verify/npm-release-verify
  rm -rf "${PROJECT_ROOT}"/verify/golang-release-verify
  mkdir -p "${PROJECT_ROOT}"/verify/golang-release-verify
  rm -rf "${PROJECT_ROOT}"/verify/maven-release-verify
  mkdir -p "${PROJECT_ROOT}"/verify/maven-release-verify

}

function debug-pip {
  WARN "Maven PIP"
  debug
  echo
  ls -lah src/
  ls -lah src/"${PACKAGE_NAME//_/-}"/"${PACKAGE_DIR}"
}

function debug-maven {
  WARN "Maven Debug"
  debug
  ls -lah ~/.m2/repository/com/nutanix/api
  ls -lah ~/.m2/repository/com/nutanix/api/"${PACKAGE_NAME}"
  ls -lah ~/.m2/repository/com/nutanix/api/"${PACKAGE_NAME}"/"${VERSION}"/
  ls -lah "${PROJECT_ROOT}"/verify/maven-release-verify
  ls -lah "${PROJECT_ROOT}"/verify/maven-release-verify
  ls -lah "${PROJECT_ROOT}"/verify/maven-release-verify/package
  WARN "Failed pom.xml:"
  cat pom.xml
  WARN "Failed settings.xml"
  cat settings.xml

}

if [[ "${BASH_SOURCE[0]}" == "${0}" ]]; then
  main "$@"
else
  echo "This file is designed to be executed, not sourced."
fi

exit "${EXIT_STATUS}"
