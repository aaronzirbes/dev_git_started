#!/bin/bash

# Software discovery...
os_version=`uname -s`
git_version=`git version 2> /dev/null`
git_name=`git config --global user.name 2> /dev/null`
git_email=`git config --global user.email 2> /dev/null`
ssh_public_key=`cat ~/.ssh/id_rsa.pub`
homebrew_version=`brew --version 2> /dev/null`


default_git_name=$git_name
if [ "${default_git_name}" == "" ]; then default_git_name=`id -P azirbes | awk -F : '{print $8}'`; fi

default_git_email=$git_email
if [ "${default_git_email}" == "" ]; then default_git_email="${USER}@bloomhealthco.com"; fi

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

function installGitCoreMac() {
    if [ "${git_version}" == "" ] && [ "${os_version}" == "Darwin" ]; then
        if [ "${homebrew_version}" != "" ]; then
            brew install git
        else
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
                    echo "Please follow the installer's on-screen instructions, and then re-run this script"
                    open "${git_installer}"
                else
                    echo "Unable to find installer.  Please run Git Core installer and then re-run this script"
                fi
            else
                echo "Failed to download git core."
            fi
        fi
    fi
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

    if [ "${ssh_public_key}" == "" ]; then
        echo "It appears you do not have an SSH private/public key pair."

        generate_key_pair=$(askYesNo "Do you wish to generate a SSH key pair")

        if [ "${generate_key_pair}" == "Y" ]; then
            if [ "${git_email}" == ""]; then
                echo "You must configure Git before we can generate a SSH key pair"
                exit
            else
                echo "Generating SSH keypair..."
                #ssh-keygen -t rsa -C "${git_email}"
            fi
        else
            echo "You should generated an SSH key pair for use with GitHub so you dont' have to enter your username and password every time."
        fi
    fi

    if [ "${ssh_public_key}" != "" ]; then
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

function installHomebrew() {
    if [ "${homebrew_version}" == "" ] && [ "${os_version}" == "Darwin" ]; then
        echo "${RED}Homebrew was not found on your system.${RESET}"
        install_homebrew=$(askYesNo "Do you wish to install Homebrew (${YELLOW}HIGHLY RECOMMENDED!${RESET})")

        if [ "${install_homebrew}" == "Y" ]; then
            ruby -e "$(curl -fsSkL raw.github.com/mxcl/homebrew/go)"
    fi
    fi
}

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

function setupSandbox() {
}

function main() {
    echo ""
    bloomLogo
    installGitCore
    downloadGiHubMac
    configureGit
}

main
