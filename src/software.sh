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
