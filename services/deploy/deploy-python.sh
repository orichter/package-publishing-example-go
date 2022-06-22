#!/bin/bash
#shellcheck disable=SC1091
#export TERM=ansi
export TERM=xterm-color

function debug {
  env |grep -v PASSWORD
  pwd
  ls -lah
  ls -lah ./services
  ls -lah ./services/deploy
}

source ./release-config.source
echo "Deploying python packages with Release Params:"
cat ./release-config.source
echo
pushd ./pypi/packaging_tutorial/ || exit 1
rm ./dist/*
python3 -m pip install --upgrade build
python3 -m build
python3 -m pip install --upgrade twine
python3 -m twine upload --repository testpypi --username __token__ --password "${PASSWORD_PUBLISH_TESTPYPI}" dist/*
popd || exit 1