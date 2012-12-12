############################ GIT CONFIGURATION #######################

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
        configure_git='Y'
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

