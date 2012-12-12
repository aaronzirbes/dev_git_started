#!/bin/bash

####################### GLOBAL ########################

# Software discovery...
os_version=`uname -s`
git_version=`git version 2> /dev/null`
git_name=`git config --global user.name 2> /dev/null`
git_email=`git config --global user.email 2> /dev/null`
ssh_public_key=`cat ~/.ssh/id_rsa.pub`
homebrew_version=`brew --version 2> /dev/null`
github_base_api_url="https://api.github.com/"

# These are all the bloom git repositories that engineering should have cloned locally
git_repositories="
radiant_service_event
lib_service_client
"
git_ping_repo="radiant_service_event"

default_git_name=$git_name
if [ "${default_git_name}" == "" ]; then default_git_name=`id -P azirbes | awk -F : '{print $8}'`; fi

default_git_email=$git_email
if [ "${default_git_email}" == "" ]; then default_git_email="${USER}@bloomhealthco.com"; fi

default_git_sandbox=${BITBUCKET_SANDBOX}
if [ "${default_git_sandbox}" == "" ]; then default_git_sandbox="${HOME}/bloom"; fi

# Colors
RESET=$'\e[0m'
RED=$'\e[1;31m'
GREEN=$'\e[1;32m'
YELLOW=$'\e[1;33m'
BLUE=$'\e[1;34m'
PURPLE=$'\e[1;35m'
CYAN=$'\e[0;36m'
GREY=$'\e[0;37m'
WHITE=$'\e[1;37m'

github_password=""

############################# MAIN ###############################

function main() {
    echo ""
    bloomLogo
    installGitCore
    downloadGiHubMac
    configureGit
    setupGithubUsername
    checkGitHubConnection
    setupSandbox
    verifyAllRepos
}
