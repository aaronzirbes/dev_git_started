######################### UTILITIES #################################

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
