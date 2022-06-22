#!/bin/bash
#export TERM=ansi
export TERM=xterm-color
PROJECT_ROOT=${PROJECT_ROOT:-"/home/circleci/project"}
export EXIT_STATUS=0
#shellcheck disable=SC1091
#shellcheck disable=SC1090
source "${PROJECT_ROOT}"/services/deploy/release-utils.source
#shellcheck disable=SC1091
#shellcheck disable=SC1090
source "${PROJECT_ROOT}"/release-config.source
echo "Deploying maven package using Release Params:"
cat "${PROJECT_ROOT}"/release-config.source

# Defualt Deployment Parameters
export GROUP_ID=com.nutanix.example
export ARTIFACT_ID=hellonutanixworld
export VERSION="${DEPLOYMENT_TAG}"
export DEPLOYMENT_FILE=./hellonutanixworld-"${DEPLOYMENT_TAG}".jar
export DEPLOYMENT_POM_FILE=./pom.xml
export DEPLOYMENT_DOC_FILE=./javadoc.jar
export DEPLOYMENT_SOURCES_FILE=./sources.jar
export REPOSITORY_ID=nutanix-public
export REPOSITORY_URL=https://maven.pkg.github.com/orichter/package-publishing-examples

function main {
  import-gpg-keys
  mvn-github-external-release
  mvn-central-external-release
}

function import-gpg-keys {
  echo "${PASSWORD_GPG_KEY}" | base64 -d| gpg --import --no-tty --batch --yes --passphrase "${PASSWORD_PUBLISH_MVN_CENTRAL}" --pinentry-mode loopback
  gpg --keyserver keyserver.ubuntu.com --recv-keys "${GPG_PUBLIC_KEY}"

  #echo -e "${PASSWORD_GPG_KEY}" | base64 -d > PRIVATE_GPG_KEY.key
  #echo "${PASSWORD_PUBLISH_MVN_CENTRAL}" | gpg --import PRIVATE_GPG_KEY.key
  #gpg --no-tty --batch --passphrase "${PASSWORD_PUBLISH_MVN_CENTRAL}" --pinentry-mode loopback --output secrets.env --decrypt secrets.env.gpg
  #DEPLOYMENT_FILE=hellonutanixworld-"${DEPLOYMENT_TAG}".jar

  #gpg-sign "${PRIMARY_JAR_FILE}"
  #gpg-sign pom.xml
  #gpg-sign javadoc.jar
  #gpg-sign sources.jar

  #gpg --verify "${PRIMARY_JAR_FILE}"
  #gpg --verify pom.xml
  #gpg --verify javadoc.jar.asc
  #gpg --verify sources.jar.asc
}

function sign-and-deploy-file {
  pushd "${PROJECT_ROOT}"/mvn-external-release/"${DEPLOYMENT_TAG}" || exit 1

  cp "${PROJECT_ROOT}"/services/deploy/mvn/settings.xml .
  cp "${PROJECT_ROOT}"/mvn/hellonutanixworld/pom.xml .
  sed -i "s/.{env.DEPLOYMENT_TAG}/${DEPLOYMENT_TAG}/g" pom.xml

  echo "This Sample Project has no JavaDocs" > javadoc.md
  echo "This Sample Project has no Sources" > sources.md
  zip javadoc.jar javadoc.md
  zip sources.jar sources.md
  export DEPLOYMENT_POM_FILE=./pom.xml
  export DEPLOYMENT_DOC_FILE=./javadoc.jar
  export DEPLOYMENT_SOURCES_FILE=./sources.jar

  if mvn gpg:sign-and-deploy-file \
    -DgroupId="${GROUP_ID}" \
    -DartifactId="${ARTIFACT_ID}" \
    -Dversion="${VERSION}" \
    -Dfile="${DEPLOYMENT_FILE}" \
    -DpomFile="${DEPLOYMENT_POM_FILE}" \
    -Djavadoc="${DEPLOYMENT_DOC_FILE}" \
    -Dsources="${DEPLOYMENT_SOURCES_FILE}" \
    -DrepositoryId="${REPOSITORY_ID}" \
    -Durl="${REPOSITORY_URL}" \
    --settings settings.xml ; then PASS "Successful Deployment of ${DEPLOYMENT_FILE} version ${VERSION} to ${REPOSITORY_URL}"
  else
    ERROR "Deployment of ${DEPLOYMENT_FILE} version ${VERSION} to ${REPOSITORY_URL} Failed"
    INFO "Debug Info:"
    debug
    EXIT_STATUS=1
  fi

  popd || exit 1
}

function mvn-github-external-release {
  export VERSION="${DEPLOYMENT_TAG}"
  export REPOSITORY_ID=nutanix-public
  export REPOSITORY_URL=https://maven.pkg.github.com/orichter/package-publishing-examples

  sign-and-deploy-file
}

function mvn-central-external-release {
  export REPOSITORY_ID=maven-central
  if [[ "${DEPLOYMENT_TAG}" =~ "-rc" ]]; then
    export REPOSITORY_URL=https://s01.oss.sonatype.org/content/repositories/snapshots
    # Ideally we would do this for a Github release candidate as well,
    # but gpg:sign-and-deploy-file doesnt work with -SNAPSHOT
    # In theory, there is no reason to sign a SNAPSHOT, but in practice,
    # having a single deployment method makes deployments more consistent.
    export VERSION="${DEPLOYMENT_TAG}"-SNAPSHOT
  else
    export REPOSITORY_URL=https://s01.oss.sonatype.org/service/local/staging/deploy/maven2/
    export VERSION="${DEPLOYMENT_TAG}"
  fi

  sign-and-deploy-file
}

function debug {
  env |grep -v PASSWORD
  echo
  pwd
  ls -lah
  git status
  git log --oneline --decorate -n 15
  echo
}

function alternate-approaches {
  echo "This function should not be called, but illustrates some useful code for alternate approaches"
  # Manually Deploy Files
  #mvn deploy:deploy-file -Durl=https://maven.pkg.github.com/orichter/package-publishing-examples -Dpackaging=jar.asc -DrepositoryId=nutanix-public -Dfile=javadoc.jar.asc -DpomFile=pom.xml --settings settings.xml
  #mvn deploy:deploy-file -Durl=https://maven.pkg.github.com/orichter/package-publishing-examples -Dpackaging=jar.asc -DrepositoryId=nutanix-public -Dfile=sources.jar.asc -DpomFile=pom.xml --settings settings.xml

  #GPG_TTY=$(tty)
  #export GPG_TTY

  #if [[ "${DEPLOYMENT_TAG}" =~ "-rc" ]]; then
  #  MVN_DEPLOYMENT_TAG="${DEPLOYMENT_TAG}"-SNAPSHOT
  #else
  #  MVN_DEPLOYMENT_TAG="${DEPLOYMENT_TAG}"
  #fi

  #-Dmaven.wagon.http.pool=false \
  #-Dmaven.wagon.httpconnectionManager.ttlSeconds=25 \
  #-Dmaven.wagon.http.retryHandler.count=3 \
  #-Dpackaging=jar \
  #-Djavadoc=./javadoc.jar \

  #function gpg-sign {
  #  FILE_TO_SIGN=$1
  #  #echo "${PASSWORD_PUBLISH_MVN_CENTRAL}" | gpg --batch --yes --passphrase-fd 0 --pinentry-mode loopback  -ab "${FILE_TO_SIGN}"
  #  gpg -ab "${FILE_TO_SIGN}"
  #}


}

if [[ "${BASH_SOURCE[0]}" == "${0}" ]]; then
  main "$@"
else
  echo "This file is designed to be executed, not sourced."
fi

exit "${EXIT_STATUS}"
