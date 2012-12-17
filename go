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

########################################################################################
########### Begin Common Configuration Options #########################################
########################################################################################

# These are all the bloom git repositories that engineering should have cloned locally
git_repositories="
    lib_common
    lib_domain
    webapp_bloomhealth
    webapp_bhbo
    dev_scripts
    dev_config
    test_geb_page_objects
"

# This is the repo used to ping GitHub and check if the connection is setup properly, so everyone should have read access.
git_ping_repo="lib_common"


########################################################################################
########### End Common Configuration Options ###########################################
########################################################################################

#################### BLOOM REPO SETUP ##########################
github_base_api_url="https://api.github.com/"

default_git_sandbox=${BITBUCKET_SANDBOX}
if [ "${default_git_sandbox}" == "" ]; then default_git_sandbox="${HOME}/bloom"; fi

github_password=""

function setupSandbox() {

    if [ "${BLOOM_GIT_SANDBOX}" == "" ]; then
        echo "You do not have a location set for BLOOM_GIT_SANDBOX, the folder where all the bloom git repos will be checked out to."
        read -p "What location would you like to use for your BLOOM_GIT_SANDBOX? [${GREEN}${default_git_sandbox}${RESET}]: " new_bloom_git_sandbox
        if [ "${new_bloom_git_sandbox}" == "" ]; then
            new_bloom_git_sandbox="${default_git_sandbox}"
        fi

        sandbox_dir=`dirname "${new_bloom_git_sandbox}"`

        if [ ! -d "${sandbox_dir}" ]; then
            echo "The path '${RED}${new_bloom_git_sandbox}${RESET}' cannot be used as '${RED}${sandbox_dir}${RESET}' is not a folder."
        else
            export BLOOM_GIT_SANDBOX="${new_bloom_git_sandbox}"
            if [ ! -f ~/.profile ]; then
                touch ~/.profile
            fi
            if (grep -q 'BLOOM_GIT_SANDBOX' ~/.profile); then
                echo "Updating your ${BLUE}BLOOM_GIT_SANDBOX${RESET} environment variable in ${BLUE}~/.profile${RESET}"
                sed -i -e "s/.*BLOOM_GIT_SANDBOX=.*/export BLOOM_GIT_SANDBOX='${BLOOM_GIT_SANDBOX}'/" ~/.profile
            else
                echo "Adding the ${BLUE}BLOOM_GIT_SANDBOX${RESET} environment variable to the end of ${BLUE}~/.profile${RESET}"
                echo "export BLOOM_GIT_SANDBOX='${BLOOM_GIT_SANDBOX}'" >> ~/.profile
            fi
            howToReloadProfile
        fi 
    fi

    if [ ! -d "${BLOOM_GIT_SANDBOX}" ]; then
        echo "Creating folder '${BLUE}${BLOOM_GIT_SANDBOX}${RESET}'."
        mkdir -p "${BLOOM_GIT_SANDBOX}"
    fi
    if [ ! -d "${BLOOM_GIT_SANDBOX}" ]; then
        echo "The folder '${RED}${BLOOM_GIT_SANDBOX}${RESET}' cannot be found or created."
        exit
    fi
}

function verifyAllRepos() {

    if (( ${#git_version} > 0 )); then

        forked_repositories=$(listForkedRepos)

        echo "It looks like you have the following forked repositories on GitHub:"
        for your_repo in ${forked_repositories}; do
            echo " * ${WHITE}${your_repo}${RESET}"
        done

        for repo in ${git_repositories}; do

            # Ensure user has forked all required bloomhealth repos
            if (echo "${forked_repositories}" | grep -q "${repo}"); then
                echo "You have already forked: ${WHITE}${repo}${RESET}.  Good job."
            else
                echo "You have ${RED}not${RESET} forked: ${WHITE}${repo}${RESET}.  Forking via GitHub API...."
                forkRepo ${repo}
                echo "There is some ${WHITE}hardcore forking action${RESET} happening on GiHub right now."
            fi

            # Verify Repo is setup correctly locally
            verifyRepo ${repo}
        done
    else
        echo "You cannot setup the Bloomhealth repos until you have ${WHITE}git${RESET} installed and configured."
    fi
}

function pingRepo() {
    repo_ssh_url="git@github.com:bloomhealth/${git_ping_repo}.git"
        echo ""
    echo "Pinging test Bloomhealth repo to ensure your git is setup properly."
    if (git ls-remote "${repo_ssh_url}"); then
        echo "Git Repo Ping was ${GREEN}successful${RESET}."
        echo ""
    else
        echo "${RED} Failed to access ${WHITE}${repo_ssh_url}${RED}. Please ensure your SSH keys are setup correctly, and you have added your public SSH key to GitHub through the SSH Keys admin screen."
        exit
    fi
}

function verifyRepo() {
    repo_name=$1
    repo_dir="${BLOOM_GIT_SANDBOX}/${repo_name}"
    repo_bloom_ssh_url="git@github.com:bloomhealth/${repo_name}.git"
    repo_user_ssh_url="git@github.com:${GITHUB_USERNAME}/${repo_name}.git"
    found_git_repo=0

    # Verify that the dir is not an Hg repo
    if [ -d "${repo_dir}" ]; then
        pushd "${repo_dir}" > /dev/null
        if (git status 1>&2> /dev/null); then
            echo "Found git repository: ${repo_dir}."
            if (git remote -v |grep -q "upstream.*${repo_bloom_ssh_url}"); then
                echo "  Found remote upstream $WHITE${repo_bloom_ssh_url}${RESET}."
            else
                echo "  ${RED}Missing remote upstream server ${WHITE}${repo_bloom_ssh_url}${RED}.  Adding it.${RESET}"
                git remote add upstream ${repo_bloom_ssh_url} 
            fi
            if (git remote -v |grep -q "origin.*${repo_user_ssh_url}"); then
                echo "  Found remote origin $WHITE${repo_user_ssh_url}${RESET}."
            else
                echo "  ${RED}Missing remote origin server ${WHITE}${repo_user_ssh_url}${RED}.  Adding it.${RESET}"
                git remote add origin ${repo_user_ssh_url}
                echo ""
                echo "  You can fetch the latest repo by issuing the command:"
                echo "    git fetch -u origin master"
            fi
            found_git_repo=1

            # setting head to origin master
            git fetch --all
            git remote set-head origin master > /dev/null
        else
            echo "${WHITE}${repo_dir}${RED} is not a Git repository.  Renaming it to ${repo_dir}-GITBACKUP."
            mv "${repo_dir}" "${repo_dir}-GITBACKUP"
        fi
        popd > /dev/null
    fi

    if (( $found_git_repo == 0 )); then
        echo "Initializing git repository for ${WHITE}${repo_name}${RESET}..."
        mkdir -p "${repo_dir}"
        pushd "${repo_dir}" > /dev/null
        git init > /dev/null
        git remote add origin ${repo_user_ssh_url} > /dev/null
        git remote add upstream ${repo_bloom_ssh_url} > /dev/null
        popd > /dev/null
    fi
}

function listForkedRepos() {
    githubApi GET "user/repos?type=private" |grep '"name":' |sed -E 's#.*"name": "##' |sed -e 's/",.*//' |sort -u
}

function forkRepo() {
    repo_to_fork=$1
    githubApi POST "repos/bloomhealth/${repo_to_fork}/forks" |grep '"name":' |sed -E 's#.*"name": "##' |sed -e 's/",.*//'
}

function checkGitHubConnection() {
    setupGitHubPassword
    checkGitHubPassword
    pingRepo
}

function setupGitHubPassword() {
    echo "" 1>&2
    read -s -p "password [${GREEN}github.com/${GITHUB_USERNAME}${RESET}]: " new_github_password
    export github_password="${new_github_password}"
    echo "" 1>&2
}

function checkGitHubPassword() {
    auth_result=$(githubApi GET "user")
    
    if (echo "${auth_result}" | grep -q '"message": "Bad credentials"'); then
        echo "Invalid GitHub password."
        exit
    fi
}

function githubApi() {
    http_command="${1}"
    url="${2}"
    if [ "$github_password}" == "" ]; then
        echo "You must specify a GitHub password." 1>&2
        exit
    else
        curl -s --request ${http_command} --user "${GITHUB_USERNAME}:${github_password}" -i "${github_base_api_url}${url}" || echo "REST request fails" 1>&2
    fi
}

############################ GIT CONFIGURATION #######################
ssh_public_key=`cat ~/.ssh/id_rsa.pub`

git_name=`git config --global user.name 2> /dev/null`
default_git_name=$git_name
if [ "${default_git_name}" == "" ]; then default_git_name=`id -P azirbes | awk -F : '{print $8}'`; fi

git_email=`git config --global user.email 2> /dev/null`
default_git_email=$git_email
if [ "${default_git_email}" == "" ]; then default_git_email="${USER}@bloomhealthco.com"; fi

function configureGit() {
    if (( "${#git_name}" > 0 )) && (( "${#git_email}" > 0 )) && (( "${#ssh_public_key}" > 0 )); then
        configure_git=$(askYesNo "Do you wish to configure git")
    fi

    if [ "${configure_git}" == "Y" ]; then

        signupGitHub

        read -p "What is your full name to appear on GitHub? [${GREEN}${default_git_name}${RESET}]: " new_git_name
        read -p "What is your email_address to appear on GitHub? [${GREEN}${default_git_email}${RESET}]: " new_git_email

        if [ "${new_git_name}" == "" ]; then new_git_name="${default_git_name}"; fi
        if [ "${new_git_email}" == "" ]; then new_git_email="${default_git_email}"; fi

        if [ "${new_git_name}" != "${git_name}" ]; then
            echo "Setting git name to: ${new_git_name}"
            git config --global user.name "${new_git_name}"
            git_name=new_git_name
        fi
        if [ "${new_git_email}" != "${git_email}" ]; then
            echo "Setting git email to: ${new_git_email}"
            git config --global user.email "${new_git_email}"
            git_email=new_git_email
        fi
    fi

    if (( "${#ssh_public_key}" == 0 )); then
        configure_git='Y'
        echo "It appears you do not have an SSH private/public key pair."

        generate_key_pair=$(askYesNo "Do you wish to generate a SSH key pair")

        if [ "${generate_key_pair}" == "Y" ]; then
            if (( "${#git_email}" == 0 )); then
                echo "You must configure Git before we can generate a SSH key pair"
                exit
            else
                echo "Generating SSH keypair..."
                ssh-keygen -t rsa -C "${git_email}"
            fi
        else
            echo "You should generated an SSH key pair for use with GitHub so you dont' have to enter your username and password every time."
        fi
    fi

    if [ "${ssh_public_key}" != "" ] && [ "${configure_git}" == "Y" ]; then
        echo "The following is your SSH public key to be used on GitHub."
        echo ""
        echo "${GREEN}${ssh_public_key}${RESET}" | fold -w 72
        echo ""

        copy_ssh_key=$(askYesNo "Do you wish to copy this to your clipboard so you can paste it into your GituHub Account settings page")

        if [ "${copy_ssh_key}" == "Y" ]; then
            pbcopy < ~/.ssh/id_rsa.pub

            echo "Please click '${BLUE}Add SSH key${RESET}' in your GitHub ${WHITE}Account Settings -> SSH Keys${RESET} page"
            echo "You can use '${GREEN}Bloom $(hostname -s)${RESET}' for the '${BLUE}Title${RESET}', as that is the name of your computer."
            open "https://github.com/settings/ssh"
        fi
    fi
}

############################ GITHUB SETUP ##########################

function signupGitHub() {
    if (( "${#git_name}" > 0 )) && (( "${#git_email}" > 0 )) && (( "${#ssh_public_key}" > 0 )); then
        have_github_account=$(askYesNo "Do you have a GitHub account")
    fi

    if [ "${have_github_account}" == "N" ]; then
        signup_url='https://github.com/signup/free'
        echo "${RED}Please sign-up for a GitHub account, and then re-run this script.${RESET}"
        echo ""
        echo "Suggested username: ${PURPLE}${USER}${RESET}"
        echo "Suggested email: ${PURPLE}${default_git_email}${RESET}"
        echo "You can ${PURPLE}pick your own${RESET} password."
        echo ""

        open "${signup_url}"
        echo ""
        exit
    fi
}

function setupGithubUsername() {

    if [ "${GITHUB_USERNAME}" == "" ]; then
        echo "You do not have your GitHub username set via the \$GITHUB_USERNAME environemnt variable."
        read -p "What is your GitHub username? [${GREEN}${USER}${RESET}]: " new_github_username
        if [ "${new_github_username}" == "" ]; then
            new_github_username="${USER}"
        fi

        if [ ! -f ~/.profile ]; then
            touch ~/.profile
        fi
        export GITHUB_USERNAME="${new_github_username}"
        if (grep -q 'GITHUB_USERNAME' ~/.profile); then
            echo "Updating your ${BLUE}GITHUB_USERNAME${RESET} environment variable in ${BLUE}~/.profile${RESET}"
            sed -i -e "s/.*GITHUB_USERNAME=.*/export GITHUB_USERNAME='${GITHUB_USERNAME}'/" ~/.profile
        else
            echo "Adding the ${BLUE}GITHUB_USERNAME${RESET} environment variable to the end of ${BLUE}~/.profile${RESET}"
            echo "${WHITE}export GITHUB_USERNAME='${GITHUB_USERNAME}'${RESET}" 
            echo "export GITHUB_USERNAME='${GITHUB_USERNAME}'" >> ~/.profile
        fi
        howToReloadProfile
    fi
}

######################## SOFTWARE INSTALLATION ######################
homebrew_version=`brew --version 2> /dev/null`
git_version=`git version 2> /dev/null`

# This function will download GitHub Mac client
function downloadGiHubMac() {
    app_name='GitHub.app'

    if [ ! -d /Applications/${app_name} ]; then
        dl_url='https://central.github.com/mac/latest'
        dl_file='github-mac-latest.zip'
        dl_location="${HOME}/Downloads/"
        dl_path="${dl_location}${dl_file}"

        if [ ! -d ${dl_location} ]; then
            mkdir -p ${dl_location}
        fi

        if [ ! -d ${dl_location}${app_name} ]; then

            if [ ! -f ${dl_path} ]; then
                echo "Downloading GitHub Mac."
                curl -L ${dl_url} > ${dl_path} || rm -f ${dl_path}
            fi

            if [ -f ${dl_path} ]; then
                echo "Extracting GitHub Mac..."
                unzip -q -d ${dl_location} ${dl_path}
            else
                echo "Failed to download GitHub Mac."
            fi
        fi

        if [ -d ${dl_location}${app_name} ]; then
            echo "Installing GitHub Mac."
            mv ${dl_location}${app_name} /Applications/ && open /Applications/
        else
            echo "Failed to install GitHub Mac."
            echo "Expecting: ${dl_location}${app_name}"
        fi
    else
        echo "${BLUE}GitHub Mac${RESET} is already installed."
    fi
}

function installGitCore() {
    if [ "${git_version}" == "" ]; then
        if [ "${os_version}" == "Linux" ]; then
            installGitCoreLinux
        elif [ "${os_version}" == "Darwin" ]; then
            installGitCoreMac
        fi
    else
        echo "${BLUE}Git${RESET} is already installed."
    fi
}

function installGitCoreLinux() {
    if [ "${git_version}" == "" ] && [ "${os_version}" == "Linux" ]; then
        sudo apt-get -y -y install git-core
    fi
}

function installGitFlowMac() {
    if [ "${git_version}" == "" ] && [ "${os_version}" == "Darwin" ]; then
        if [ "${homebrew_version}" != "" ]; then
            brew install git-flow
        else
            gitflow_install_url="https://github.com/nvie/gitflow/wiki/Mac-OS-X"
            dl_url='https://github.com/downloads/timcharper/git_osx_installer/git-1.8.0.1-intel-universal-snow-leopard.dmg'
            dl_file='git-mac-latest.dmg'
            dl_location="${HOME}/Downloads/"
            dl_path="${dl_location}${dl_file}"

            if [ ! -d ${dl_location} ]; then
                mkdir -p ${dl_location}
            fi

            if [ ! -f ${dl_path} ]; then
                echo "Downloading git core."
                curl -L ${dl_url} > ${dl_path} || rm -f ${dl_path}
            fi

            if [ -f ${dl_path} ]; then
                echo "Mounting git core DMG..."
                hdiutil attach ${dl_path}

                git_installer=`ls -d /Volumes/Git*/*.pkg`
                if [ "${git_installer}" != "" ]; then
                    echo "${GREEN}Please follow the installer's on-screen instructions, and then re-run this script.${RESET}"
                    echo ""
                    open "${git_installer}"
                else
                    echo "Unable to find installer.  Please run Git Core installer and then re-run this script"
                fi
                echo "You will need to install git-flow ${RED}by hand${RESET} since you do not have ${WHITE}homebrew${RESET} installed."
                open "${gitflow_install_url}"
                echo ""
            else
                echo "Failed to download git core."
            fi
        fi
    fi
}

function installGitCoreMac() {
    if [ "${git_version}" == "" ] && [ "${os_version}" == "Darwin" ]; then
        if [ "${homebrew_version}" != "" ]; then
            brew install git
        else
            gitflow_install_url="https://github.com/nvie/gitflow/wiki/Mac-OS-X"
            dl_url='https://github.com/downloads/timcharper/git_osx_installer/git-1.8.0.1-intel-universal-snow-leopard.dmg'
            dl_file='git-mac-latest.dmg'
            dl_location="${HOME}/Downloads/"
            dl_path="${dl_location}${dl_file}"

            if [ ! -d ${dl_location} ]; then
                mkdir -p ${dl_location}
            fi

            if [ ! -f ${dl_path} ]; then
                echo "Downloading git core."
                curl -L ${dl_url} > ${dl_path} || rm -f ${dl_path}
            fi

            if [ -f ${dl_path} ]; then
                echo "Mounting git core DMG..."
                hdiutil attach ${dl_path}

                git_installer=`ls -d /Volumes/Git*/*.pkg`
                if [ "${git_installer}" != "" ]; then
                    echo "${GREEN}Please follow the installer's on-screen instructions, and then re-run this script.${RESET}"
                    echo ""
                    open "${git_installer}"
                else
                    echo "Unable to find installer.  Please run Git Core installer and then re-run this script"
                fi
                echo "You will need to install git-flow ${RED}by hand${RESET} since you do not have ${WHITE}homebrew${RESET} installed."
                open "${gitflow_install_url}"
                echo ""
            else
                echo "Failed to download git core."
            fi
        fi
    fi
}

function installHomebrew() {
    if [ "${homebrew_version}" == "" ] && [ "${os_version}" == "Darwin" ]; then
        echo "${RED}Homebrew was not found on your system.${RESET}"
        install_homebrew=$(askYesNo "Do you wish to install Homebrew (${YELLOW}HIGHLY RECOMMENDED!${RESET})")

        if [ "${install_homebrew}" == "Y" ]; then
            ruby -e "$(curl -fsSkL raw.github.com/mxcl/homebrew/go)"
        fi
    fi
}
######################### UTILITIES #################################

# Software discovery...
os_version=`uname -s`

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

function bloomLogo() {

    echo "${YELLOW}MMM.                 MMMMMMM${RESET}"
    echo "${YELLOW} MMMM                MMMMMMM${RESET}"
    echo "${YELLOW} MMMM.               MMMMMMM${RESET}"
    echo "${YELLOW} MMMMM               MMMMMMM${RESET}"
    echo "${YELLOW} MMMMM               MMMMMMM${RESET}"
    echo "${YELLOW} MMMMM               MMMMMMM${RESET}"
    echo "${YELLOW} MMMMM.MMMMMMMMM.    MMMMMMM   ..MMMMMMMMM         . MMMMMMMMM .         MMMMMMMM..  MMMMMMMMM.${RESET}"
    echo "${YELLOW} MMMMMMMMMMMMMMMMMM  MMMMMMM  MMMMMMMMMMMMMMM.    MMMMMMMMMMMMMMM     MMMMMMMMMMMMMMMMMMMMMMMMMM${RESET}"
    echo "${YELLOW} MMMMMMMMMMMMMMMMMMM MMMMMMM.MMMMMMMMMMMMMMMMMM. MMMMMMMMMMMMMMMMM   MMMMMMMMMMMMMMMMMMMMMMMMMMMMM${RESET}"
    echo "${YELLOW} MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM${RESET}"
    echo "${YELLOW} MMMMMMMM.   .MMMMMMMMMMMMMMMMMMMMM.. . MMMMMMMMMMMMMMM    .MMMMMMMMMMMMMMM...MMMMMMMMMM   MMMMMMMM.${RESET}"
    echo "${YELLOW} MMMMMM         MMMMMMMMMMMMMMMMM.        MMMMMMMMMMM         MMMMMMMMMMMM     .MMMMMMM     MMMMMMM.${RESET}"
    echo "${YELLOW} MMMMM.          MMMMMMMMMMMMMMM.          MMMMMMMMM.          MMMMMMMMMMM     .MMMMMM      .MMMMMM.${RESET}"
    echo "${YELLOW} MMMMM           MMMMMMMMMMMMMMM           MMMMMMMMM           MMMMMMMMMM      .MMMMMM      .MMMMMM.${RESET}"
    echo "${YELLOW} MMMMM           MMMMMMMMMMMMMMM           MMMMMMMMM           MMMMMMMMMM      .MMMMMM      .MMMMMM.${RESET}"
    echo "${YELLOW} MMMMMM.        MMMMMMMMMMMMMMMMM.        MMMMMMMMMMM        .MMMMMMMMMMM      .MMMMMM      .MMMMMM.${RESET}"
    echo "${YELLOW} MMMMMMM       MMMMMMMMMMMMMMMMMMM.      MMMMMMMMMMMMM.     .MMMMMMMMMMMM      .MMMMMM      .MMMMMM.${RESET}"
    echo "${YELLOW}  MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM      .MMMMMM      .MMMMMM${RESET}"
    echo "${YELLOW}  MMMMMMMMMMMMMMMMMMM.MMMMMMMMMMMMMMMMMMMMMMMMM.MMMMMMMMMMMMMMMMMMMMMMMMM      .MMMMMM      .MMMMMM${RESET}"
    echo "${YELLOW}    MMMMMMMMMMMMMMM    MMMMM  MMMMMMMMMMMMMMM    .MMMMMMMMMMMMMMM  MMMMMM      .MMMMMM      .MMMMM.${RESET}"
    echo "${YELLOW}     ..MMMMMMMMM..      MMMM   . MMMMMMMMM         ..MMMMMMMMM     MMMMMM      .MMMMMM      .MMM${RESET}"
    echo "${YELLOW}        ...  ...         .         ..  .  .            ...   .     .      .     .             ..${RESET}"
}

function askYesNo() {
    prompt_text="${1}"
    while true; do
        read -p "${prompt_text}? [${GREEN}y/n${RESET}]: " yn
        case ${yn} in
            [Yy]* ) result='Y'; break;;
            [Nn]* ) result='N'; break;;
            * ) echo "${RED}Please choose (y)es or (n)o.${RESET}" 1>&2
        esac
    done
    echo ${result}
}

function howToReloadProfile() {
    echo "You'll have to run the following command to re-load your ~/.profile (don't forget the leading '.'):"
    echo ""
    echo ". ~/.profile"
    echo ""
}

main

