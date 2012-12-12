#################### BLOOM REPO SETUP ##########################

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
        git remote add origin ${repo_user_ssh_url}
        git remote add upstream ${repo_bloom_ssh_url}
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

