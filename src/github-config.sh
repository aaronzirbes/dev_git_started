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

