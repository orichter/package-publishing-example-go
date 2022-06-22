#!/bin/bash
# manage-release-branch.sh is designed to automatically manage a release branch
# which has gone through the stages dev, test, stage, and prod in the circle.yml
# CircleCI based deployment. It assumes dev deployment runs on merge to master,
# test deployment runs upon tagging a release candidate (i.e. vA.B.C-rcD),
# the stage deployment runs upon tagging a release (i.e. vA.B.C) if this script
# is in the stage portion of the circle.yaml, tagging a release automatically
# creates a release branch from master for a pull request against production
# Merge the release into master (either before or after the stage release)
# In order for this branch to be properly pushed to your repository,
# You must go to https://canaveral-ui.canaveral-corp.us-west-2.aws/team-manager
# and add svc-xi-circlebot to the contributors or external-contributors  team and make sure
# that team has write permissions to your repository
# enabling write permissions for the robot-automation-guest-account in your git repo
# may have the same effect

#export TERM=ansi
export TERM=xterm-color

ENV_TAG=dev
if [ "$1" ];then
  ENV_TAG=$1
fi

MOST_RECENT_COMMIT_WITH_TAGS=$(git rev-list --tags --max-count=1)
MOST_RECENT_TAG=$(git describe --tags --abbrev=0)
DEPLOYMENT_TAG=${FORCE_DEPLOYMENT_TAG:-$MOST_RECENT_TAG}

{
  echo 'export DEPLOYMENT_TAG="'"${DEPLOYMENT_TAG}"'"'
  echo 'export FORCE_DEPLOYMENT_TAG="'"${FORCE_DEPLOYMENT_TAG}"'"'
  echo 'export BRANCH_FILTER="'"${BRANCH_FILTER}"'"'
  echo 'export TAG_FILTER="'"${TAG_FILTER}"'"'
} >> ./release-config.source

# Penultimate: second to last
# Requires a git unshallow or a git clone
#PENULTIMATE_COMMIT_WITH_TAGS=$(git rev-list --tags --max-count=2)
#PENULTIMATE_COMMIT_WITH_TAGS=$(echo "$PENULTIMATE_COMMIT_WITH_TAGS"| awk -F ' ' '{print $2}')
#PENULTIMATE_TAG=$(git describe --tags "$PENULTIMATE_COMMIT_WITH_TAGS")
echo "Most Recent Tag: ${MOST_RECENT_TAG} Commit: ${MOST_RECENT_COMMIT_WITH_TAGS}"
#echo "Penultimate Tag: ${PENULTIMATE_TAG} Commit: ${PENULTIMATE_COMMIT_WITH_TAGS}"

if [ "$ENV_TAG" = "dev" ];then
  echo "Dev Deployment Config"
  cat ./release-config.source

  echo
  echo "Tag Release Candidate Suggestion:"
  echo "When you tag your release candidate we suggest something like the following."
  echo "The first line is the title, everything else is the body:"
  echo
  PR_TITLE=$(git log --pretty=format:"%s" "${MOST_RECENT_TAG}"^..HEAD)
  PR_TITLE=$(echo "${PR_TITLE}"|cut -c -100)
  echo "${PR_TITLE}"
  echo
  echo "Commits since the last release (inclusive):"
  git log --oneline --decorate "${MOST_RECENT_TAG}"^..HEAD
  echo
  echo
  exit 0
fi

if [ "$ENV_TAG" = "test" ];then
  echo "Test Deployment Config"
  cat ./release-config.source

  echo
  echo "Tag Release Candidate Suggestion:"
  echo "When you tag your release candidate we suggest something like the following."
  echo "The first line is the title, everything else is the body:"
  echo
  PR_TITLE=$(git log --pretty=format:"%s" "${MOST_RECENT_TAG}"^..HEAD)
  PR_TITLE=$(echo "${PR_TITLE}"|cut -c -100)
  echo "${PR_TITLE}"
  echo
  echo "Commits since the last release (inclusive):"
  git log --oneline --decorate "${MOST_RECENT_TAG}"^..HEAD
  echo
  echo
  exit 0
fi

if [ "$ENV_TAG" = "stage" ];then
  echo "Stage Deployment Config"
  cat ./release-config.source
  exit 0
  #open pull request with production as the base.
fi

if [ "$ENV_TAG" = "prod" ];then
  echo "Prod Deployment Config"
  cat ./release-config.source

  # Consider implementing a hotfix release branch creation.
  #git config --global user.email "xi-mgmt-gateway@nutanix.com>"
  #git config --global user.name "Circle CI Release Agent"

  #git checkout master
  #git pull

  exit 0
fi

echo "create-release-branch.sh called with invalid parameter!"
echo "parameters are positional and valid parameters are: dev, stage, prod"
exit 1
