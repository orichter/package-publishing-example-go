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
# Python 2.x has problems with the - character as a deployment tag
# So we strip it.
PYTHON_DEPLOY_TO_TAG=${DEPLOY_TO_TAG//-/}
#PYTHON_DEPLOY_TO_TAG=${DEPLOY_TO_TAG}
# There may also be problems with the _ character
# If so, uncomment the following.
#PYTHON_DEPLOY_TO_TAG=${DEPLOY_TO_TAG//_/}
export VERSION="${PYTHON_DEPLOY_TO_TAG}"
INFO "Deploying npm packages using Release Params:"
cat "${PROJECT_ROOT}"/release-config.source
INFO "Version: ${VERSION}"
echo

function main {
  #deploy-to-stage-internal
  deploy-to-stage
  #deploy-to-github-prod
  #deploy-to-prod
}

function deploy-to-stage {
  pushd "${PROJECT_ROOT}"/pip-release-verify || exit 1

  cp -rf package stage-package

  pushd stage-package || exit 1

  PUBLISH_FROM_NAME=categories-sdk
  PACKAGE_NAME=release-candidate-"${PUBLISH_FROM_NAME}"
  sed -i "s|${PUBLISH_FROM_NAME}|${PACKAGE_NAME}|g" setup.py

  # This override is necessary because the tag in setup.py doesn't match the actual tag.
  # It should be fixed before final deployment
  DEPLOY_FROM_TAG=16.7.0
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
  if python3 -m twine upload --repository testpypi --username __token__ --password "${PASSWORD_PUBLISH_TESTPYPI}" dist/* ; then
    PASS "Python Package ${PACKAGE_NAME} Successfully Deployed to Github Internal ${PACKAGE_URL}"
  else
    ERROR "Failed to Deploy Python Package ${PACKAGE_NAME} to test.pypi.org"
    debug
    export EXIT_STATUS=1
  fi

  popd || exit 1

  popd || exit 1

}

function deploy-to-prod {
  pushd "${PROJECT_ROOT}"/pip-release-verify || exit 1

  cp -rf package prod-package

  pushd prod-package || exit 1

  PUBLISH_FROM_NAME=categories-sdk
  PACKAGE_NAME="${PUBLISH_FROM_NAME}"
  sed -i "s|${PUBLISH_FROM_NAME}|${PACKAGE_NAME}|g" setup.py

  # This override is necessary because the tag in setup.py doesn't match the actual tag.
  # It should be fixed before final deployment
  DEPLOY_FROM_TAG=16.7.0
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
  if python3 -m twine upload --username __token__ --password "${PASSWORD_PUBLISH_TESTPYPI}" dist/* ; then
    PASS "Python Package ${PACKAGE_NAME} Successfully Deployed to Github Internal ${PACKAGE_URL}"
  else
    ERROR "Failed to Deploy Python Package ${PACKAGE_NAME} to test.pypi.org"
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
