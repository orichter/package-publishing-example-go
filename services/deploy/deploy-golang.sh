#!/bin/bash
#export TERM=ansi
export TERM=xterm-color

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
echo "Deploying golang package using Release Params:"
cat ./release-config.source
mkdir -p ./golang-release
pushd ./golang-release || exit 1
#git clone git@github.com:orichter/package-publishing-example-go.git
git config --global user.email "api-packaging@nutanix.com"
git config --global user.name "Nutanix Circle CI Release Agent"

git clone https://github.com/orichter/package-publishing-example-go.git
pushd ./package-publishing-example-go || exit 1
cp -rf ../../golang/* .
go mod tidy
go test ./...
git add ./*
git commit -m "hello: changes for ${DEPLOYMENT_TAG}"
git tag "${DEPLOYMENT_TAG}"
#URL=https://"${PASSWORD_PUBLISH_GOLANG}"@github.com/orichter/package-publishing-example-go.git
#echo "${URL}"

curl -fsSL https://cli.github.com/packages/githubcli-archive-keyring.gpg | sudo dd of=/usr/share/keyrings/githubcli-archive-keyring.gpg
echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/githubcli-archive-keyring.gpg] https://cli.github.com/packages stable main" | sudo tee /etc/apt/sources.list.d/github-cli.list > /dev/null
sudo apt update
sudo apt install gh

echo "${PASSWORD_PUBLISH_GOLANG}" | gh auth login --with-token

git push https://"${PASSWORD_PUBLISH_GOLANG}"@github.com/orichter/package-publishing-example-go.git
git push https://"${PASSWORD_PUBLISH_GOLANG}"@github.com/orichter/package-publishing-example-go.git "${DEPLOYMENT_TAG}"
gh release create "${DEPLOYMENT_TAG}" -F changelog.md
gh release upload "${DEPLOYMENT_TAG}" .
#git push https://"${PASSWORD_PUBLISH_GOLANG}"@github.com/orichter/package-publishing-example-go.git "${DEPLOYMENT_TAG}"

#git push https://"${PASSWORD_PUBLISH_GOLANG}"@github.com/orichter/package-publishing-example-go.git "${DEPLOYMENT_TAG}"
#git push origin "${DEPLOYMENT_TAG}"
debug
echo

popd || exit 1
popd || exit 1
