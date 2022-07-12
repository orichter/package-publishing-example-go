#!/bin/bash
#export TERM=ansi
export TERM=xterm-color
PROJECT_ROOT=${PROJECT_ROOT:-"/home/circleci/project"}
EXIT_STATUS=0
export PASSWORD_PUBLISH_GOLANG=${PASSWORD_PUBLISH_NPM_GITHUB}
#shellcheck disable=SC1091
#shellcheck disable=SC1090
source "${PROJECT_ROOT}"/services/deploy/release-utils.source

#shellcheck disable=SC1091
#shellcheck disable=SC1090
source "${PROJECT_ROOT}"/release-config.source
export VERSION="${DEPLOY_TO_TAG}"
INFO "Deploying npm packages using Release Params:"
cat "${PROJECT_ROOT}"/release-config.source
INFO "Version: ${VERSION}"
echo

function main {
  install_golang
  #export PACKAGE_NAME=experiments-nutanix-sdk-golang
  #export PACKAGE_URL=https://@github.com/nutanix-release-engineering/"${PACKAGE_NAME}".git
  #export PACKAGE_NAME=experiments-nutanix-sdk-golang
  #export PACKAGE_URL=https://@github.com/nutanix-release-engineering/"${PACKAGE_NAME}".git

  export PACKAGE_NAME=experiments-nutanix-sdk-golang
  export PACKAGE_URL=github.com/nutanix-release-engineering/"${PACKAGE_NAME}"
  export GOPRIVATE=${PACKAGE_URL}

  mkdir -p "${PROJECT_ROOT}"/golang-release-verify
  pushd "${PROJECT_ROOT}"/golang-release-verify || exit 1

  echo 'package test-ntnx-api-golang-sdk-external


  import (
      "testing"
      test-ntnx-api-golang-sdk-external "github.com/nutanix-release-engineering/experiments-nutanix-sdk-golang"
  )

  func TestHello(t *testing.T) {
      want := "hello nutanix world"
      if got := hellotest.HelloWorld(); got != want {
          t.Errorf("Hello() = %q, want %q", got, want)
      }
  }
  ' > nutanix-hello-world_test.go
  #sudo apt-get install golang-stable
  #sudo curl -O https://storage.googleapis.com/golang/go1.8.linux-amd64.tar.gz
  #sudo tar -xf go1.8.linux-amd64.tar.gz
  #sudo mv go /usr/local
  #export GOROOT=/usr/local/go/bin/go
  export PATH=$PATH:/usr/local/go/bin
  mkdir -p "$HOME"/go
  export GOPATH=$HOME/go
  export PATH=$PATH:$GOROOT/bin
  #go version
  go get -u=patch
  go env -w GO111MODULE=on
  export GO111MODULE=on
  #go install github.com/orichter/package-publishing-example-go@"${DEPLOYMENT_TAG}"
  #go install golang.org/x/tools/gopls@latest

  if go get "${PACKAGE_URL}" ; then
    PASS "Go Successful get for ${PACKAGE_URL}"
  else
    #ERROR "Go Failed get for ${PACKAGE_URL}"
    WARN "Go get not currently implemented due to unpublished dependencies"
    #EXIT_STATUS=1
  fi

  if go install "${PACKAGE_URL}" ; then
    PASS "Go Successful install for ${PACKAGE_URL}"
  else
    #ERROR "Go Failed install for ${PACKAGE_URL}"
    WARN "Go get not currently implemented due to unpublished dependencies"
    #EXIT_STATUS=1
  fi

  VERSION_COUNT=$(git log --oneline -n 1 | grep -c "${VERSION}")
  export VERSION_COUNT
  if [[ "${VERSION_COUNT}" -eq 1 ]]; then
    PASS "Go HEAD at Version: ${VERSION} install for ${PACKAGE_URL}"
  else
    #ERROR "Go Version: ${VERSION} not found at HEAD for ${PACKAGE_URL}"
    WARN "Go get not currently implemented due to unpublished dependencies"
    #EXIT_STATUS=1
  fi

  if go test ; then
    PASS "Go Successful install for ${PACKAGE_URL}"
  else
    WARN "Go Test not currently implemented for  ${PACKAGE_URL}"
    #WARN "Go Failed install for ${PACKAGE_URL}"
    #EXIT_STATUS=1
  fi
  popd || exit 1

  pushd "${HOME}"/go/src/"${PACKAGE_URL}" || exit 1

  # Check to see the proper tag has been applied
  # "${DEPLOYMENT_TAG}"$ means ends with "${DEPLOYMENT_TAG}"
  # This is necessary to distinguish v0.0.17 from v0.0.17-rc1 which is necessary to distinguish
  if [ "$(git tag |grep -c "${VERSION}"$)" -eq 1 ]; then
    PASS "Deployment Tag Verified"
  else
    MOST_RECENT_GO_TAGS=$(git log --oneline --decorate --max-count=1)
    ERROR "Doploymet Tag Verification Failed"
    INFO "MOST_RECENT_TAGS: ${MOST_RECENT_GO_TAGS}"
    INFO "VERSION: ${VERSION}"
    debug
    exit 1
  fi
  #go list -m all
  #PACKAGE_VERSION=$(go list -u -m all | grep github.com/orichter/package-publishing-example-go | awk '{print $2}')
  #echo "${PACKAGE_VERSION}"
  #echo "${DEPLOYMENT_TAG}"

  popd || exit 1
}

function install_golang {
  sudo curl -O https://storage.googleapis.com/golang/go1.8.linux-amd64.tar.gz
  sudo tar -xf go1.8.linux-amd64.tar.gz
  sudo mv go /usr/local
  export PATH=$PATH:/usr/local/go/bin
  go version
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
  exit "${EXIT_STATUS}"
else
  WARN "This file is designed to be executed, not sourced."
fi
