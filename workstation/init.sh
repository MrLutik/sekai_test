#!/bin/bash
ETC_PROFILE="/etc/profile"
set +e && chmod 555 $ETC_PROFILE && source $ETC_PROFILE &>/dev/null && set -e

INFRA_BRANCH="${1,,}"
SKIP_UPDATE=$2
START_TIME_INIT=$3

[ ! -z "$SUDO_USER" ] && KIRA_USER=$SUDO_USER
[ -z "$KIRA_USER" ] && KIRA_USER=$USER

[ "$KIRA_USER" == "root" ] && KIRA_USER=$(logname)
if [ "$KIRA_USER" == "root" ]; then
    echo "ERROR: You must login as non root user to your machine!"
    exit 1
fi

if [ "${USER,,}" != root ]; then
    echo "ERROR: You have to run this application as root, try 'sudo -s' command first"
    exit 1
fi

# Used To Initialize essential dependencies, MUST be iterated if essentials require updating
CDHELPER_VERSION="v0.6.51"
INFRA_REPO="https://github.com/KiraCore/kira"
ARCHITECTURE=$(uname -m)

[ -z "$INFRA_BRANCH" ] && INFRA_BRANCH="master"
[ -z "$START_TIME_INIT" ] && START_TIME_INIT="$(date -u +%s)"
[ -z "$SKIP_UPDATE" ] && SKIP_UPDATE="false"
[ -z "$DEFAULT_SSH_PORT" ] && DEFAULT_SSH_PORT="22"

DEFAULT_P2P_PORT="26656"
DEFAULT_RPC_PORT="26657"
DEFAULT_PROMETHEUS_PORT="26660"
DEFAULT_GRPC_PORT="9090"
DEFAULT_INTERX_PORT="11000"

KIRA_INTERX_PORT="11000"

KIRA_SEED_P2P_PORT="16656"
KIRA_SEED_RPC_PORT="16657"
KIRA_SEED_GRPC_PORT="19090"
KIRA_SEED_PROMETHEUS_PORT="16660"

KIRA_SENTRY_RPC_PORT="26657"
KIRA_SENTRY_P2P_PORT="26656"
KIRA_SENTRY_GRPC_PORT="29090"
KIRA_SENTRY_PROMETHEUS_PORT="26660"

KIRA_VALIDATOR_P2P_PORT="36656"
KIRA_VALIDATOR_RPC_PORT="36657"
KIRA_VALIDATOR_GRPC_PORT="39090"
KIRA_VALIDATOR_PROMETHEUS_PORT="36660"

KIRA_REGISTRY_PORT="5000"

KIRA_HOME="/home/$KIRA_USER"
KIRA_DUMP="$KIRA_HOME/dump"
KIRA_SNAP="$KIRA_HOME/snap"
KIRA_SECRETS="$KIRA_HOME/.secrets"
KIRA_CONFIGS="$KIRA_HOME/.kira"
SETUP_LOG="$KIRA_DUMP/setup.log"

echo "------------------------------------------------"
echo "|      STARTED: INIT"
echo "|-----------------------------------------------"
echo "|  SKIP UPDATE: $SKIP_UPDATE"
echo "|   START TIME: $START_TIME_INIT"
echo "| INFRA BRANCH: $INFRA_BRANCH"
echo "|   INFRA REPO: $INFRA_REPO"
echo "|    KIRA USER: $KIRA_USER"
echo "| ARCHITECTURE: $ARCHITECTURE"
echo "------------------------------------------------"

rm -rfv $KIRA_DUMP
mkdir -p "$KIRA_DUMP" "$KIRA_SNAP" "$KIRA_CONFIGS" "$KIRA_SECRETS" "/var/kiraglob"

set +x
if [ -z "$SKIP_UPDATE" ]; then
    echo -e "\e[35;1mMMMMMMMMMWX0kdloxOKNWMMMMMMMMMMMMMMMMMMMMMMMMMMM"
    echo "MMMMMWNKOxlc::::::cok0XWWMMMMMMMMMMMMMMMMMMMMMMM"
    echo "MMWX0kdlc::::::::::::clxkOKNMMMMMMMMMMWKkk0NWMMM"
    echo "MNkoc:::::::::::::::::::::cok0NWMMMMMMWKxlcld0NM"
    echo "W0l:cllc:::::::::::::::::::::coKWMMMMMMMWKo:;:xN"
    echo "WOlcxXNKOdlc::::::::::::::::::l0WMMMMMWNKxc;;;oX"
    echo "W0olOWMMMWX0koc::::::::::::ldOXWMMMWXOxl:;;;;;oX"
    echo "MWXKNMMMMMMMWNKOdl::::codk0NWMMWNKkdc:;;;;;;;;oX"
    echo "MMMMMMMMMMMMMMMMWX0kkOKNWMMMWX0xl:;;;;;;;;;;;;oX"
    echo "MMMMMMMMMWXOkOKNMMMMMMMMMMMW0l:;;;;;;;;;;;;;;;oX"
    echo "MMMMMMMMMXo:::cox0XWMMMMMMMNx:;;;;;;;;;;;;;;;;oX"
    echo "MMMMMMMMMKl:::::::ldOXWMMMMNx:;;;;;;;;;;;;;;co0W"
    echo "MMMMMMMMMKl::::;;;;;:ckWMMMNx:;;;;;;;;;;:ldOKNMM"
    echo "MMMMMMMMMKl;;;;;;;;;;;dXMMMNx:;;;;;;;:ox0XWMMMMM"
    echo "MMMMMMMMMKl;;;;;;;;;;;dXMMMWk:;;;:cdkKNMMMMMMMMM"
    echo "MMMMMMMMMKl;;;;;;;;;;;dXMMMMXkoox0XWMMMMMMMMMMMM"
    echo "MMMMMMMMMKl;;;;;;;;;;;dXMMMMMWWWMMMMMMMMMMMMMMMM"
    echo "MMMMMMMMMKl;;;;;;;;;;;dXMMMMMMMMMMMMMMMMMMMMMMMM"
    echo "MMMMMMMMMKo;;;;;;;;;;;dXMMMMMMMMMMMMMMMMMMMMMMMM"
    echo "MMMMMMMMMWKxl:;;;;;;;;oXMMMWNWMMMMMMMMMMMMMMMMMM"
    echo "MMMMMMMMMMMWNKkdc;;;;;:dOOkdlkNMMMMMMMMMMMMMMMMM"
    echo "MMMMMMMMMMMMMMMWXOxl:;;;;;cokKWMMMMMMMMMMMMMMMMM"
    echo "MMMMMMMMMMMMMMMMMMWN0kdxxOKWMMMMMMMMMMMMMMMMMMMM"
    echo "M              KIRA NETWORK SETUP              M"
    echo -e "MMMMMMMMMMMMMMMMMMMMMMWWMMMMMMMMMMMMMMMMMMMMMMMM\e[0m\c\n"
    sleep 3
else
    echo "INFO: Initalizing setup script..."
fi

systemctl stop kiraup || echo "WARNING: KIRA Update service could NOT be stopped, service might not exist yet!"
systemctl stop kiraplan || echo "WARNING: KIRA Upgrade Plan service could NOT be stopped, service might not exist yet!"

echo -n ""
set -x
rm -fv /dev/null && mknod -m 666 /dev/null c 1 3 || :

CPU_CORES=$(cat /proc/cpuinfo | grep processor | wc -l || echo "0")
RAM_MEMORY=$(grep MemTotal /proc/meminfo | awk '{print $2}' || echo "0")

if [[ $CPU_CORES -lt 2 ]] ; then
    echo -en "\e[31;1mERROR: KIRA Manager requires at lest 2 CPU cores but your machine has only $CPU_CORES\e[0m"
    echo "INFO: Recommended CPU is 4 cores"
    echo -en "\e[31;1mPress any key to continue or Ctrl+C to abort...\e[0m" && read -n 1 -s && echo ""
fi

if [[ $RAM_MEMORY -lt 3145728 ]] ; then
    echo -en "\e[31;1mERROR: KIRA Manager requires at lest 4 GB RAM but your machine has only $RAM_MEMORY kB\e[0m"
    echo "INFO: Recommended RAM is 8GB"
    echo -en "\e[31;1mPress any key to continue or Ctrl+C to abort...\e[0m" && read -n 1 -s && echo ""
fi

# All branches should have the same name across all repos to be considered compatible
if [[ $INFRA_BRANCH == mainnet* ]] || [[ $INFRA_BRANCH == testnet* ]] ; then
    DEFAULT_BRANCH="$INFRA_BRANCH"
    SEKAI_BRANCH="$DEFAULT_BRANCH"
    INTERX_BRANCH="$DEFAULT_BRANCH"
else
    DEFAULT_BRANCH="master"
    [ -z "$SEKAI_BRANCH" ] && SEKAI_BRANCH="$DEFAULT_BRANCH"
    [ -z "$INTERX_BRANCH" ] && INTERX_BRANCH="$DEFAULT_BRANCH"
fi

[ -z "$SEKAI_REPO" ] && SEKAI_REPO="https://github.com/KiraCore/sekai"
[ -z "$INTERX_REPO" ] && INTERX_REPO="https://github.com/KiraCore/sekai"

if [ "${SKIP_UPDATE,,}" != "true" ]; then
    #########################################
    # START Installing Essentials
    #########################################
    KIRA_REPOS=/kira/repos

    KIRA_INFRA="$KIRA_REPOS/kira"
    KIRA_SEKAI="$KIRA_REPOS/sekai"
    KIRA_INTERX="$KIRA_REPOS/interx"

    KIRA_SETUP=/kira/setup
    KIRA_UPDATE=/kira/update
    KIRA_MANAGER="/kira/manager"

    KIRA_SCRIPTS="${KIRA_INFRA}/common/scripts"
    KIRA_WORKSTATION="${KIRA_INFRA}/workstation"

    SEKAID_HOME="/root/.sekaid"

    DOCKER_COMMON="/docker/shared/common"
    # read only common directory
    DOCKER_COMMON_RO="/docker/shared/common_ro"
    GLOBAL_COMMON_RO="/docker/shared/common_ro/kiraglob"

    rm -rfv $KIRA_DUMP
    mkdir -p "$KIRA_DUMP/INFRA/manager" $KIRA_INFRA $KIRA_SEKAI $KIRA_INTERX $KIRA_SETUP $KIRA_MANAGER $DOCKER_COMMON $DOCKER_COMMON_RO $GLOBAL_COMMON_RO

    ESSENTIALS_HASH=$(echo "$CDHELPER_VERSION-$KIRA_HOME-$INFRA_BRANCH-$INFRA_REPO-$ARCHITECTURE-24" | md5sum | awk '{ print $1 }' || echo -n "")
    KIRA_SETUP_ESSSENTIALS="$KIRA_SETUP/essentials-$ESSENTIALS_HASH"
    if [ ! -f "$KIRA_SETUP_ESSSENTIALS" ] ; then
        echo "INFO: Installing Essential Packages & Env Variables..."
        rm -fv /var/lib/apt/lists/lock || echo "WARINING: Failed to remove APT lock"
        apt-get update -y
        apt-get install -y --allow-unauthenticated --allow-downgrades --allow-remove-essential --allow-change-held-packages \
            software-properties-common apt-transport-https ca-certificates gnupg curl wget git build-essential \
            nghttp2 libnghttp2-dev libssl-dev fakeroot dpkg-dev libcurl4-openssl-dev net-tools jq aptitude \
            zip unzip p7zip-full 
        
        apt update -y
        apt install -y bc dnsutils psmisc netcat nmap parallel

        ln -s /usr/bin/git /bin/git || echo "WARNING: Git symlink already exists"
        git config --add --global core.autocrlf input || echo "WARNING: Failed to set global autocrlf"
        git config --unset --global core.filemode || echo "WARNING: Failed to unset global filemode"
        git config --add --global core.filemode false || echo "WARNING: Failed to set global filemode"
        git config --add --global pager.branch false || echo "WARNING: Failed to disable branch pager"
        git config --add --global http.sslVersion "tlsv1.2" || echo "WARNING: Failed to set ssl version"

        echo "INFO: Base Tools Setup..."
        export DOTNET_SYSTEM_GLOBALIZATION_INVARIANT=1
        cd /tmp

        if [[ "${ARCHITECTURE,,}" == *"arm"* ]] || [[ "${ARCHITECTURE,,}" == *"aarch"* ]] ; then
            CDHELPER_ARCH="arm64"
            EXPECTED_HASH="c2e40c7143f4097c59676f037ac6eaec68761d965bd958889299ab32f1bed6b3"
        else
            CDHELPER_ARCH="x64"
            EXPECTED_HASH="082e05210f93036e0008658b6c6bd37ab055bac919865015124a0d72e18a45b7"
        fi

        FILE_HASH=$(sha256sum ./CDHelper-linux-$CDHELPER_ARCH.zip | awk '{ print $1 }' || echo -n "")

        if [ "$FILE_HASH" != "$EXPECTED_HASH" ]; then
            rm -f -v ./CDHelper-linux-$CDHELPER_ARCH.zip
            wget "https://github.com/asmodat/CDHelper/releases/download/$CDHELPER_VERSION/CDHelper-linux-$CDHELPER_ARCH.zip"
            FILE_HASH=$(sha256sum ./CDHelper-linux-$CDHELPER_ARCH.zip | awk '{ print $1 }')

            if [ "$FILE_HASH" != "$EXPECTED_HASH" ]; then
                set +x
                echo -e "\nDANGER: Failed to check integrity hash of the CDHelper tool !!!\nERROR: Expected hash: $EXPECTED_HASH, but got $FILE_HASH\n"
                SELECT="" && while [ "${SELECT,,}" != "x" ] && [ "${SELECT,,}" != "c" ] ; do echo -en "\e[31;1mPress e[X]it or [C]ontinue to disregard the issue\e[0m\c" && read -d'' -s -n1 ACCEPT && echo ""; done
                [ "${SELECT,,}" == "x" ] && exit
                echo "DANGER: You decided to disregard a potential vulnerability !!!"
                echo -en "\e[31;1mPress any key to continue or Ctrl+C to abort...\e[0m" && read -n 1 -s && echo ""
                set -x
            fi
        else
            echo "INFO: CDHelper tool was already downloaded"
        fi

        INSTALL_DIR="/usr/local/bin/CDHelper"
        rm -rfv $INSTALL_DIR
        mkdir -pv $INSTALL_DIR
        unzip CDHelper-linux-$CDHELPER_ARCH.zip -d $INSTALL_DIR
        chmod -R -v 555 $INSTALL_DIR

        ls -l /bin/CDHelper || echo "Symlink not found"
        rm -fv /bin/CDHelper || echo "Removing old symlink"
        ln -s $INSTALL_DIR/CDHelper /bin/CDHelper || echo "CDHelper symlink already exists"

        CDHelper version

        CDHelper text lineswap --insert="DOTNET_SYSTEM_GLOBALIZATION_INVARIANT=1" --prefix="DOTNET_SYSTEM_GLOBALIZATION_INVARIANT=" --path=$ETC_PROFILE --append-if-found-not=True
        CDHelper text lineswap --insert="KIRA_HOME=$KIRA_HOME" --prefix="KIRA_HOME=" --path=$ETC_PROFILE --append-if-found-not=True
        CDHelper text lineswap --insert="KIRA_DUMP=$KIRA_DUMP" --prefix="KIRA_DUMP=" --path=$ETC_PROFILE --append-if-found-not=True
        CDHelper text lineswap --insert="KIRA_SNAP=$KIRA_SNAP" --prefix="KIRA_SNAP=" --path=$ETC_PROFILE --append-if-found-not=True
        CDHelper text lineswap --insert="KIRA_SCAN=$KIRA_HOME/kirascan" --prefix="KIRA_SCAN=" --path=$ETC_PROFILE --append-if-found-not=True
        CDHelper text lineswap --insert="KIRA_SECRETS=$KIRA_SECRETS" --prefix="KIRA_SECRETS=" --path=$ETC_PROFILE --append-if-found-not=True
        CDHelper text lineswap --insert="KIRA_CONFIGS=$KIRA_CONFIGS" --prefix="KIRA_CONFIGS=" --path=$ETC_PROFILE --append-if-found-not=True
        CDHelper text lineswap --insert="PUBLIC_PEERS=$KIRA_CONFIGS/public_peers" --prefix="PUBLIC_PEERS=" --path=$ETC_PROFILE --append-if-found-not=True
        CDHelper text lineswap --insert="PUBLIC_SEEDS=$KIRA_CONFIGS/public_seeds" --prefix="PUBLIC_SEEDS=" --path=$ETC_PROFILE --append-if-found-not=True

        CDHelper text lineswap --insert="KIRA_MANAGER=$KIRA_MANAGER" --prefix="KIRA_MANAGER=" --path=$ETC_PROFILE --append-if-found-not=True
        CDHelper text lineswap --insert="KIRA_REPOS=$KIRA_REPOS" --prefix="KIRA_REPOS=" --path=$ETC_PROFILE --append-if-found-not=True
        CDHelper text lineswap --insert="KIRA_SETUP=$KIRA_SETUP" --prefix="KIRA_SETUP=" --path=$ETC_PROFILE --append-if-found-not=True
        CDHelper text lineswap --insert="KIRA_UPDATE=$KIRA_UPDATE" --prefix="KIRA_UPDATE=" --path=$ETC_PROFILE --append-if-found-not=True

        CDHelper text lineswap --insert="KIRA_INFRA=$KIRA_INFRA" --prefix="KIRA_INFRA=" --path=$ETC_PROFILE --append-if-found-not=True
        CDHelper text lineswap --insert="KIRA_SEKAI=$KIRA_SEKAI" --prefix="KIRA_SEKAI=" --path=$ETC_PROFILE --append-if-found-not=True
        CDHelper text lineswap --insert="KIRA_INTERX=$KIRA_INTERX" --prefix="KIRA_INTERX=" --path=$ETC_PROFILE --append-if-found-not=True

        CDHelper text lineswap --insert="KIRA_SCRIPTS=$KIRA_SCRIPTS" --prefix="KIRA_SCRIPTS=" --path=$ETC_PROFILE --append-if-found-not=True
        CDHelper text lineswap --insert="KIRA_WORKSTATION=$KIRA_WORKSTATION" --prefix="KIRA_WORKSTATION=" --path=$ETC_PROFILE --append-if-found-not=True
        
        CDHelper text lineswap --insert="DOCKER_COMMON=$DOCKER_COMMON" --prefix="DOCKER_COMMON=" --path=$ETC_PROFILE --append-if-found-not=True
        CDHelper text lineswap --insert="DOCKER_COMMON_RO=$DOCKER_COMMON_RO" --prefix="DOCKER_COMMON_RO=" --path=$ETC_PROFILE --append-if-found-not=True
        CDHelper text lineswap --insert="GLOBAL_COMMON_RO=$GLOBAL_COMMON_RO" --prefix="GLOBAL_COMMON_RO=" --path=$ETC_PROFILE --append-if-found-not=True
        CDHelper text lineswap --insert="LOCAL_GENESIS_PATH=$DOCKER_COMMON_RO/genesis.json" --prefix="LOCAL_GENESIS_PATH=" --path=$ETC_PROFILE --append-if-found-not=True

        CDHelper text lineswap --insert="ETC_PROFILE=$ETC_PROFILE" --prefix="ETC_PROFILE=" --path=$ETC_PROFILE --append-if-found-not=True
        CDHelper text lineswap --insert="SEKAID_HOME=$SEKAID_HOME" --prefix="SEKAID_HOME=" --path=$ETC_PROFILE --append-if-found-not=True

        CDHelper text lineswap --insert="DEFAULT_SSH_PORT=$DEFAULT_SSH_PORT" --prefix="DEFAULT_SSH_PORT=" --path=$ETC_PROFILE --append-if-found-not=True
        CDHelper text lineswap --insert="DEFAULT_P2P_PORT=$DEFAULT_P2P_PORT" --prefix="DEFAULT_P2P_PORT=" --path=$ETC_PROFILE --append-if-found-not=True
        CDHelper text lineswap --insert="DEFAULT_RPC_PORT=$DEFAULT_RPC_PORT" --prefix="DEFAULT_RPC_PORT=" --path=$ETC_PROFILE --append-if-found-not=True
        CDHelper text lineswap --insert="DEFAULT_GRPC_PORT=$DEFAULT_GRPC_PORT" --prefix="DEFAULT_GRPC_PORT=" --path=$ETC_PROFILE --append-if-found-not=True
        CDHelper text lineswap --insert="DEFAULT_INTERX_PORT=$DEFAULT_INTERX_PORT" --prefix="DEFAULT_INTERX_PORT=" --path=$ETC_PROFILE --append-if-found-not=True
        CDHelper text lineswap --insert="DEFAULT_PROMETHEUS_PORT=$DEFAULT_PROMETHEUS_PORT" --prefix="DEFAULT_PROMETHEUS_PORT=" --path=$ETC_PROFILE --append-if-found-not=True

        CDHelper text lineswap --insert="KIRA_SEED_PROMETHEUS_PORT=$KIRA_SEED_PROMETHEUS_PORT" --prefix="KIRA_SEED_PROMETHEUS_PORT=" --path=$ETC_PROFILE --append-if-found-not=True
        CDHelper text lineswap --insert="KIRA_SENTRY_PROMETHEUS_PORT=$KIRA_SENTRY_PROMETHEUS_PORT" --prefix="KIRA_SENTRY_PROMETHEUS_PORT=" --path=$ETC_PROFILE --append-if-found-not=True
        CDHelper text lineswap --insert="KIRA_VALIDATOR_PROMETHEUS_PORT=$KIRA_VALIDATOR_PROMETHEUS_PORT" --prefix="KIRA_VALIDATOR_PROMETHEUS_PORT=" --path=$ETC_PROFILE --append-if-found-not=True

        CDHelper text lineswap --insert="KIRA_INTERX_PORT=$KIRA_INTERX_PORT" --prefix="KIRA_INTERX_PORT=" --path=$ETC_PROFILE --append-if-found-not=True
        
        CDHelper text lineswap --insert="KIRA_SENTRY_RPC_PORT=$KIRA_SENTRY_RPC_PORT" --prefix="KIRA_SENTRY_RPC_PORT=" --path=$ETC_PROFILE --append-if-found-not=True
        CDHelper text lineswap --insert="KIRA_SEED_RPC_PORT=$KIRA_SEED_RPC_PORT" --prefix="KIRA_SEED_RPC_PORT=" --path=$ETC_PROFILE --append-if-found-not=True
        CDHelper text lineswap --insert="KIRA_VALIDATOR_RPC_PORT=$KIRA_VALIDATOR_RPC_PORT" --prefix="KIRA_VALIDATOR_RPC_PORT=" --path=$ETC_PROFILE --append-if-found-not=True
        
        CDHelper text lineswap --insert="KIRA_SEED_P2P_PORT=$KIRA_SEED_P2P_PORT" --prefix="KIRA_SEED_P2P_PORT=" --path=$ETC_PROFILE --append-if-found-not=True
        CDHelper text lineswap --insert="KIRA_SENTRY_P2P_PORT=$KIRA_SENTRY_P2P_PORT" --prefix="KIRA_SENTRY_P2P_PORT=" --path=$ETC_PROFILE --append-if-found-not=True
        CDHelper text lineswap --insert="KIRA_VALIDATOR_P2P_PORT=$KIRA_VALIDATOR_P2P_PORT" --prefix="KIRA_VALIDATOR_P2P_PORT=" --path=$ETC_PROFILE --append-if-found-not=True

        CDHelper text lineswap --insert="KIRA_SEED_GRPC_PORT=$KIRA_SEED_GRPC_PORT" --prefix="KIRA_SEED_GRPC_PORT=" --path=$ETC_PROFILE --append-if-found-not=True
        CDHelper text lineswap --insert="KIRA_SENTRY_GRPC_PORT=$KIRA_SENTRY_GRPC_PORT" --prefix="KIRA_SENTRY_GRPC_PORT=" --path=$ETC_PROFILE --append-if-found-not=True
        CDHelper text lineswap --insert="KIRA_VALIDATOR_GRPC_PORT=$KIRA_VALIDATOR_GRPC_PORT" --prefix="KIRA_VALIDATOR_GRPC_PORT=" --path=$ETC_PROFILE --append-if-found-not=True

        CDHelper text lineswap --insert="KIRA_REGISTRY_PORT=$KIRA_REGISTRY_PORT" --prefix="KIRA_REGISTRY_PORT=" --path=$ETC_PROFILE --append-if-found-not=True
        

        touch $KIRA_SETUP_ESSSENTIALS
    else
        echo "INFO: Essentials were already installed: $(git --version), Curl, Wget..."
    fi

    #########################################
    # END Installing Essentials
    #########################################

    echo "INFO: Updating kira Repository..."
    rm -rfv $KIRA_INFRA
    mkdir -p $KIRA_INFRA
    git clone --branch $INFRA_BRANCH $INFRA_REPO $KIRA_INFRA
    cd $KIRA_INFRA
    git describe --all --always
    chmod -R 555 $KIRA_INFRA

    # update old processes
    rm -rfv $KIRA_MANAGER && mkdir -p "$KIRA_MANAGER"
    cp -rfv "$KIRA_WORKSTATION/." $KIRA_MANAGER
    chmod -R 555 $KIRA_MANAGER

    echo "INFO: ReStarting init script to launch setup menu..."
    source $KIRA_MANAGER/init.sh "$INFRA_BRANCH" "True" "$START_TIME_INIT"
    echo "INFO: Init script restart finished."
    exit 0
fi

SETUP_VER=$(cat $KIRA_INFRA/version || echo "")
[ -z "SETUP_VER" ] && echo -en "\e[31;1mERROR: Invalid setup release version!\e[0m" && exit 1

CDHelper text lineswap --insert="KIRA_SETUP_VER=\"$SETUP_VER\"" --prefix="KIRA_SETUP_VER=" --path=$ETC_PROFILE --append-if-found-not=True
CDHelper text lineswap --insert="KIRA_USER=\"$KIRA_USER\"" --prefix="KIRA_USER=" --path=$ETC_PROFILE --append-if-found-not=True

CDHelper text lineswap --insert="INFRA_BRANCH=\"$INFRA_BRANCH\"" --prefix="INFRA_BRANCH=" --path=$ETC_PROFILE --append-if-found-not=True
CDHelper text lineswap --insert="SEKAI_BRANCH=\"$SEKAI_BRANCH\"" --prefix="SEKAI_BRANCH=" --path=$ETC_PROFILE --append-if-found-not=True
CDHelper text lineswap --insert="INTERX_BRANCH=\"$INTERX_BRANCH\"" --prefix="INTERX_BRANCH=" --path=$ETC_PROFILE --append-if-found-not=True

CDHelper text lineswap --insert="INFRA_REPO=\"$INFRA_REPO\"" --prefix="INFRA_REPO=" --path=$ETC_PROFILE --append-if-found-not=True
CDHelper text lineswap --insert="SEKAI_REPO=\"$SEKAI_REPO\"" --prefix="SEKAI_REPO=" --path=$ETC_PROFILE --append-if-found-not=True
CDHelper text lineswap --insert="INTERX_REPO=\"$INTERX_REPO\"" --prefix="INTERX_REPO=" --path=$ETC_PROFILE --append-if-found-not=True

CDHelper text lineswap --insert="source \$KIRA_MANAGER/utils.sh" --suffix="/utils.sh" --path=$ETC_PROFILE --append-if-found-not=True

echo "INFO: Startting cleanup..."
apt-get autoclean || echo "WARNING: autoclean failed"
apt-get clean || echo "WARNING: clean failed"
apt-get autoremove || echo "WARNING: autoremove failed"
journalctl --vacuum-time=3d || echo "WARNING: journalctl vacuum failed"

$KIRA_MANAGER/setup/tools.sh

set +x
echo "INFO: Your host environment was initialized"
echo -e "\e[33;1mTERMS & CONDITIONS: Make absolutely sure that you are NOT running this script on your primary PC operating system, it can cause irreversible data loss and change of firewall rules which might make your system vulnerable to various security threats or entirely lock you out of the system. By proceeding you take full responsibility for your own actions and accept that you continue on your own risk. You also acknowledge that malfunction of any software you run might potentially cause irreversible loss of assets due to unforeseen issues and circumstances including but not limited to hardware and/or software faults and/or vulnerabilities.\e[0m"
echo -en "\e[31;1mPress any key to accept terms & continue or Ctrl+C to abort...\e[0m" && read -n 1 -s && echo ""
echo "INFO: Launching setup menu..."
set -x
source $KIRA_MANAGER/menu.sh "true"

set +x
echo "------------------------------------------------"
echo "| FINISHED: INIT                               |"
echo "|  ELAPSED: $(($(date -u +%s) - $START_TIME_INIT)) seconds"
echo "------------------------------------------------"
set -x
exit 0
