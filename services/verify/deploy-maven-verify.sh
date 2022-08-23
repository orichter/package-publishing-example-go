#!/bin/bash
#export TERM=ansi
#export TERM=xterm-color
PROJECT_ROOT=${PROJECT_ROOT:-"/home/circleci/project"}
#shellcheck disable=SC1091
#shellcheck disable=SC1090
source "${PROJECT_ROOT}"/services/deploy/release-utils.source

#shellcheck disable=SC1091
#shellcheck disable=SC1090
source "${PROJECT_ROOT}"/release-config.source
if [[ "${SUPPRESS_MVN}" == "true" ]]; then
  WARN "Suppressing mvn debloyments. See manage-package-deployments.sh"
  exit
fi

echo "Verifying maven package deployment using Release Params:"
cat "${PROJECT_ROOT}"/release-config.source

function main {
  #export DEFAULT_REPOSITORY_URL=https://maven.pkg.github.com/nutanix-release-engineering/experiments-example-github-package-npm
  export DEFAULT_REPOSITORY_URL=https://maven.pkg.github.com/nutanix-core/ntnx-api-java-sdk-external
  export DEFAULT_REPOSITORY_ID=nutanix-private

  NAMESPACES="vmm prism clustermgmt aiops iam storage"
  # Iterate the string variable using for loop
  for NAMESPACE in ${NAMESPACES}; do
    echo "MVN Verifying Namespace: ${NAMESPACE}"
    # HACK: Crude deployment manifest which needs to be reworked.
    if [ "${NAMESPACE}" = "storage" ] ; then
      export DEPLOY_FROM_TAG="4.0.1-alpha-2"
      VERSION=${DEPLOY_TO_TAG}-${DEPLOY_FROM_TAG/./-}
    else
      export DEPLOY_FROM_TAG="4.0.1-alpha-1"
      VERSION=${DEPLOY_TO_TAG}-${DEPLOY_FROM_TAG/./-}
    fi

    mvn-github-external-verify "${NAMESPACE}" "${VERSION}"
    mvn-central-external-verify "${NAMESPACE}" "${VERSION}"

  done

}

function mvn-github-external-verify {
  NAMESPACE=$1
  VERSION=$2
  PACKAGE_NAME=${NAMESPACE}-java-client
  ARTIFACT_ID=${PACKAGE_NAME}
  export DEPLOYMENT_FILE=./"${ARTIFACT_ID}"-"${DEPLOY_FROM_TAG}".jar
  export DEPLOYMENT_POM_FILE=./"${ARTIFACT_ID}"-"${DEPLOY_FROM_TAG}".pom

  export REPOSITORY_ID=nutanix-public
  export REPOSITORY_URL=https://maven.pkg.github.com/orichter/package-publishing-examples

  verify-deployment "${NAMESPACE}" "${VERSION}"
}

function mvn-central-external-verify {
  NAMESPACE=$1
  VERSION=$2
  PACKAGE_NAME=${NAMESPACE}-java-client
  export ARTIFACT_ID=${PACKAGE_NAME}
  export DEPLOYMENT_FILE=./"${ARTIFACT_ID}"-"${DEPLOY_FROM_TAG}".jar
  export DEPLOYMENT_POM_FILE=./"${ARTIFACT_ID}"-"${DEPLOY_FROM_TAG}".pom


  export REPOSITORY_ID=maven-central
  if [[ "${VERSION}" =~ "-alpha" ]]; then
    export REPOSITORY_URL=https://s01.oss.sonatype.org/content/repositories/snapshots
    # Ideally we would do this for a Github release candidate as well,
    # but gpg:sign-and-deploy-file doesnt work with -SNAPSHOT
    # In theory, there is no reason to sign a SNAPSHOT, but in practice,
    # having a single deployment method makes deployments more consistent.
    VERIFY_VERSION="${VERSION}"-SNAPSHOT
  else
    export REPOSITORY_URL=https://s01.oss.sonatype.org/service/local/staging/deploy/maven2/
    VERIFY_VERSION="${VERSION}"
  fi

  verify-deployment "${NAMESPACE}" "${VERIFY_VERSION}"
}


function verify-deployment {
  NAMESPACE=$1
  VERSION=$2
  PACKAGE_NAME=${NAMESPACE}-java-client
  INFO "Verifying deployment of ${NAMESPACE} to ${REPOSITORY_ID}: ${REPOSITORY_URL}"

  rm -rf "${PROJECT_ROOT}"/maven-release-verify
  rm -rf "${HOME}"/.m2
  mkdir -p "${PROJECT_ROOT}"/maven-release-verify

  pushd "${PROJECT_ROOT}"/maven-release-verify || exit 1

  mvn archetype:generate -q -DgroupId=com.nutanix.test -DartifactId=sdk-test-app -DarchetypeArtifactId=maven-archetype-quickstart -DarchetypeVersion=1.4 -DinteractiveMode=false

  #pushd test-"${NAMESPACE}"-sdk-app || exit 1
  pushd sdk-test-app || exit 1

  #mv pom.xml pom.xml.old
  rm ./pom.xml
  cp -f "${PROJECT_ROOT}"/services/deploy/mvn/pom.template ./pom.xml

  sed -i "s|.{env.DEPLOYMENT_TAG}|${VERSION}|g" pom.xml
  sed -i "s|${DEFAULT_REPOSITORY_ID}|${REPOSITORY_ID}|g" pom.xml
  sed -i "s|${DEFAULT_REPOSITORY_URL}|${REPOSITORY_URL}|g" pom.xml
  sed -i "s|-java-client|${PACKAGE_NAME}|g" pom.xml

  echo "Using pom.xml:"
  cat ./pom.xml

  cp "${PROJECT_ROOT}"/services/deploy/mvn/settings.xml .

  mvn package -q --settings settings.xml

  EXIT_STATUS=0
  CMD="java -cp target/sdk-test-app-1.0-SNAPSHOT.jar com.nutanix.test.sdk-test-app.App"
  OUTPUT=$(eval "${CMD}")

  if [ "${OUTPUT}" == "Hello World!" ]; then
    PASS "Maven Test Rig Tests Pass for ${REPOSITORY_ID}:${REPOSITORY_URL}"
  else
    #ERROR "Maven Test Rig Tests Failed for ${REPOSITORY_ID}:${REPOSITORY_URL}"
    WARN "Maven Test Rig Tests Not Yet Implemented"
    INFO "${OUTPUT}"
    #debug
    #EXIT_STATUS=1
  fi
  echo

  # The following test is currently not implemented on maven-central for releases
  # https://central.sonatype.org/publish/release/#releasing-deployment-from-ossrh-to-the-central-repository-introduction
  # https://central.sonatype.org/publish/publish-maven
  # https://github.com/sonatype/nexus-maven-plugins/tree/main/staging/maven-plugin
  if [[ "${VERSION}" =~ "-alpha" ]] || [[ "${REPOSITORY_ID}" != "maven-central" ]] ; then
    if mvn -q org.simplify4u.plugins:pgpverify-maven-plugin:check clean install; then
      PASS "Maven PGP Verify Tests Pass for ${REPOSITORY_ID}:${REPOSITORY_URL}"
    else
      ERROR "Maven PGP Verify Failed for ${REPOSITORY_ID}:${REPOSITORY_URL}"
      echo "${OUTPUT}"
      debug
      EXIT_STATUS=1
    fi
  else
    WARN "Maven PGP Verify Tests Not Implemented for ${REPOSITORY_ID}:${REPOSITORY_URL}"

    #gpg --keyserver keyserver.ubuntu.com --recv-keys "${GPG_PUBLIC_KEY}"

    # The following would provide some verification, but there appears to be no easy way to parameterize the "1003" in the deployment which appears to increase monatonically
    #wget "https://s01.oss.sonatype.org/service/local/repositories/comnutanix-1003/content/com/nutanix/example/hellonutanixworld/${VERSION}/hellonutanixworld-${VERSION}.jar"
    #wget "https://s01.oss.sonatype.org/service/local/repositories/comnutanix-1003/content/com/nutanix/example/hellonutanixworld/${VERSION}/hellonutanixworld-${VERSION}.pom"
    #wget "https://s01.oss.sonatype.org/service/local/repositories/comnutanix-1003/content/com/nutanix/example/hellonutanixworld/${VERSION}/hellonutanixworld-${VERSION}.jar.asc"
    #wget "https://s01.oss.sonatype.org/service/local/repositories/comnutanix-1003/content/com/nutanix/example/hellonutanixworld/${VERSION}/hellonutanixworld-${VERSION}.pom.asc"

    #if gpg --verify "hellonutanixworld-${VERSION}.jar.asc" && gpg --verify "hellonutanixworld-${VERSION}.pom.asc" ; then
    #  PASS "JAR and pom files manually verified"
    #else
    #  ERROR "JAR and pom files verification failed"
    #  INFO "Check the GPG_PUBLIC_KEY in circle-ci env settings"
    #  EXIT_STATUS=1
    #fi

    WARN "Manual deployment steps needed at: https://s01.oss.sonatype.org/#stagingRepositories with username nutanix. Contact api-packaging@nutanix.com for the credentials"
  fi
  echo

  CMD="java -cp ${HOME}/.m2/repository/com/api/${ARTIFACT_ID}/${VERSION}/${ARTIFACT_ID}-${VERSION}.jar com.nutanix.api.vmm-java-client.App"
  OUTPUT=$(eval "${CMD}")

  if [ "${OUTPUT}" == "Hello World!" ]; then
    PASS "Maven Dependency Tests Pass for ${REPOSITORY_ID}:${REPOSITORY_URL}"
  else
    #ERROR "Maven Dependency Tests Failed ${REPOSITORY_ID}:${REPOSITORY_URL}"
    WARN "Maven Dependency Test Not Yet Implemented"
    echo "${OUTPUT}"
    #debug
    #EXIT_STATUS=1
  fi
  echo

  popd || exit 1
  popd || exit 1
}

function debug {
  #env |grep -v PASSWORD
  echo
  pwd
  env
  ls -lah
  git status
  git log --oneline --decorate -n 15
  mvn --version
  echo

  WARN "Failed pom.xml:"
  cat pom.xml
  WARN "Failed settings.xml"
  cat settings.xml

}

if [[ "${BASH_SOURCE[0]}" == "${0}" ]]; then
  main "$@"
else
  echo "This file is designed to be executed, not sourced."
fi

exit "${EXIT_STATUS}"
