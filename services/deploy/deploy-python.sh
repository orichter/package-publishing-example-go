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
export VERSION="${DEPLOY_TO_TAG}"
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

  cp -rf package stage-package-internal

  pushd stage-package-internal || exit 1

  rm ./dist/*
  python3 -m pip install --upgrade build
  python3 -m build
  python3 -m pip install --upgrade twine
  python3 -m twine upload --repository testpypi --username __token__ --password "${PASSWORD_PUBLISH_TESTPYPI}" dist/*
  popd || exit 1

  popd || exit 1

}

if [[ "${BASH_SOURCE[0]}" == "${0}" ]]; then
  main "$@"
else
  echo "This file is designed to be executed, not sourced."
fi

exit "${EXIT_STATUS}"
