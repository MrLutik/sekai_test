#!/bin/bash
set +e && source $ETC_PROFILE &>/dev/null && set -e
source $SELF_SCRIPTS/utils.sh
set -x

LATEST_BLOCK_HEIGHT=$1
PREVIOUS_HEIGHT=$2
HEIGHT=$3
CATCHING_UP=$4
CONSENSUS_STOPPED=$5

LIP_FILE="$COMMON_READ/local_ip"
PIP_FILE="$COMMON_READ/public_ip"

if [[ $PREVIOUS_HEIGHT -ge $HEIGHT ]]; then
    set +x
    echoWarn "WARNING: Blocks are not beeing produced or synced"
    echoWarn "WARNING: Current height: $HEIGHT"
    echoWarn "WARNING: Previous height: $PREVIOUS_HEIGHT"
    echoWarn "WARNING: Latest height: $LATEST_BLOCK_HEIGHT"
    echoWarn "WARNING: Consensus Stopped: $CONSENSUS_STOPPED"
  
      if [[ $LATEST_BLOCK_HEIGHT -ge 1 ]] && [[ $LATEST_BLOCK_HEIGHT -le $HEIGHT ]] && [ "$CONSENSUS_STOPPED" == "true" ] ; then
          echoWarn "WARNINIG: Cosnensus halted, lack of block production is not result of the issue with the node"
      else
          echoErr "ERROR: Block production stopped"
          sleep 10
          exit 1
      fi
else
    echoInfo "INFO, Success, new blocks were created or synced: $HEIGHT"
fi

echoInfo "INFO: Latest Block Height: $HEIGHT"
set -x

if [ ! -z "$EXTERNAL_ADDR" ] ; then
    echoInfo "INFO: Checking availability of the external address '$EXTERNAL_ADDR'"
    if timeout 15 nc -z $EXTERNAL_ADDR $EXTERNAL_P2P_PORT ; then 
        echoInfo "INFO: Success, your node external address '$EXTERNAL_ADDR' is exposed"
        echo "ONLINE" > "$COMMON_DIR/external_address_status"
    else
        echoErr "ERROR: Your node external address is NOT visible to other nodes"
        echo "OFFLINE" > "$COMMON_DIR/external_address_status"
    fi
else
    echoWarn "WARNING: This node is NOT advertising its it's public or local external address to other nodes in the network!"
    echo "OFFLINE" > "$COMMON_DIR/external_address_status"
fi

exit 0