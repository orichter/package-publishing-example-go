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
if [[ "${SUPPRESS_GOLANG}" == "true" ]]; then
  WARN "Suppressing golang debloyments. See manage-package-deployments.sh"
  exit
fi

export VERSION="${DEPLOY_TO_TAG}"
INFO "Deploying Golang packages using Release Params:"
cat "${PROJECT_ROOT}"/release-config.source
INFO "Version: ${VERSION}"
echo

function main {
#  sudo curl -O https://storage.googleapis.com/golang/go1.8.linux-amd64.tar.gz
#  sudo tar -xf go1.8.linux-amd64.tar.gz
#  sudo mv go /usr/local
  export PATH=$PATH:/usr/local/go/bin
#  go version

  clean-all

  NAMESPACES="vmm prism clustermgmt aiops iam storage"
  # Iterate the string variable using for loop
  for NAMESPACE in ${NAMESPACES}; do
    echo "Golang Deploying Namespace: ${NAMESPACE}"
    # HACK: Crude deployment manifest which needs to be reworked.
    if [ "${NAMESPACE}" = "storage" ] ; then
      DEPLOY_FROM_TAG="4.0.1-alpha.2"
      VERSION=${DEPLOY_TO_TAG}-${DEPLOY_FROM_TAG/./-}
    else
      DEPLOY_FROM_TAG="4.0.1-alpha.1"
      VERSION=${DEPLOY_TO_TAG}-${DEPLOY_FROM_TAG/./-}
    fi

    deploy-to-stage-internal "${NAMESPACE}" "${VERSION}" "${DEPLOY_FROM_TAG}"
    #deploy-to-stage "${NAMESPACE}" "${VERSION}" "${DEPLOY_FROM_TAG}"
    deploy-to-github-prod "${NAMESPACE}" "${VERSION}" "${DEPLOY_FROM_TAG}"

  done

  PASS "Successful Deployments can be found at:"
  cat "${PROJECT_ROOT}"/verify/golang-release-verify/successful-deployments.txt
  if test -f "${PROJECT_ROOT}/verify/golang-release-verify/failed-deployments.txt"; then
    ERROR "Failed Deployments to:"
    cat "${PROJECT_ROOT}"/verify/golang-release-verify/failed-deployments.txt
  fi

}

function reset-git-repo {
  REPO=$1
  cp -rf ./package "${REPO}"
  pushd "${REPO}" || exit 1
  rm -rf .git

  git config --global user.email "owen@nutanix.com"
  git config --global user.name "Owen Richter"

  git init
  # HACK: Prod Replacement
  sed -i 's#nutanix-core/ntnx-api-golang-sdk-external#orichter/package-publishing-example-go#g' README.md
  git add README.md
  git add LICENSE.md
  git commit -a -m "Initial Top Level Commit with LICENSE.md and README.md"

  popd || exit 1
}

function clean-all {
  pushd "${PROJECT_ROOT}"/verify/golang-release-verify || exit 1
  rm -rf "${PROJECT_ROOT}"/verify/golang-release-verify/package-stage-internal/
  # This seemingly unnecessary assignment prevents prior settings of this flag from affecting this deployment.
  RESET_REPO=false
  #RESET_REPO=true
  if [[ "${RESET_REPO}" == "true" ]]; then
    reset-git-repo "${PROJECT_ROOT}"/verify/golang-release-verify/package-stage-internal
  else
    git clone git@github.com:nutanix-release-engineering/experiments-nutanix-sdk-golang.git
    mv ./experiments-nutanix-sdk-golang "${PROJECT_ROOT}"/verify/golang-release-verify/package-stage-internal/
  fi

  pushd "${PROJECT_ROOT}"/verify/golang-release-verify || exit 1
  rm -rf "${PROJECT_ROOT}"/verify/golang-release-verify/package-github-prod/
  # This seemingly unnecessary assignment prevents prior settings of this flag from affecting this deployment.
  RESET_REPO=false
  #RESET_REPO=true
  if [[ "${RESET_REPO}" == "true" ]]; then
    reset-git-repo "${PROJECT_ROOT}"/verify/golang-release-verify/package-github-prod
  else
    git clone git@github.com:orichter/package-publishing-example-go.git
    mv ./experiments-nutanix-sdk-golang "${PROJECT_ROOT}"/verify/golang-release-verify/package-github-prod/
  fi

  popd || exit 1

}

function deploy-to-stage-internal {
  NAMESPACE=$1
  VERSION=$2
  DEPLOY_FROM_TAG=$3
  export PACKAGE_NAME="${NAMESPACE}-go-client"

  export REPO_NAME=experiments-nutanix-sdk-golang
  export PACKAGE_URL=https://github.com/nutanix-release-engineering/"${REPO_NAME}"/releases/tag/"${VERSION}"
  export PACKAGE_AUTH_URL=https://"${PASSWORD_PUBLISH_GOLANG}"@github.com/nutanix-release-engineering/"${REPO_NAME}".git
  #export PACKAGE_URL=https://github.com/nutanix-release-engineering/"${REPO_NAME}".git
  export GH_TOKEN=${PASSWORD_PUBLISH_GOLANG}

  # This directory should be created in the verify-private packages step.
  pushd "${PROJECT_ROOT}"/verify/golang-release-verify || exit 1
  # Set publishing credentials
  git config --global user.email "api-packaging@nutanix.com"
  git config --global user.name "Nutanix Circle CI Release Agent"

  # Install Github Release Tools
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
  #git grep -l 'nutanix-core/ntnx-api-golang-sdk-external' | xargs sed -i '' -e 's#nutanix-core/ntnx-api-golang-sdk-external#nutanix/ntnx-api-go-clients#g'
  git checkout "${PACKAGE_NAME}/v${DEPLOY_FROM_TAG}"

  # Something like this should be done in verify-private-packages.sh so we assume it.
  #git clone https://github.com/nutanix-core/ntnx-api-golang-sdk-external
  pushd ./package || exit 1

  pushd "${PACKAGE_NAME}" || exit 1
  INFO "Rewriting repository url"
  # HACK: Modify Sujeets code for Linux compatibility.
  git grep -l 'nutanix-core/ntnx-api-golang-sdk-external' | xargs sed -i -e 's#nutanix-core/ntnx-api-golang-sdk-external#nutanix-release-engineering/experiments-nutanix-sdk-golang#g'
  popd || exit 1

  cp -rf ./"${PACKAGE_NAME}" "${PROJECT_ROOT}"/verify/golang-release-verify/package-stage-internal/
  git checkout main

  pushd "${PROJECT_ROOT}"/verify/golang-release-verify/package-stage-internal || exit 1
  # The following line can be used to test multiple subsequent versioned commits.
  echo "Testing update to verstion ${VERSION}" >> ./"${PACKAGE_NAME}"/test-updates.txt; git add ./test-updates.txt; # HACK: comment this out.

  INFO "Deploying ntnx-api-golang-sdk-external Namespace: ${NAMESPACE} Version: ${VERSION}"
  git add ./"${PACKAGE_NAME}"
  git commit -a -m "Namespace: ${NAMESPACE} Version: ${VERSION}"

  if git tag "${PACKAGE_NAME}"/v"${VERSION}" ; then
    PASS "Golang Package: ${PACKAGE_NAME} Namespace: ${NAMESPACE} Version: ${VERSION} Successfully Tagged:"
    INFO "${PACKAGE_URL}"
  else
    ERROR "Failed to Tag Golang Package: ${PACKAGE_NAME} Namespace: ${NAMESPACE} Version: ${VERSION}:"
    INFO  "${PACKAGE_URL}"
    INFO "This usually indicates an attempt to re-deploy a tagged commit without manually removing the existing tag first."
    debug
    export EXIT_STATUS=1
    # If we fail to tag, we want an immediate exit rather than a delayed exit.
    # It is possible this will result in an inconsistent state which needs to be cleaned up,
    # but hopefully this will interrupt the first deploy in which case there will be no inconsistent state.
    exit 1
  fi

  git remote add deploy "${PACKAGE_AUTH_URL}"
  #git push -u -f deploy main
  INFO "Pushing to deploy remote: ${PACKAGE_URL}"
  if git push -f deploy HEAD:main ; then
    PASS "Golang Package: ${PACKAGE_NAME} Namespace: ${NAMESPACE} Version: ${VERSION} Successfully Deployed to Github Internal:"
    INFO "${PACKAGE_URL}"
  else
    ERROR "Failed to Deploy Golang Package: ${PACKAGE_NAME} Namespace: ${NAMESPACE} Version: ${VERSION} to Github Internal:"
    INFO  "${PACKAGE_URL}"
    debug
    export EXIT_STATUS=1
    exit 1
  fi

  INFO "Update Tags on ${PACKAGE_URL}"
  if git push deploy "${PACKAGE_NAME}"/v"${VERSION}" ; then
    PASS "Golang Package: ${PACKAGE_NAME}/v${VERSION} Namespace: ${NAMESPACE} Version: ${VERSION} Successfully Tagged on Github Internal:"
    INFO "${PACKAGE_URL}"
    PASS "${PACKAGE_URL}" >> "${PROJECT_ROOT}"/verify/golang-release-verify/successful-deployments.txt
  else
    ERROR "Failed to Tag Golang Package: ${PACKAGE_NAME}/v${VERSION} Namespace: ${NAMESPACE} Version: ${VERSION} on Github Internal:"
    INFO  "${PACKAGE_URL}"
    ERROR "${PACKAGE_URL}" >> "${PROJECT_ROOT}"/verify/golang-release-verify/failed-deployments.txt
    debug
    export EXIT_STATUS=1
    exit 1
  fi

  #git push -f deploy "${VERSION}"
  gh release create "${PACKAGE_NAME}/v${VERSION}" -F changelog.md --repo "${PACKAGE_AUTH_URL}"

  # It appears releases can't be pushed to Nutanix internal repositories
  # the org owners don't delegate enough permissions for that purpose.
  #INFO "Uploading Release"
  #if gh release upload "${VERSION}" . ; then
  #  PASS "Golang Package ${PACKAGE_NAME}/v${VERSION} Successfully Released on Github Internal ${PACKAGE_URL}"
  #else
  #  ERROR "Failed to Release Golang Package ${PACKAGE_NAME}/v${VERSION} on Github Internal ${PACKAGE_URL}"
  #  debug
  #  export EXIT_STATUS=1
  #  exit 1
  #fi

  #git push https://"${PASSWORD_PUBLISH_GOLANG}"@github.com/orichter/package-publishing-example-go.git "${DEPLOYMENT_TAG}"

  #git push https://"${PASSWORD_PUBLISH_GOLANG}"@github.com/orichter/package-publishing-example-go.git "${DEPLOYMENT_TAG}"
  #git push origin "${DEPLOYMENT_TAG}"
  #debug
  echo

  popd || exit 1
  popd || exit 1

}

function deploy-to-github-prod {
  NAMESPACE=$1
  VERSION=$2
  DEPLOY_FROM_TAG=$3
  export PACKAGE_NAME="${NAMESPACE}-go-client"

  # HACK: This needs to change for production release
  export REPO_NAME=package-publishing-example-go
  # HACK: This needs to change for production release
  export PACKAGE_URL=https://github.com/orichter/"${REPO_NAME}"/releases/tag/"${VERSION}"
  # HACK: This needs to change for production release
  export PACKAGE_AUTH_URL=https://"${PASSWORD_PUBLISH_GOLANG}"@github.com/orichter/"${REPO_NAME}".git

  #export PACKAGE_URL=https://github.com/orichter/"${REPO_NAME}".git
  export GH_TOKEN=${PASSWORD_PUBLISH_GOLANG}

  # This directory should be created in the verify-private packages step.
  pushd "${PROJECT_ROOT}"/verify/golang-release-verify || exit 1
  # Set publishing credentials
  git config --global user.email "api-packaging@nutanix.com"
  git config --global user.name "Nutanix Circle CI Release Agent"

  # Install Github Release Tools
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
  #git grep -l 'nutanix-core/ntnx-api-golang-sdk-external' | xargs sed -i '' -e 's#nutanix-core/ntnx-api-golang-sdk-external#nutanix/ntnx-api-go-clients#g'
  git checkout "${PACKAGE_NAME}/v${DEPLOY_FROM_TAG}"

  # Something like this should be done in verify-private-packages.sh so we assume it.
  #git clone https://github.com/nutanix-core/ntnx-api-golang-sdk-external
  pushd ./package || exit 1

  pushd "${PACKAGE_NAME}" || exit 1
  INFO "Rewriting repository url"
  # HACK: Modify Sujeets code for Linux compatibility.
  # HACK: Prod replacement
  git grep -l 'nutanix-core/ntnx-api-golang-sdk-external' | xargs sed -i -e 's#nutanix-core/ntnx-api-golang-sdk-external#orichter/package-publishing-example-go#g'
  popd || exit 1

  cp -rf ./"${PACKAGE_NAME}" "${PROJECT_ROOT}"/verify/golang-release-verify/package-github-prod/
  git checkout main

  pushd "${PROJECT_ROOT}"/verify/golang-release-verify/package-github-prod || exit 1
  # The following line can be used to test multiple subsequent versioned commits.
  echo "Testing update to verstion ${VERSION}" >> ./"${PACKAGE_NAME}"/test-updates.txt; git add ./test-updates.txt; # HACK: comment this out.

  INFO "Deploying ntnx-api-golang-sdk-external Namespace: ${NAMESPACE} Version: ${VERSION}"
  git add ./"${PACKAGE_NAME}"
  git commit -a -m "Namespace: ${NAMESPACE} Version: ${VERSION}"

  if git tag "${PACKAGE_NAME}"/v"${VERSION}" ; then
    PASS "Golang Package: ${PACKAGE_NAME} Namespace: ${NAMESPACE} Version: ${VERSION} Successfully Tagged:"
    INFO "${PACKAGE_URL}"
  else
    ERROR "Failed to Tag Golang Package: ${PACKAGE_NAME} Namespace: ${NAMESPACE} Version: ${VERSION}:"
    INFO  "${PACKAGE_URL}"
    INFO "This usually indicates an attempt to re-deploy a tagged commit without manually removing the existing tag first."
    debug
    export EXIT_STATUS=1
    # If we fail to tag, we want an immediate exit rather than a delayed exit.
    # It is possible this will result in an inconsistent state which needs to be cleaned up,
    # but hopefully this will interrupt the first deploy in which case there will be no inconsistent state.
    exit 1
  fi

  git remote add deploy "${PACKAGE_AUTH_URL}"
  #git push -u -f deploy main
  INFO "Pushing to deploy remote: ${PACKAGE_URL}"
  if git push -f deploy HEAD:main ; then
    PASS "Golang Package: ${PACKAGE_NAME} Namespace: ${NAMESPACE} Version: ${VERSION} Successfully Deployed to Github Internal:"
    INFO "${PACKAGE_URL}"
  else
    ERROR "Failed to Deploy Golang Package: ${PACKAGE_NAME} Namespace: ${NAMESPACE} Version: ${VERSION} to Github Internal:"
    INFO  "${PACKAGE_URL}"
    debug
    export EXIT_STATUS=1
    exit 1
  fi

  INFO "Update Tags on ${PACKAGE_URL}"
  if git push deploy "${PACKAGE_NAME}"/v"${VERSION}" ; then
    PASS "Golang Package: ${PACKAGE_NAME}/v${VERSION} Namespace: ${NAMESPACE} Version: ${VERSION} Successfully Tagged on Github Internal:"
    INFO "${PACKAGE_URL}"
    PASS "${PACKAGE_URL}" >> "${PROJECT_ROOT}"/verify/golang-release-verify/successful-deployments.txt
  else
    ERROR "Failed to Tag Golang Package: ${PACKAGE_NAME}/v${VERSION} Namespace: ${NAMESPACE} Version: ${VERSION} on Github Internal:"
    INFO  "${PACKAGE_URL}"
    ERROR "${PACKAGE_URL}" >> "${PROJECT_ROOT}"/verify/golang-release-verify/failed-deployments.txt
    debug
    export EXIT_STATUS=1
    exit 1
  fi

  #git push -f deploy "${VERSION}"
  gh release create "${PACKAGE_NAME}/v${VERSION}" -F changelog.md --repo "${PACKAGE_AUTH_URL}"

  # It appears releases can't be pushed to Nutanix internal repositories
  # the org owners don't delegate enough permissions for that purpose.
  #INFO "Uploading Release"
  #if gh release upload "${VERSION}" . ; then
  #  PASS "Golang Package ${PACKAGE_NAME}/v${VERSION} Successfully Released on Github Internal ${PACKAGE_URL}"
  #else
  #  ERROR "Failed to Release Golang Package ${PACKAGE_NAME}/v${VERSION} on Github Internal ${PACKAGE_URL}"
  #  debug
  #  export EXIT_STATUS=1
  #  exit 1
  #fi

  #git push https://"${PASSWORD_PUBLISH_GOLANG}"@github.com/orichter/package-publishing-example-go.git "${DEPLOYMENT_TAG}"

  #git push https://"${PASSWORD_PUBLISH_GOLANG}"@github.com/orichter/package-publishing-example-go.git "${DEPLOYMENT_TAG}"
  #git push origin "${DEPLOYMENT_TAG}"
  #debug
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
  git log --oneline --decorate -n 10
  git remote -v
  echo
}

if [[ "${BASH_SOURCE[0]}" == "${0}" ]]; then
  main "$@"
  exit "${EXIT_STATUS}"
else
  WARN "This file is designed to be executed, not sourced."
fi
