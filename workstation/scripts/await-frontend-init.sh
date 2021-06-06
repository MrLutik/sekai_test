#!/bin/bash
set +e && source "/etc/profile" &>/dev/null && set -e
source $KIRA_MANAGER/utils.sh
set -x

CONTAINER_NAME="frontend"
COMMON_PATH="$DOCKER_COMMON/$CONTAINER_NAME"
COMMON_LOGS="$COMMON_PATH/logs"
TIMER_NAME="${CONTAINER_NAME^^}_INIT"
TIMEOUT=1800

set +x
echoWarn "--------------------------------------------------"
echoWarn "|  STARTING ${CONTAINER_NAME^^} INIT $KIRA_SETUP_VER"
echoWarn "|-------------------------------------------------"
echoWarn "| COMMON DIR: $COMMON_PATH"
echoWarn "|    TIMEOUT: $TIMEOUT seconds"
echoWarn "|-------------------------------------------------"
set -x

globDel "${CONTAINER_NAME}_STATUS" "${CONTAINER_NAME}_EXISTS"
timerStart $TIMER_NAME

systemctl restart kirascan || echoWarn "WARNING: Could NOT restart kira scan service"

while [[ $(timerSpan $TIMER_NAME) -lt $TIMEOUT ]] ; do
    echoInfo "INFO: Waiting for container $CONTAINER_NAME to start..."
    if [ "$(globGet ${CONTAINER_NAME}_EXISTS)" != "true" ]; then
        echoWarn "WARNING: $CONTAINER_NAME container does not exists yet, waiting up to $(timerSpan $TIMER_NAME $TIMEOUT) seconds ..." && sleep 30 && continue
    else echoInfo "INFO: Success, container $CONTAINER_NAME was found" ; fi

    echoInfo "INFO: Awaiting $CONTAINER_NAME initialization..."
    if [ "$(globGet ${CONTAINER_NAME}_STATUS)" != "running" ] ; then
        echoWarn "WARNING: $CONTAINER_NAME is not initialized yet, waiting up to $(timerSpan $TIMER_NAME $TIMEOUT) seconds ..." && sleep 30 && continue
    else echoInfo "INFO: Success, $CONTAINER_NAME was initialized" && break ; fi
done

echoInfo "INFO: Printing all $CONTAINER_NAME health logs..."
docker inspect --format "{{json .State.Health }}" $($KIRA_SCRIPTS/container-id.sh "$CONTAINER_NAME") | jq '.Log[-1].Output' | xargs | sed 's/\\n/\n/g' || echo "INFO: Failed to display $CONTAINER_NAME container health logs"

echoInfo "INFO: Printing all $CONTAINER_NAME start logs..."
cat $COMMON_LOGS/start.log | tail -n 75 || echoWarn "WARNING: Failed to display $CONTAINER_NAME container start logs"

[ "$(globGet ${CONTAINER_NAME}_STATUS)" != "running" ] && echoErr "ERROR: $CONTAINER_NAME was not started sucessfully within defined time" && exit 1

set +x
echoWarn "------------------------------------------------"
echoWarn "| FINISHED: ${CONTAINER_NAME^^} INIT"
echoWarn "|  ELAPSED: $(timerSpan $TIMER_NAME) seconds"
echoWarn "------------------------------------------------"
set -x