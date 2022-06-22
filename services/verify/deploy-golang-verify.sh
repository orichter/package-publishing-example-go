#!/bin/bash
#export TERM=ansi
#export TERM=xterm-color
PROJECT_ROOT=${PROJECT_ROOT:-"/home/circleci/project"}
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
source ./release-config.source
echo "Verifying golang package deployment using Release Params:"
cat ./release-config.source

mkdir -p ./golang-release-verify
pushd ./golang-release-verify || exit 1

echo 'package test_hello_world

import (
    "testing"
    hellotest "github.com/orichter/package-publishing-example-go"
)

func TestHello(t *testing.T) {
    want := "hello nutanix world"
    if got := hellotest.HelloWorld(); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}
' > nutanix-hello-world_test.go
#sudo apt-get install golang-stable
sudo curl -O https://storage.googleapis.com/golang/go1.8.linux-amd64.tar.gz
sudo tar -xf go1.8.linux-amd64.tar.gz
sudo mv go /usr/local
export PATH=$PATH:/usr/local/go/bin
#mkdir -p "$HOME"/go
#export GOROOT=$HOME/go
#export PATH=$PATH:$GOROOT/bin
go version
#go get -u=patch
#go env -w GO111MODULE=on
#export GO111MODULE=on
#go install github.com/orichter/package-publishing-example-go@"${DEPLOYMENT_TAG}"
#go install golang.org/x/tools/gopls@latest
go get github.com/orichter/package-publishing-example-go
go install github.com/orichter/package-publishing-example-go
if go test; then
  PASS "Go Tests Pass"
else
  ERROR "Go Tests Failed"
  exit 1
fi
popd || exit 1


pushd "${HOME}"/go/src/github.com/orichter/package-publishing-example-go || exit 1

# Check to see the proper tag has been applied
# "${DEPLOYMENT_TAG}"$ means ends with "${DEPLOYMENT_TAG}"
# This is necessary to distinguish v0.0.17 from v0.0.17-rc1 which is necessary to distinguish
if [ "$(git tag |grep -c "${DEPLOYMENT_TAG}"$)" -eq 1 ]; then
  PASS "Deployment Tag Verified"
else
  MOST_RECENT_GO_TAGS=$(git log --oneline --decorate --max-count=1)
  ERROR "Doploymet Tag Verification Failed"
  INFO "MOST_RECENT_TAGS: ${MOST_RECENT_GO_TAGS}"
  INFO "DEPLOYMENT_TAG: ${DEPLOYMENT_TAG}"
  debug
  exit 1
fi
#go list -m all
#PACKAGE_VERSION=$(go list -u -m all | grep github.com/orichter/package-publishing-example-go | awk '{print $2}')
#echo "${PACKAGE_VERSION}"
#echo "${DEPLOYMENT_TAG}"

popd || exit 1
