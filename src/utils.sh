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
