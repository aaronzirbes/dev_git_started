#!/bin/bash

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

