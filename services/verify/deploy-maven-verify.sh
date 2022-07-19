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
echo "Verifying maven package deployment using Release Params:"
cat "${PROJECT_ROOT}"/release-config.source

function main {
  #export DEFAULT_REPOSITORY_URL=https://maven.pkg.github.com/nutanix-release-engineering/experiments-example-github-package-npm
  export DEFAULT_REPOSITORY_URL=https://maven.pkg.github.com/nutanix-core/ntnx-api-java-sdk-external
  export DEFAULT_REPOSITORY_ID=nutanix-private
  mvn-github-external-verify
  mvn-central-external-verify
}

function mvn-github-external-verify {
  export VERSION="${DEPLOY_TO_TAG}"
  export REPOSITORY_ID=nutanix-public
  export REPOSITORY_URL=https://maven.pkg.github.com/orichter/package-publishing-examples

  verify-deployment
}

function mvn-central-external-verify {
  export REPOSITORY_ID=maven-central
  if [[ "${DEPLOY_TO_TAG}" =~ "-rc" ]]; then
    export REPOSITORY_URL=https://s01.oss.sonatype.org/content/repositories/snapshots
    # Ideally we would do this for a Github release candidate as well,
    # but gpg:sign-and-deploy-file doesnt work with -SNAPSHOT
    # In theory, there is no reason to sign a SNAPSHOT, but in practice,
    # having a single deployment method makes deployments more consistent.
    export VERSION="${DEPLOY_TO_TAG}"-SNAPSHOT
  else
    export REPOSITORY_URL=https://s01.oss.sonatype.org/service/local/staging/deploy/maven2/
    export VERSION="${DEPLOY_TO_TAG}"
  fi

  verify-deployment
}


function verify-deployment {
  ARTIFACT_ID=vmm-java-client
  rm -rf "${PROJECT_ROOT}"/maven-release-verify
  rm -rf "${HOME}"/.m2
  mkdir -p "${PROJECT_ROOT}"/maven-release-verify

  pushd "${PROJECT_ROOT}"/maven-release-verify || exit 1

  mvn archetype:generate -q -DgroupId=com.nutanix.test -DartifactId=sdk-test-app -DarchetypeArtifactId=maven-archetype-quickstart -DarchetypeVersion=1.4 -DinteractiveMode=false

  pushd sdk-test-app || exit 1
  mv pom.xml pom.xml.old
  cp "${PROJECT_ROOT}"/services/deploy/mvn/pom.xml .

  sed -i "s|.{env.DEPLOYMENT_TAG}|${VERSION}|g" pom.xml
  sed -i "s|${DEFAULT_REPOSITORY_ID}|${REPOSITORY_ID}|g" pom.xml
  sed -i "s|${DEFAULT_REPOSITORY_URL}|${REPOSITORY_URL}|g" pom.xml

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
    ERROR "Maven Test Rig Tests Failed for ${REPOSITORY_ID}:${REPOSITORY_URL}"
    INFO "${OUTPUT}"
  #  debug
  #  EXIT_STATUS=1
  fi
  echo

  # The following test is currently not implemented on maven-central for releases
  # https://central.sonatype.org/publish/release/#releasing-deployment-from-ossrh-to-the-central-repository-introduction
  # https://central.sonatype.org/publish/publish-maven
  # https://github.com/sonatype/nexus-maven-plugins/tree/main/staging/maven-plugin
  if [[ "${DEPLOY_TO_TAG}" =~ "-rc" ]] || [[ "${REPOSITORY_ID}" != "maven-central" ]] ; then
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
    ERROR "Maven Dependency Tests Failed ${REPOSITORY_ID}:${REPOSITORY_URL}"
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
}

if [[ "${BASH_SOURCE[0]}" == "${0}" ]]; then
  main "$@"
else
  echo "This file is designed to be executed, not sourced."
fi

exit "${EXIT_STATUS}"
