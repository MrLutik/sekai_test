#!/bin/bash
set +e && source "/etc/profile" &>/dev/null && set -e
source $KIRA_MANAGER/utils.sh
# quick edit: FILE="$KIRA_MANAGER/containers/start-snapshot.sh" && rm $FILE && nano $FILE && chmod 555 $FILE

set -x

MAX_HEIGHT=$1
SYNC_FROM_SNAP=$2

[ -z "$MAX_HEIGHT" ] && MAX_HEIGHT=$(globGet LATEST_BLOCK) && (! $(isNaturalNumber "$MAX_HEIGHT")) && MAX_HEIGHT=0
# ensure to create parent directory for shared status info
CONTAINER_NAME="snapshot"
SNAP_STATUS="$KIRA_SNAP/status"
COMMON_PATH="$DOCKER_COMMON/$CONTAINER_NAME"
COMMON_LOGS="$COMMON_PATH/logs"
HALT_FILE="$COMMON_PATH/halt"
SNAP_DONE="$SNAP_STATUS/done"
SNAP_PROGRESS="$SNAP_STATUS/progress"

CPU_CORES=$(cat /proc/cpuinfo | grep processor | wc -l || echo "0")
RAM_MEMORY=$(grep MemTotal /proc/meminfo | awk '{print $2}' || echo "0")
[ "${DEPLOYMENT_MODE,,}" == "minimal" ] && UTIL_DIV=4 || UTIL_DIV=6
CPU_RESERVED=$(echo "scale=2; ( $CPU_CORES / $UTIL_DIV )" | bc)
RAM_RESERVED="$(echo "scale=0; ( $RAM_MEMORY / $UTIL_DIV ) / 1024 " | bc)m"

SNAP_FILENAME="${NETWORK_NAME}-$MAX_HEIGHT-$(date -u +%s).zip"
SNAP_FILE="$KIRA_SNAP/$SNAP_FILENAME"

set +x
echoWarn "------------------------------------------------"
echoWarn "| STARTING $CONTAINER_NAME NODE"
echoWarn "|-----------------------------------------------"
echoWarn "|    HOSTNAME: $KIRA_SNAPSHOT_DNS"
echoWarn "| SYNC HEIGHT: $MAX_HEIGHT"
echoWarn "|  SNAP DEST.: $SNAP_FILE"
echoWarn "| SNAP SOURCE: $SYNC_FROM_SNAP"
echoWarn "|     MAX CPU: $CPU_RESERVED / $CPU_CORES"
echoWarn "|     MAX RAM: $RAM_RESERVED"
echoWarn "------------------------------------------------"
set -x

echoInfo "INFO: Wiping '$CONTAINER_NAME' resources..."
$KIRA_SCRIPTS/container-delete.sh "$CONTAINER_NAME"

chattr -iR $COMMON_PATH || echoWarn "WARNING: Failed to remove integrity protection from $COMMON_PATH"
rm -rfv "$COMMON_PATH" "$SNAP_STATUS" "$SNAP_DONE" "$SNAP_PROGRESS"
mkdir -p "$COMMON_LOGS" "$SNAP_STATUS"

echoInfo "INFO: Loading secrets..."
set +x
set +e
source $KIRAMGR_SCRIPTS/load-secrets.sh
set -x
set -e

echoInfo "INFO: Checking peers info..."
SEED_SEED=$(echo "${SEED_NODE_ID}@$KIRA_SEED_DNS:$DEFAULT_P2P_PORT" | xargs | tr -d '\n' | tr -d '\r')
SENTRY_SEED=$(echo "${SENTRY_NODE_ID}@$KIRA_SENTRY_DNS:$DEFAULT_P2P_PORT" | xargs | tr -d '\n' | tr -d '\r')
VALIDATOR_SEED=$(echo "${VALIDATOR_NODE_ID}@$KIRA_VALIDATOR_DNS:$DEFAULT_P2P_PORT" | xargs | tr -d '\n' | tr -d '\r')
PRIV_SENTRY_SEED=$(echo "${PRIV_SENTRY_NODE_ID}@$KIRA_PRIV_SENTRY_DNS:$DEFAULT_P2P_PORT" | xargs | tr -d '\n' | tr -d '\r')

if [ "${DEPLOYMENT_MODE,,}" == "minimal" ] && [ "${INFRA_MODE,,}" == "validator" ] ; then
    CFG_persistent_peers="tcp://$VALIDATOR_SEED"
    CONTAINER_NETWORK="$KIRA_VALIDATOR_NETWORK"
else
    CFG_persistent_peers="tcp://$SENTRY_SEED"
    CONTAINER_NETWORK="$KIRA_SENTRY_NETWORK"
fi

cp -afv $KIRA_SECRETS/snapshot_node_key.json $COMMON_PATH/node_key.json

SNAP_DESTINATION="$COMMON_PATH/snap.zip"
rm -rfv $SNAP_DESTINATION
if [ -f "$KIRA_SNAP_PATH" ] ; then
    echoInfo "INFO: State snapshot was found, cloning..."
    ln -fv "$KIRA_SNAP_PATH" "$SNAP_DESTINATION"
fi

echoInfo "INFO: Starting '$CONTAINER_NAME' container..."
docker run -d \
    --cpus="$CPU_RESERVED" \
    --memory="$RAM_RESERVED" \
    --oom-kill-disable \
    -p $KIRA_SNAPSHOT_RPC_PORT:$DEFAULT_RPC_PORT \
    --hostname $KIRA_SNAPSHOT_DNS \
    --restart=always \
    --name $CONTAINER_NAME \
    --net=$CONTAINER_NETWORK \
    --log-opt max-size=5m \
    --log-opt max-file=5 \
    -e HALT_HEIGHT="$MAX_HEIGHT" \
    -e SNAP_FILENAME="$SNAP_FILENAME" \
    -e NETWORK_NAME="$NETWORK_NAME" \
    -e HOSTNAME="$KIRA_SNAPSHOT_DNS" \
    -e CONTAINER_NETWORK="$CONTAINER_NETWORK" \
    -e CFG_moniker="KIRA ${CONTAINER_NAME^^} NODE" \
    -e CFG_seeds="$CFG_seeds" \
    -e CFG_persistent_peers="$CFG_persistent_peers" \
    -e CFG_grpc_laddr="tcp://0.0.0.0:$DEFAULT_GRPC_PORT" \
    -e CFG_rpc_laddr="tcp://0.0.0.0:$DEFAULT_RPC_PORT" \
    -e CFG_p2p_laddr="tcp://0.0.0.0:$DEFAULT_P2P_PORT" \
    -e CFG_private_peer_ids="" \
    -e CFG_unconditional_peer_ids="$VALIDATOR_NODE_ID,$PRIV_SENTRY_NODE_ID,$SEED_NODE_ID,$SENTRY_NODE_ID" \
    -e CFG_pex="false" \
    -e CFG_addr_book_strict="false" \
    -e CFG_seed_mode="false" \
    -e CFG_max_num_outbound_peers="0" \
    -e CFG_max_num_inbound_peers="0" \
    -e CFG_handshake_timeout="60s" \
    -e CFG_allow_duplicate_ip="true" \
    -e CFG_dial_timeout="30s" \
    -e CFG_max_txs_bytes="131072000" \
    -e CFG_max_tx_bytes="131072" \
    -e CFG_send_rate="65536000" \
    -e CFG_recv_rate="65536000" \
    -e CFG_max_packet_msg_payload_size="131072" \
    -e MIN_HEIGHT="$(globGet MIN_HEIGHT)" \
    -e NEW_NETWORK="$NEW_NETWORK" \
    -e EXTERNAL_SYNC="$EXTERNAL_SYNC" \
    -e KIRA_SETUP_VER="$KIRA_SETUP_VER" \
    -e INTERNAL_P2P_PORT="$DEFAULT_P2P_PORT" \
    -e INTERNAL_RPC_PORT="$DEFAULT_RPC_PORT" \
    -e NODE_TYPE=$CONTAINER_NAME \
    -e NODE_ID="$SNAPSHOT_NODE_ID" \
    -e DEPLOYMENT_MODE="$DEPLOYMENT_MODE" \
    -v $COMMON_PATH:/common \
    -v $KIRA_SNAP:/snap \
    -v $DOCKER_COMMON_RO:/common_ro:ro \
    kira:latest # use sentry image as base

echoInfo "INFO: Waiting for $CONTAINER_NAME node to start..."
CONTAINER_CREATED="true" && $KIRAMGR_SCRIPTS/await-sentry-init.sh "$CONTAINER_NAME" "$SNAPSHOT_NODE_ID" || CONTAINER_CREATED="false"

if [ "${CONTAINER_CREATED,,}" != "true" ]; then
    echoErr "ERROR: Snapshot failed, '$CONTAINER_NAME' container did not start"
    $KIRA_SCRIPTS/container-pause.sh $CONTAINER_NAME || echoErr "ERROR: Failed to pause container"
else
    echoInfo "INFO: Success '$CONTAINER_NAME' container was started"
    rm -fv "$SNAP_DESTINATION"

    echoInfo "INFO: Checking genesis SHA256 hash"
    TEST_SHA256=$(docker exec -i "$CONTAINER_NAME" /bin/bash -c ". /etc/profile;sha256 \$SEKAID_HOME/config/genesis.json" || echo -n "")
    if [ -z "$TEST_SHA256" ] || [ "$TEST_SHA256" != "$GENESIS_SHA256" ]; then
        echoErr "ERROR: Snapshot failed, expected genesis checksum to be '$GENESIS_SHA256' but got '$TEST_SHA256'"
        $KIRA_SCRIPTS/container-pause.sh $CONTAINER_NAME || echoErr "ERROR: Failed to pause container"
        exit 1
    fi

    echoInfo "INFO: Snapshot destination: $SNAP_FILE"
    echoInfo "INFO: Work in progress, await snapshot container to reach 100% sync status"
fi