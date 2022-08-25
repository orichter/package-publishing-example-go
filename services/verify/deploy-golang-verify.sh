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
if [[ "${SUPPRESS_GOLANG}" == "true" ]]; then
  WARN "Suppressing golang debloyments. See manage-package-deployments.sh"
  exit
fi

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

  mkdir -p "${PROJECT_ROOT}"/golang-release-verify
  pushd "${PROJECT_ROOT}"/golang-release-verify || exit 1

  # HACK: This call currently only verifies the internal repository.
  # It should also verify the external repositories.
  validate "${PROJECT_ROOT}"/services/deploy/manifest.json
  return
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

  if go get "${PACKAGE_URL}" ; then
    PASS "Go Successful get for ${PACKAGE_URL}"
  else
    #ERROR "Go Failed get for ${PACKAGE_URL}"
    WARN "Go get not currently implemented due to unpublished dependencies"
    #EXIT_STATUS=1
  fi

  # Should be populated in verify-private-packages.sh.
  cp "${PROJECT_ROOT}"/go.mod .
  if go install "${PACKAGE_URL}" ; then
    PASS "Go Successful install for ${PACKAGE_URL}"
  else
    ERROR "Go Failed install for ${PACKAGE_URL}"
    #WARN "Go get not currently implemented due to unpublished dependencies"
    EXIT_STATUS=1
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

function validate {
  #!/bin/bash
  set -e

  # This script installs golang if not present on the system and then it installs all the golang sdks provided in manifest.json.
  # It does a go get for all the SDKs and see if it successful or not.
  # PARAM :-
  # [ REQUIRED ] 1ST ARGUMENT := PATH TO MANIFEST.JSON
  # [ OPTIONAL ] 2ND ARGUMENT := URL OF THE GITHUB REPO WHERE SDKS ARE HOSTED [ DEFAULT VALUE = github.com/nutanix-core/ntnx-api-golang-sdk-external ]
  # NOTE: This script should not be run with sh command for .e.g `sh testSdk.sh` is not supported.

  MANIFEST=$1
  GO_SRC_URL=$2
  RED="\e[31m"
  GREEN="\e[32m"
  YELLOW="\e[33m"
  ENDCOLOR="\e[0m"

  # initArch discovers the architecture for this system.
  initArch() {
    ARCH=$(uname -m)
    case $ARCH in
      armv5*) ARCH="armv5";;
      armv6*) ARCH="armv6";;
      armv7*) ARCH="arm";;
      aarch64) ARCH="arm64";;
      x86) ARCH="386";;
      x86_64) ARCH="amd64";;
      i686) ARCH="386";;
      i386) ARCH="386";;
    esac
  }

  # initOS discovers the operating system for this system.
  initOS() {
    #shellcheck disable=2005
    #shellcheck disable=2006
    #shellcheck disable=2046
    OS=$(echo `uname`|tr '[:upper:]' '[:lower:]')
    case "$OS" in
      # Minimalist GNU for Windows
      mingw*|cygwin*) OS='windows';;
    esac
  }

  # installed the desired version of go defined in $GOVERSION variable and install
  # it at the location defined by $GOINSTALLDIR.
  installGo() {
    initArch
    initOS
    mkdir -p "${GOINSTALLDIR}"
    wget https://dl.google.com/go/go"${GOVERSION}"."${OS}"-"${ARCH}".tar.gz -O "${GOINSTALLDIR}"/"${GOTAR}"
    tar -xzf "${GOINSTALLDIR}/${GOTAR}" -C "${GOINSTALLDIR}"
    # removing the tar as it is no longer needed
    rm -rf "${GOINSTALLDIR:?}/${GOTAR:?}"
  }

  printGreen() {
    #shellcheck disable=2046
    #shellcheck disable=2059
    echo -e  $(printf "${GREEN}$1${ENDCOLOR}")
  }

  printRed() {
    #shellcheck disable=2046
    #shellcheck disable=2059
    echo -e  $(printf "${RED}$1${ENDCOLOR}")
  }

  printYellow() {
    #shellcheck disable=2046
    #shellcheck disable=2059
    echo -e  $(printf "${YELLOW}$1${ENDCOLOR}")
  }

  if [ -z "$1" ]
    then
      printRed "First argument should be the path to manifest.json. Please correct it and retry."
      exit 1
  fi

  if [ -z "$2" ]
    then
      printYellow "Source code url is missing, using github.com/nutanix-core/ntnx-api-golang-sdk-external as default."
      GO_SRC_URL="github.com/nutanix-core/ntnx-api-golang-sdk-external"
  fi

  # if ! [ -x "$(command -v go)" ]; then
  # HACK: We should really replace the earlier go install with this one.
  # For now, we simply always overwrite.
  if true; then
    GOVERSION=1.16 # the default version used
    GOTAR=go.tar.gz # go tar file will be renamed to this
    PD=$(pwd) # capturing PD just to ensure working with absolute paths
    GOINSTALLDIR="$PD"/.install-go # by default it gets installed inside .install-go
    printYellow "Installing go of version ${GOVERSION}"
    installGo
    export GOROOT=$GOINSTALLDIR/go
    export PATH=$GOINSTALLDIR/go/bin:$PATH
  else
    printYellow "Skipping go installation as it has already been installed."
  fi

  export GOPRIVATE="github.com/nutanix-core,github.com/nutanix"


  printYellow "Printing GO Version"
  go version
  echo

  # creating a go.mod file, just to make sure this script works with the latest versions of go.
  go mod init testingv4GoSDKs
  #shellcheck disable=2154
  echo "$x" | jq -r '.[]|"\(.Namespace) \(.LanguageAndVersion) \(.APIVersion)"' "${MANIFEST}" | while read -r Namespace LanguageAndVersion APIVersion; do
      printYellow "|.................................................................|"
      printYellow "    Processing Namespace ${Namespace} with APIVersion ${APIVersion}"
      printYellow "|.................................................................|"

      #shellcheck disable=2259
      #shellcheck disable=2162
      echo "$x" | jq -r '.[]|"\(.Language) \(.Version)"' <<<  "${LanguageAndVersion}" | while read Language Version ; do
          if [[ ${Language} == "Golang" ]] ; then
            echo "Testing Golang SDK for namespace ${Namespace} with version ${Version} "
            #shellcheck disable=SC2086
            go get ${GO_SRC_URL}/${Namespace}-go-client/v4/...@${Version}
            printGreen "---------------------Test Succesfull------------------------------"
          fi
      done
      printYellow "|.............................xxx.................................|"
      echo
      echo
      echo
  done


  printYellow "|.................................................................|"
  printYellow "                    Printing go.mod contents"
  printYellow "|.................................................................|"
  echo
  echo
  cat go.mod
  echo
  echo
  printYellow "|.............................xxx.................................|"
  echo
  echo
  echo

  printGreen "All tests passed successfully"
  printYellow "Cleaning up files"
  rm -rf go.mod
  rm -rf go.sum
  rm -rf "${GOINSTALLDIR}"

}


if [[ "${BASH_SOURCE[0]}" == "${0}" ]]; then
  main "$@"
  exit "${EXIT_STATUS}"
else
  WARN "This file is designed to be executed, not sourced."
fi
