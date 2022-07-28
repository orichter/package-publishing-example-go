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
export VERSION="${DEPLOY_TO_TAG}"
# HACK: Pinned to DEPLOY_FROM_TAG=4.0.1-alpha-1
# This java package is currently pinned to the version below
# rather than inheriting from config.yml. This should be fixed before
# production deployment
export DEPLOY_FROM_TAG=4.0.1-alpha-1
INFO "Version: ${VERSION}"
echo

# Defualt Deployment Parameters
export GROUP_ID=com.nutanix.api
export ARTIFACT_ID=vmm-java-client
export VERSION="${VERSION}"
export DEPLOYMENT_FILE=./"${ARTIFACT_ID}"-"${DEPLOY_FROM_TAG}".jar
export DEPLOYMENT_POM_FILE=./"${ARTIFACT_ID}"-"${DEPLOY_FROM_TAG}".pom
export DEPLOYMENT_DOC_FILE=./javadoc.jar
export DEPLOYMENT_SOURCES_FILE=./sources.jar
export REPOSITORY_ID=nutanix-public
export REPOSITORY_URL=https://maven.pkg.github.com/orichter/package-publishing-examples

function main {
  import-gpg-keys
  mvn-github-external-release
  mvn-central-external-release
  echo
  PASS "Successful Deployments can be found at:"
  cat "${PROJECT_ROOT}"/verify/maven-release-verify/successful-deployments.txt
  if test -f "${PROJECT_ROOT}/verify/maven-release-verify/failed-deployments.txt"; then
    ERROR "Failed Deployments to:"
    cat "${PROJECT_ROOT}"/verify/maven-release-verify/failed-deployments.txt
  fi

}

function import-gpg-keys {
  echo "${PASSWORD_GPG_KEY}" | base64 -d| gpg --import --no-tty --batch --yes --passphrase "${PASSWORD_PUBLISH_MVN_CENTRAL}" --pinentry-mode loopback
  gpg --keyserver keyserver.ubuntu.com --recv-keys "${GPG_PUBLIC_KEY}"

  #echo -e "${PASSWORD_GPG_KEY}" | base64 -d > PRIVATE_GPG_KEY.key
  #echo "${PASSWORD_PUBLISH_MVN_CENTRAL}" | gpg --import PRIVATE_GPG_KEY.key
  #gpg --no-tty --batch --passphrase "${PASSWORD_PUBLISH_MVN_CENTRAL}" --pinentry-mode loopback --output secrets.env --decrypt secrets.env.gpg
  #DEPLOYMENT_FILE=hellonutanixworld-"${VERSION}".jar

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
  #pushd "${PROJECT_ROOT}"/verify/mvn-external-release/"${VERSION}" || exit 1
  pushd "${PROJECT_ROOT}"/verify/maven-release-verify/package || exit 1

  cp "${PROJECT_ROOT}"/services/deploy/mvn/settings.xml .
  #cp "${PROJECT_ROOT}"/mvn/hellonutanixworld/pom.xml .
  #sed -i "s/.{env.DEPLOYMENT_TAG}/${VERSION}/g" pom.xml
  INFO "Depolying using ${DEPLOYMENT_POM_FILE}:"
  cat "${DEPLOYMENT_POM_FILE}"

  # The following can be used to generate stub jar files.
  # These files are required for deployment, and must be added
  # if they don't exist in the underlying repositories.
  #echo "This Sample Project has no JavaDocs" > javadoc.md
  #echo "This Sample Project has no Sources" > sources.md
  #zip javadoc.jar javadoc.md
  #zip sources.jar sources.md
  #export DEPLOYMENT_POM_FILE=./pom.xml
  export DEPLOYMENT_DOC_FILE=./"${ARTIFACT_ID}"-"${DEPLOY_FROM_TAG}"-javadoc.jar
  export DEPLOYMENT_SOURCES_FILE=./"${ARTIFACT_ID}"-"${DEPLOY_FROM_TAG}"-sources.jar

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
    PASS "${PACKAGE_URL}" >> "${PROJECT_ROOT}"/verify/maven-release-verify/successful-deployments.txt
  else
    ERROR "Deployment of ${DEPLOYMENT_FILE} version ${VERSION} to ${REPOSITORY_URL} Failed"
    ERROR "${PACKAGE_URL}" >> "${PROJECT_ROOT}"/verify/maven-release-verify/failed-deployments.txt
    INFO "Debug Info:"
    debug
    EXIT_STATUS=1
  fi

  popd || exit 1
}

function mvn-github-external-release {
  export VERSION="${VERSION}"
  export REPOSITORY_ID=nutanix-public
  export REPOSITORY_URL=https://maven.pkg.github.com/orichter/package-publishing-examples
  # It is unclear if /packages/1552392 is static within github, so it may need to be dropped.
  export PACKAGE_URL="${REPOSITORY_URL//maven.pkg.github.com/github.com}"/packages/1552392

  sign-and-deploy-file
}

function mvn-central-external-release {
  export REPOSITORY_ID=maven-central
  if [[ "${VERSION}" =~ "-rc" ]]; then
    export REPOSITORY_URL=https://s01.oss.sonatype.org/content/repositories/snapshots
    # Ideally we would do this for a Github release candidate as well,
    # but gpg:sign-and-deploy-file doesnt work with -SNAPSHOT
    # In theory, there is no reason to sign a SNAPSHOT, but in practice,
    # having a single deployment method makes deployments more consistent.
    export VERSION="${VERSION}"-SNAPSHOT
    export PACKAGE_URL="${REPOSITORY_URL}/${GROUP_ID//.//}/${ARTIFACT_ID}/${VERSION}"
  else
    export REPOSITORY_URL=https://s01.oss.sonatype.org/service/local/staging/deploy/maven2/
    export VERSION="${VERSION}"
    export PACKAGE_URL="${REPOSITORY_URL}/${GROUP_ID//.//}/${ARTIFACT_ID}/${VERSION}"
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

  #if [[ "${VERSION}" =~ "-rc" ]]; then
  #  VERSION="${VERSION}"-SNAPSHOT
  #else
  #  MVN_VERSION="${VERSION}"
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
