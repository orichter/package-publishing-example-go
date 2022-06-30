#!/bin/bash
#export TERM=ansi
export TERM=xterm-color
PROJECT_ROOT=${PROJECT_ROOT:-"/home/circleci/project"}
EXIT_STATUS=0
#export PASSWORD_PUBLISH_GOLANG=${PASSWORD_PUBLISH_NPM_GITHUB}
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
  sudo curl -O https://storage.googleapis.com/golang/go1.8.linux-amd64.tar.gz
  sudo tar -xf go1.8.linux-amd64.tar.gz
  sudo mv go /usr/local
  export PATH=$PATH:/usr/local/go/bin
  go version

  deploy-to-stage-internal
  #deploy-to-stage
  #deploy-to-github-prod
  #deploy-to-prod
}

function deploy-to-stage-internal {
  export PACKAGE_NAME=experiments-nutanix-sdk-golang
  export PACKAGE_URL=https://"${PASSWORD_PUBLISH_GOLANG}"@github.com/nutanix-release-engineering/"${PACKAGE_NAME}".git
  #export PACKAGE_URL=https://github.com/nutanix-release-engineering/"${PACKAGE_NAME}".git
  export GH_TOKEN=${PASSWORD_PUBLISH_GOLANG}

  # This directory should be created in the verify-private packages step.
  pushd "${PROJECT_ROOT}"/golang-release-verify || exit 1
  # Set publishing credentials
  git config --global user.email "api-packaging@nutanix.com"
  git config --global user.name "Nutanix Circle CI Release Agent"

  # Something like this should be done in verify-private-packages.sh so we assume it.
  #git clone https://github.com/nutanix-core/ntnx-api-golang-sdk-external
  pushd ./package || exit 1

  curl -fsSL https://cli.github.com/packages/githubcli-archive-keyring.gpg | sudo dd of=/usr/share/keyrings/githubcli-archive-keyring.gpg
  echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/githubcli-archive-keyring.gpg] https://cli.github.com/packages stable main" | sudo tee /etc/apt/sources.list.d/github-cli.list > /dev/null
  sudo apt update
  sudo apt install gh

  go version
  go mod tidy
  go test ./...

  INFO "Logging into Github using gh auth login"
  echo "${PASSWORD_PUBLISH_GOLANG}" | gh auth login --with-token
  #rm -rf .git

  #git init
  #git add .
  git commit -m "Deploying ntnx-api-golang-sdk-external Version: ${VERSION}"
  git tag "${VERSION}"
  git remote add deploy "${PACKAGE_URL}"
  #git push -u -f deploy main

  INFO "Pushing to deploy remote"
  if git push -f deploy HEAD:main ; then
    PASS "Golang Package ${PACKAGE_NAME} Successfully Deployed to Github Internal ${PACKAGE_URL}"
  else
    ERROR "Failed to Deploy Golang Package ${PACKAGE_NAME} to Github Internal ${PACKAGE_URL}"
    debug
    export EXIT_STATUS=1
    exit 1
  fi

  #git push -f deploy "${VERSION}"
  gh release create "${VERSION}" -F changelog.md --repo "${PACKAGE_URL}"

  # It appears releases can't be pushed to Nutanix internal repositories
  # the org owners don't delegate enough permissions for that purpose.
  #INFO "Uploading Release"
  #if gh release upload "${VERSION}" . ; then
  #  PASS "Golang Package ${PACKAGE_NAME} Successfully Released on Github Internal ${PACKAGE_URL}"
  #else
  #  ERROR "Failed to Release Golang Package ${PACKAGE_NAME} on Github Internal ${PACKAGE_URL}"
  #  debug
  #  export EXIT_STATUS=1
  #  exit 1
  #fi

  #git push https://"${PASSWORD_PUBLISH_GOLANG}"@github.com/orichter/package-publishing-example-go.git "${DEPLOYMENT_TAG}"

  #git push https://"${PASSWORD_PUBLISH_GOLANG}"@github.com/orichter/package-publishing-example-go.git "${DEPLOYMENT_TAG}"
  #git push origin "${DEPLOYMENT_TAG}"
  debug
  echo

  popd || exit 1
  popd || exit 1

}

function debug {
  #env |grep -v PASSWORD
  echo
  pwd
  ls -lah
  git status
  git log --oneline --decorate -n 15
  git remote -v
  echo
}

if [[ "${BASH_SOURCE[0]}" == "${0}" ]]; then
  main "$@"
  exit "${EXIT_STATUS}"
else
  WARN "This file is designed to be executed, not sourced."
fi
