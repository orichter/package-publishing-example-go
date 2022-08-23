#!/bin/bash
#export TERM=ansi
export TERM=xterm-color
PROJECT_ROOT=${PROJECT_ROOT:-"/home/circleci/project"}
EXIT_STATUS=0
#shellcheck disable=SC1091
#shellcheck disable=SC1090
source "${PROJECT_ROOT}"/services/deploy/release-utils.source

#shellcheck disable=SC1091
#shellcheck disable=SC1090
source "${PROJECT_ROOT}"/release-config.source
if [[ "${SUPPRESS_PIP}" == "true" ]]; then
  WARN "Suppressing pip debloyments. See manage-package-deployments.sh"
  exit
fi

# Python 2.x has problems with the - character as a deployment tag
# So we strip it.
PYTHON_DEPLOY_TO_TAG=${DEPLOY_TO_TAG//-/}
#PYTHON_DEPLOY_TO_TAG=${DEPLOY_TO_TAG}
# There may also be problems with the _ character
# If so, uncomment the following.
#PYTHON_DEPLOY_TO_TAG=${DEPLOY_TO_TAG//_/}
export VERSION="${PYTHON_DEPLOY_TO_TAG}"
INFO "Deploying pip packages using Release Params:"
cat "${PROJECT_ROOT}"/release-config.source
INFO "Version: ${VERSION}"
echo

function main {

  NAMESPACES="vmm prism clustermgmt aiops iam storage"
  # Iterate the string variable using for loop
  for NAMESPACE in ${NAMESPACES}; do
    echo "PIP Deploying Namespace: ${NAMESPACE}"
    # HACK: Crude deployment manifest which needs to be reworked.
    if [ "${NAMESPACE}" = "storage" ] ; then
      export DEPLOY_FROM_TAG="4.0.1-alpha.2"
      VERSION=${DEPLOY_TO_TAG}${DEPLOY_FROM_TAG//./}
    else
      export DEPLOY_FROM_TAG="4.0.1-alpha.1"
      VERSION=${DEPLOY_TO_TAG}${DEPLOY_FROM_TAG//./}
    fi

    #deploy-to-stage-internal "${NAMESPACE}" "${VERSION}"
    deploy-to-stage "${NAMESPACE}" "${VERSION}"
    #deploy-to-github-prod
    #deploy-to-prod

  done

  PASS "Successful Deployments can be found at:"
  cat "${PROJECT_ROOT}"/verify/pip-release-verify/successful-deployments.txt
  if test -f "${PROJECT_ROOT}/pip-release-verify/failed-deployments.txt"; then
    ERROR "Failed Deployments to:"
    cat "${PROJECT_ROOT}"/verify/pip-release-verify/failed-deployments.txt
  fi

}

function deploy-to-stage {
  NAMESPACE=$1
  VERSION=$2
  export PACKAGE_NAME="ntnx_${NAMESPACE}_py_client"
  #export PACKAGE_NAME="${PACKAGE_NAME//_/-}"

  pushd "${PROJECT_ROOT}"/verify/pip-release-verify || exit 1

  echo "package-${NAMESPACE}"
  echo "stage-package-${NAMESPACE}"
  cp -rf package-"${NAMESPACE}" stage-package-"${NAMESPACE}"
  ls -lah

  pushd stage-package-"${NAMESPACE}" || exit 1

  PUBLISH_FROM_NAME=${PACKAGE_NAME}
  #PACKAGE_NAME=release-candidate-"${PUBLISH_FROM_NAME}"
  sed -i "s|${PUBLISH_FROM_NAME}|${PACKAGE_NAME}|g" setup.py

  sed -i "s|${DEPLOY_FROM_TAG}|${VERSION}|g" setup.py

  INFO "Rewritten Setup.py"
  cat setup.py

  python3 -m pip install --upgrade build > /dev/null
  python3 -m pip install --upgrade twine > /dev/null

  rm ./dist/*

  INFO "Building Pythong Package for test.pypi.org remote"
  if python3 -m build > /dev/null ; then
    PASS "Python Package ${PACKAGE_NAME} Successfully Built for test.pypi.org"
  else
    ERROR "Failed to Build Python Package ${PACKAGE_NAME} for test.pypi.org"
    debug
    export EXIT_STATUS=1
  fi

  INFO "Pushing to test.pypi.org remote"
  PACKAGE_URL=https://test.pypi.org/project/"${PACKAGE_NAME}/${VERSION}"
  if python3 -m twine upload --repository testpypi --username __token__ --password "${PASSWORD_PUBLISH_TESTPYPI}" dist/* ; then
    PASS "Python Package ${PACKAGE_NAME} Successfully Deployed to ${PACKAGE_URL}"
    PASS "${PACKAGE_URL}" >> "${PROJECT_ROOT}"/verify/pip-release-verify/successful-deployments.txt
  else
    ERROR "Failed to Deploy Python Package ${PACKAGE_NAME} to ${PACKAGE_URL}"
    ERROR "${PACKAGE_URL}" >> "${PROJECT_ROOT}"/verify/pip-release-verify/failed-deployments.txt
    debug
    export EXIT_STATUS=1
  fi

  popd || exit 1

  popd || exit 1

}

function deploy-to-prod {
  NAMESPACE=$1
  VERSION=$2
  export PACKAGE_NAME="ntnx_${NAMESPACE}_py_client"
  #export PACKAGE_NAME="${PACKAGE_NAME//_/-}"

  pushd "${PROJECT_ROOT}"/verify/pip-release-verify || exit 1

  cp -rf package-"${NAMESPACE}" prod-package-internal-"${NAMESPACE}"

  pushd prod-package-"${NAMESPACE}" || exit 1

  PUBLISH_FROM_NAME=${PACKAGE_NAME}
  #PACKAGE_NAME=release-candidate-"${PUBLISH_FROM_NAME}"

  sed -i "s|${PUBLISH_FROM_NAME}|${PACKAGE_NAME}|g" setup.py

  sed -i "s|${DEPLOY_FROM_TAG}|${VERSION}|g" setup.py

  INFO "Rewritten Setup.py"
  cat setup.py

  python3 -m pip install --upgrade build > /dev/null
  python3 -m pip install --upgrade twine > /dev/null

  rm ./dist/*

  INFO "Building Pythong Package for test.pypi.org remote"
  if python3 -m build > /dev/null ; then
    PASS "Python Package ${PACKAGE_NAME} Successfully Built for test.pypi.org"
  else
    ERROR "Failed to Build Python Package ${PACKAGE_NAME} for test.pypi.org"
    debug
    export EXIT_STATUS=1
  fi

  INFO "Pushing to test.pypi.org remote"
  # HACK: We are currently using test.pypi.org for prod deployments. We need to remember to revert the if statement
  # below for final publication.
  #if python3 -m twine upload --username __token__ --password "${PASSWORD_PUBLISH_TESTPYPI}" dist/* ; then
  INFO "Pushing to test.pypi.org remote"
  PACKAGE_URL=https://test.pypi.org/project/"${PACKAGE_NAME}/${VERSION}"
  if python3 -m twine upload --repository testpypi --username __token__ --password "${PASSWORD_PUBLISH_TESTPYPI}" dist/* ; then
    PASS "Python Package ${PACKAGE_NAME} Successfully Deployed to ${PACKAGE_URL}"
    PASS "${PACKAGE_URL}" >> "${PROJECT_ROOT}"/verify/pip-release-verify/successful-deployments.txt
  else
    ERROR "Failed to Deploy Python Package ${PACKAGE_NAME} to ${PACKAGE_URL}"
    ERROR "${PACKAGE_URL}" >> "${PROJECT_ROOT}"/verify/pip-release-verify/failed-deployments.txt
    debug
    export EXIT_STATUS=1
  fi

  popd || exit 1

  popd || exit 1

}

function debug {
  #env |grep -v PASSWORD
  echo
  pwd
  rm ./dist/*
  python3 -m build
  python3 -m twine upload --repository testpypi --username __token__ --password "${PASSWORD_PUBLISH_TESTPYPI}" --verbose dist/*
  ls -lah
  git status
  git log --oneline --decorate -n 15
  git remote -v
  echo
}

if [[ "${BASH_SOURCE[0]}" == "${0}" ]]; then
  main "$@"
else
  echo "This file is designed to be executed, not sourced."
fi

exit "${EXIT_STATUS}"
