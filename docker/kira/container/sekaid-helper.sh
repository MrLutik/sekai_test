#!/bin/bash

# QUICK EDIT: FILE="$SELF_CONTAINER/sekaid-helper.sh" && rm $FILE && nano $FILE && chmod 555 $FILE

function txQuery() {
    (! $(isTxHash "$1")) && echoErr "ERROR: Infalid Transaction Hash '$1'" && sekaid query tx "$1" --output=json --home=$SEKAID_HOME | jq || echoErr "ERROR: Transaction '$1' was NOT found or failed"
}

function txAwait() {
    local START_TIME="$(date -u +%s)"
    local RAW=""
    local TIMEOUT=""
    
    if (! $(isTxHash "$1")) ; then
        RAW=$(cat)
        TIMEOUT=$1
    else
        RAW=$1
        TIMEOUT=$2
    fi

    (! $(isNaturalNumber $TIMEOUT)) && TIMEOUT=0
    [[ $TIMEOUT -le 0 ]] && MAX_TIME="∞" || MAX_TIME="$TIMEOUT"
    
    local TXHASH=""
    if (! $(isTxHash "$RAW")) ; then
        # INPUT example: {"height":"0","txhash":"DF8BFCC9730FDBD33AEA184EC3D6C37B4311BC1C0E2296893BC020E4638A0D6F","codespace":"","code":0,"data":"","raw_log":"","logs":[],"info":"","gas_wanted":"0","gas_used":"0","tx":null,"timestamp":""}
        local VAL=$(echo $RAW | jsonParse "" 2> /dev/null || echo "")
        if ($(isNullOrEmpty "$VAL")) ; then
            echoErr "ERROR: Failed to propagate transaction:"
            echoErr "$RAW"
            return 1
        fi

        TXHASH=$(echo $VAL | jsonQuickParse "txhash" 2> /dev/null || echo "")
        if ($(isNullOrEmpty "$TXHASH")) ; then
            echoErr "ERROR: Transaction hash 'txhash' was NOT found in the tx propagation response:"
            echoErr "$RAW"
            return 1
        fi
    else
        TXHASH="${RAW^^}"
    fi

    echoInfo "INFO: Transaction hash '$TXHASH' was found!"
    echoInfo "INFO: Please wait for tx confirmation, timeout will occur in $MAX_TIME seconds ..."

    while : ; do
        local ELAPSED=$(($(date -u +%s) - $START_TIME))
        local OUT=$(sekaid query tx $TXHASH --output=json --home=$SEKAID_HOME 2> /dev/null | jsonParse "" 2> /dev/null || echo -n "")
        if [ ! -z "$OUT" ] ; then
            echoInfo "INFO: Transaction query response received received:"
            echo $OUT | jq

            local CODE=$(echo $OUT | jsonQuickParse "code" 2> /dev/null || echo -n "")
            if [ "$CODE" == "0" ] ; then
                echoInfo "INFO: Transaction was confirmed sucessfully!"
                return 0
            else
                echoErr "ERROR: Transaction failed with exit code '$CODE'"
                return 1
            fi
        else
            echoWarn "WAITING: Transaction is NOT confirmed yet, elapsed ${ELAPSED}/${MAX_TIME} s"
        fi

        if [[ $TIMEOUT -gt 0 ]] && [[ $ELAPSED -gt $TIMEOUT ]] ; then
            echoInfo "INFO: Transaction query response was NOT received:"
            echo $RAW | jq 2> /dev/null || echoErr "$RAW"
            echoErr "ERROR: Timeout, failed to confirm tx hash '$TXHASH' within ${TIMEOUT} s limit"
            return 1
        else
            sleep 5
        fi
    done
}

# e.g. showAddress validator
function showAddress() {
    ($(isKiraAddress "$1")) && echo "$1" || echo $(sekaid keys show "$1" --keyring-backend=test --output=json --home=$SEKAID_HOME 2> /dev/null | jsonParse "address" 2> /dev/null || echo -n "")
}

# tryGetPermissions validator
function tryGetPermissions() {
    local ADDR=$(showAddress $1)
    echo $(sekaid query customgov permissions "$ADDR" --output=json --home=$SEKAID_HOME 2> /dev/null | jsonParse 2> /dev/null || echo -n "") && echo -n ""
}

function isPermBlacklisted() {
    local ADDR=$(showAddress "$1")
    local PERM=$2
    if (! $(isNaturalNumber $PERM)) || ($(isNullOrEmpty $ADDR)) ; then
        echo "false"
    else
        INDEX=$(tryGetPermissions $ADDR 2> /dev/null | jq ".blacklist | index($PERM)" 2> /dev/null || echo -n "")
        ($(isNaturalNumber $INDEX)) && echo "true" || echo "false"
    fi
}

function isPermWhitelisted() {
    local ADDR=$(showAddress "$1")
    local PERM=$2
    if (! $(isNaturalNumber $PERM)) || ($(isNullOrEmpty $ADDR)) ; then
        echo "false"
    else
        INDEX=$(tryGetPermissions $ADDR 2> /dev/null | jq ".whitelist | index($PERM)" 2> /dev/null || echo -n "")
        if ($(isNaturalNumber $INDEX)) && (! $(isPermBlacklisted $ADDR $PERM)) ; then
            echo "true" 
        else
            echo "false"
        fi
    fi
}

# e.g. tryGetValidator kiraXXXXXXXXXXX
# e.g. tryGetValidator kiravaloperXXXXXXXXXXX
function tryGetValidator() {
    local VAL_ADDR="${1,,}"
    if [[ $VAL_ADDR == kiravaloper* ]] ; then
        VAL_STATUS=$(sekaid query customstaking validator --val-addr="$VAL_ADDR" --output=json --home=$SEKAID_HOME 2> /dev/null | jsonParse 2> /dev/null || echo -n "")
    elif [[ $VAL_ADDR == kira* ]] ; then
        VAL_STATUS=$(sekaid query customstaking validator --addr="$VAL_ADDR" --output=json --home=$SEKAID_HOME 2> /dev/null | jsonParse 2> /dev/null || echo -n "") 
    else
        VAL_STATUS=""
    fi
    echo $VAL_STATUS
}

function lastProposal() {
    local BOTTOM_PROPOSALS=$(sekaid query customgov proposals --limit=1 --reverse --output=json --home=$SEKAID_HOME | jq -cr '.proposals | last | .proposal_id' 2> /dev/null || echo "")
    local TOP_PROPOSALS=$(sekaid query customgov proposals --limit=1 --output=json --home=$SEKAID_HOME | jq -cr '.proposals | last | .proposal_id' 2> /dev/null || echo "")
    ($(isNullOrEmpty "$BOTTOM_PROPOSALS")) && ($(isNullOrEmpty "$TOP_PROPOSALS")) && echo 0 && return 1
    (! $(isNaturalNumber $BOTTOM_PROPOSALS)) && (! $(isNaturalNumber $TOP_PROPOSALS)) && echo 0 && return 2
    [[ $TOP_PROPOSALS -le 0 ]] && [[ $BOTTOM_PROPOSALS -le 0 ]] && echo 0 && return 3
    [[ $TOP_PROPOSALS -gt $BOTTOM_PROPOSALS ]] && echo "$TOP_PROPOSALS" || echo "$BOTTOM_PROPOSALS"
    return 0
}

# voteYes $(lastProposal) validator
function voteYes() {
    local PROPOSAL=$1
    local ACCOUNT=$2
    echoInfo "INFO: Voting YES on proposal $PROPOSAL with account $ACCOUNT"
    sekaid tx customgov proposal vote $PROPOSAL 1 --from=$ACCOUNT --chain-id=$NETWORK_NAME --keyring-backend=test  --fees=100ukex --yes --log_format=json --broadcast-mode=async --output=json | txAwait
}

# voteNo $(lastProposal) validator
function voteNo() {
    local PROPOSAL=$1
    local ACCOUNT=$2
    echoInfo "INFO: Voting YES on proposal $PROPOSAL with account $ACCOUNT"
    sekaid tx customgov proposal vote $PROPOSAL 0 --from=$ACCOUNT --chain-id=$NETWORK_NAME --keyring-backend=test  --fees=100ukex --yes --log_format=json --broadcast-mode=async --output=json | txAwait
}

function networkProperties() {
    local NETWORK_PROPERTIES=$(sekaid query customgov network-properties --output=json --home=$SEKAID_HOME 2> /dev/null || echo "" | jq -rc 2> /dev/null || echo "")
    ($(isNullOrEmpty "$NETWORK_PROPERTIES")) && echo -n "" && return 1
    echo $NETWORK_PROPERTIES
    return 0
}

# showVotes $(lastProposal) 
function showVotes() {
    sekaid query customgov votes "$1" --output=json --home=$SEKAID_HOME | jsonParse
}

# showProposal $(lastProposal) 
function showProposal() {
    sekaid query customgov proposal "$1" --output json --home=$SEKAID_HOME | jsonParse
}

function showProposals() {
    sekaid query customgov proposals --limit=999999999 --output=json --home=$SEKAID_HOME | jsonParse
}

# propAwait $(lastProposal) 
function propAwait() {
    local START_TIME="$(date -u +%s)"
    local ID=""
    local STATUS=""
    local TIMEOUT=""
    if (! $(isNaturalNumber "$1")) ; then
        ID=$(cat) && STATUS=$1 && TIMEOUT=$2
    else
        ID=$1 && STATUS=$2 && TIMEOUT=$3
    fi
    
    local PROP=$(showProposal $ID 2> /dev/null || echo -n "")
    local RESULT=""
    
    if ($(isNullOrEmpty "$PROP")) ; then
        echoErr "ERROR: Proposal $ID was NOT found"
    else
        echoInfo "INFO: Waiting for proposal $ID to be finalized"
        (! $(isNaturalNumber $TIMEOUT)) && TIMEOUT=0
        [[ $TIMEOUT -le 0 ]] && MAX_TIME="∞" || MAX_TIME="$TIMEOUT"
        while : ; do
            local ELAPSED=$(($(date -u +%s) - $START_TIME))
            RESULT=$(showProposal $ID 2> /dev/null | jq ".result" 2> /dev/null | xargs 2> /dev/null || echo -n "")
            ($(isNullOrEmpty "$STATUS")) && ( [ "${RESULT,,}" == "vote_pending" ] || [ "${RESULT,,}" == "vote_result_enactment" ] ) && break
            (! $(isNullOrEmpty "$STATUS")) && [ "${RESULT,,}" == "${STATUS,,}" ] && break
            if [[ $TIMEOUT -gt 0 ]] && [[ $ELAPSED -gt $TIMEOUT ]] ; then
                echoErr "ERROR: Timeout, failed to finalize proposal '$ID' within ${TIMEOUT} s limit"
                return 1
            else
                sleep 1
            fi
        done
        echoInfo "INFO: Proposal was finalized ($RESULT)"
    fi
}

# e.g. whitelistValidator validator kiraXXXXXXXXXXX
function whitelistValidator() {
    local ACC="$1"
    local ADDR=$(showAddress $2)
    local TIMEOUT=$3
    ($(isNullOrEmpty $ACC)) && echoErr "ERROR: Account name was not defined " && return 1
    ($(isNullOrEmpty $ADDR)) && echoErr "ERROR: Validator address was not defined " && return 1
    (! $(isNaturalNumber $TIMEOUT)) && TIMEOUT=180

    if ($(isPermWhitelisted $ADDR $PermClaimValidator)) ; then
        echoWarn "WARNING: Address $ADDR was already whitelisted as validator"
    else
        echoInfo "INFO: Adding $ADDR to the validator set"
        echoInfo "INFO: Fueling address $ADDR with funds from $ACC"
        sekaid tx bank send $ACC $ADDR "954321ukex" --keyring-backend=test --chain-id=$NETWORK_NAME --fees 100ukex --yes --log_format=json --broadcast-mode=async --output=json | txAwait $TIMEOUT

        echoInfo "INFO: Assigning PermClaimValidator ($PermClaimValidator) permission"
        sekaid tx customgov proposal assign-permission $PermClaimValidator --addr=$ADDR --from=$ACC --keyring-backend=test --chain-id=$NETWORK_NAME --title="Adding Testnet Validator $ADDR" --description="Adding Validator via KIRA Manager" --fees=100ukex --yes --log_format=json --broadcast-mode=async --output=json | txAwait $TIMEOUT

        echoInfo "INFO: Searching for the last proposal submitted on-chain and voting YES"
        LAST_PROPOSAL=$(lastProposal) 
        voteYes $LAST_PROPOSAL validator

        echoInfo "INFO: Showing proposal $LAST_PROPOSAL votes"
        showVotes $LAST_PROPOSAL | jq

        echoInfo "INFO: Showing proposal $LAST_PROPOSAL status"
        showProposal $LAST_PROPOSAL | jq

        echoErr "Date Time Now: $(date '+%Y-%m-%dT%H:%M:%S')"
        echoInfo "INFO: Validator $ADDR will be added to the set after proposal $LAST_PROPOSAL passes"
        #proposalAwait $(lastProposal)
    fi
}

# whitelistValidators <account> <file-name>
# e.g.: whitelistValidators validator ./whitelist
function whitelistValidators() {
    local ACCOUNT=$1
    local WHITELIST=$2
    local TIMEOUT=$3
    (! $(isNaturalNumber $TIMEOUT)) && TIMEOUT=180
    if [ -f "$WHITELIST" ] ; then 
        echoInfo "INFO: List of validators was found ($WHITELIST)"
        while read key ; do
            key=$(echo "$key" | xargs || echo -n "")
            if ($(isNullOrEmpty "$key")) ; then
                echoWarn "INFO: Invalid key $key"
                continue
            fi
            echoInfo "INFO: Whitelisting '$key' using account '$ACCOUNT'"
            whitelistValidator validator $key $TIMEOUT || echoErr "ERROR: Failed to whitelist $key within ${TIMEOUT}s"
        done < $WHITELIST
    else
        echoErr "ERROR: List of validators was NOT found ($WHITELIST)"
    fi
}

# claimValidatorSeat <account> <moniker> <timeout-seconds>
# e.g.: claimValidatorSeat validator "BOB's NODE" 180
function claimValidatorSeat() {
    local ACCOUNT=$1
    local MONIKER=$2
    local TIMEOUT=$3
    ($(isNullOrEmpty $ACCOUNT)) && echoInfo "INFO: Account name was not defined " && return 1
    ($(isNullOrEmpty $MONIKER)) && MONIKER=$(openssl rand -hex 16)
    (! $(isNaturalNumber $TIMEOUT)) && TIMEOUT=180
    sekaid tx customstaking claim-validator-seat --from "$ACCOUNT" --keyring-backend=test --home=$SEKAID_HOME --moniker="$MONIKER" --chain-id=$NETWORK_NAME --broadcast-mode=async --fees=100ukex --yes --output=json | txAwait $TIMEOUT
}

# e.g. showBalance validator
function showBalance() {
    local ADDR=$(showAddress $1)
    local DENOM="$2"
    local RESULT=""
    (! $(isNullOrEmpty $ADDR)) && RESULT=$(sekaid query bank balances "$ADDR" --output=json --home=$SEKAID_HOME 2> /dev/null | jsonParse 2> /dev/null || echo -n "")
    if (! $(isNullOrEmpty $DENOM)) ; then
        RESULT=$(echo $RESULT | showBalance $ADDR | jq '.balances' |  jq ".[] | select(.denom==\"$DENOM\")" | jq '.amount' | xargs 2> /dev/null || echo "0")
        (! $(isNaturalNumber $RESULT)) && RESULT=0
    fi
    echo $RESULT
}

# e.g. sendTokens faucet kiraXXX...XXX 1000 ukex
function sendTokens() {
    local SOURCE=$1
    local DESTINATION=$(showAddress $2)
    local AMOUNT="$3"
    local DENOM="$4"
    echoInfo "INFO: Sending $AMOUNT $DENOM | $SOURCE -> $DESTINATION"
    OLD_BALANCE=$(showBalance "$DESTINATION" "$DENOM") && (! $(isNaturalNumber $OLD_BALANCE)) && OLD_BALANCE=0
    sekaid tx bank send $SOURCE $DESTINATION "${AMOUNT}${DENOM}" --keyring-backend=test --chain-id=$NETWORK_NAME --fees 100ukex --output=json --yes | txAwait $TIMEOUT
    NEW_BALANCE=$(showBalance "$DESTINATION" "$DENOM") && (! $(isNaturalNumber $NEW_BALANCE)) && NEW_BALANCE=0
    echoInfo "INFO: Balance change $DESTINATION | $OLD_BALANCE $DENOM -> $NEW_BALANCE $DENOM"
}

# e.g. showStatus -> { ... }
function showStatus() {
    echo $(sekaid status 2>&1 | jsonParse "" 2>/dev/null || echo -n "")
}

# e.g. showBlockHeight -> 123
function showBlockHeight() {
    SH_LATEST_BLOCK_HEIGHT=$(showStatus | jsonParse "SyncInfo.latest_block_height" 2>/dev/null || echo -n "")
    ($(isNaturalNumber "$SH_LATEST_BLOCK_HEIGHT")) && echo $SH_LATEST_BLOCK_HEIGHT || echo ""
}

# awaitBlocks <number-of-blocks> <timeout-seconds>
# e.g. awaitBlocks 5
function awaitBlocks() {
    local BLOCKS=$1
    (! $(isNaturalNumber $BLOCKS)) && echoErr "ERROR: Number of blocks to await was NOT defined" && return 1
    local SH_START_BLOCK=""
    while : ; do
        local SH_NEW_BLOCK=$(showBlockHeight)
        (! $(isNaturalNumber $SH_NEW_BLOCK)) && sleep 1 && continue
        (! $(isNaturalNumber "$SH_START_BLOCK")) && SH_START_BLOCK=$SH_NEW_BLOCK
        local SH_DELTA=$(($SH_NEW_BLOCK - $SH_START_BLOCK))
        [ $SH_DELTA -gt $BLOCKS ] && break
        sleep 1
    done
}

# e.g. showCatchingUp -> false
function showCatchingUp() {
    local SH_CATCHING_UP=$(showStatus | jsonParse "SyncInfo.catching_up" 2>/dev/null || echo -n "")
    ($(isBoolean "$SH_CATCHING_UP")) && echo "${SH_CATCHING_UP,,}" || echo ""
}

# activateValidator <account> <timeout-seconds>
# e.g. activateValidator validator 180
function activateValidator() {
    local ACCOUNT=$1
    local TIMEOUT=$2
    (! $(isNaturalNumber $TIMEOUT)) && TIMEOUT=180
    ($(isNullOrEmpty $ACCOUNT)) && echoInfo "INFO: Account name was not defined " && return 1
    sekaid tx customslashing activate --from "$ACCOUNT" --chain-id=$NETWORK_NAME --keyring-backend=test --home=$SEKAID_HOME --fees 1000ukex --yes --broadcast-mode=async --log_format=json --output=json | txAwait $TIMEOUT
}

# pauseValidator <account> <timeout-seconds>
# e.g. pauseValidator validator 180
function pauseValidator() {
    local ACCOUNT=$1
    local TIMEOUT=$2
    (! $(isNaturalNumber $TIMEOUT)) && TIMEOUT=180
    ($(isNullOrEmpty $ACCOUNT)) && echoInfo "INFO: Account name was not defined " && return 1
    sekaid tx customslashing pause --from "$ACCOUNT" --chain-id=$NETWORK_NAME --keyring-backend=test --home=$SEKAID_HOME --fees 100ukex --yes --broadcast-mode=async --log_format=json --output=json | txAwait $TIMEOUT
}

# unpauseValidator <account> <timeout-seconds>
# e.g. unpauseValidator validator 180
function unpauseValidator() {
    local ACCOUNT=$1
    local TIMEOUT=$2
    (! $(isNaturalNumber $TIMEOUT)) && TIMEOUT=180
    ($(isNullOrEmpty $ACCOUNT)) && echoInfo "INFO: Account name was not defined " && return 1
    sekaid tx customslashing unpause --from "$ACCOUNT" --chain-id=$NETWORK_NAME --keyring-backend=test --home=$SEKAID_HOME --fees 100ukex --yes --broadcast-mode=async --log_format=json --output=json | txAwait $TIMEOUT
}

# whitelistPermission <account> <permission> <address> <timeout-seconds>
# e.g. whitelistPermission validator 11 kiraXXX..YYY 180
function whitelistPermission() {
    local KM_ACC=$1
    local PERM=$2
    local ADDR=$(showAddress $3)
    local TIMEOUT=$4
    ($(isNullOrEmpty $KM_ACC)) && echoInfo "INFO: Account name was not defined '$1'" && return 1
    ($(isNullOrEmpty $ADDR)) && echoInfo "INFO: Address name was not defined '$3'" && return 1
    (! $(isNaturalNumber $PERM)) && echoInfo "INFO: Invalid permission id '$PERM' " && return 1
    (! $(isNaturalNumber $TIMEOUT)) && TIMEOUT=180
    if ($(isPermWhitelisted $ADDR $PERM)) ; then
        echoWarn "WARNING: Address '$ADDR' already has assigned permission '$PERM'"
    else
        sekaid tx customgov permission whitelist-permission --from "$KM_ACC" --keyring-backend=test --permission="$PERM" --addr="$ADDR" --chain-id=$NETWORK_NAME --home=$SEKAID_HOME --fees=100ukex --yes --broadcast-mode=async --log_format=json --output=json | txAwait $TIMEOUT
    fi
}

# blacklisttPermission <account> <permission> <address> <timeout-seconds>
# e.g. blacklisttPermission validator 11 kiraXXX..YYY 180
function blacklistPermission() {
    local KM_ACC=$1
    local PERM=$2
    local ADDR=$(showAddress $3)
    local TIMEOUT=$4
    ($(isNullOrEmpty $KM_ACC)) && echoInfo "INFO: Account name was not defined '$1'" && return 1
    ($(isNullOrEmpty $ADDR)) && echoInfo "INFO: Address name was not defined '$3'" && return 1
    (! $(isNaturalNumber $PERM)) && echoInfo "INFO: Invalid permission id '$PERM' " && return 1
    (! $(isNaturalNumber $TIMEOUT)) && TIMEOUT=180
    if ($(isPermBlacklisted $ADDR $PERM)) ; then
        echoWarn "WARNING: Address '$ADDR' already has blacklisted permission '$PERM'"
    else
        sekaid tx customgov permission blacklist-permission --from "$KM_ACC" --keyring-backend=test --permission="$PERM" --addr="$ADDR" --chain-id=$NETWORK_NAME --home=$SEKAID_HOME --fees=100ukex --yes --broadcast-mode=async --log_format=json --output=json | txAwait $TIMEOUT
    fi
}

function showCurrentPlan() {
    sekaid query upgrade current-plan --output=json --chain-id=$NETWORK_NAME --home=$SEKAID_HOME
}

function showNextPlan() {
    sekaid query upgrade next-plan --output=json --chain-id=$NETWORK_NAME --home=$SEKAID_HOME
}

# showIdentityRecord <account> <key> // shows all or a single key
# e.g. showIdentityRecord validator "mykey"
# e.g. showIdentityRecord validator 15
function showIdentityRecord() {
    local KM_ACC=$(showAddress $1)
    local KM_KEY=$2 && [ "$KM_KEY" == "*" ] && KM_KEY=""

    ($(isNullOrEmpty $KM_ACC)) && echoErr "ERROR: Account name or address '$1' is invalid" && return 1
    if ($(isNullOrEmpty $KM_KEY)) ; then
        sekaid query customgov identity-records-by-addr $KM_ACC --output=json | jq 2> /dev/null || echo ""
    else
        if ($(isNumber $KM_KEY)) ; then
            sekaid query customgov identity-records-by-addr $KM_ACC --output=json | jq ".records | .[] | select(.id==\"$KM_KEY\")" 2> /dev/null || echo ""
        else
            sekaid query customgov identity-records-by-addr $KM_ACC --output=json | jq ".records | .[] | select(.key==\"$KM_KEY\")" 2> /dev/null || echo ""
        fi
    fi
}

# upsertIdentityRecord <account> <key> <value> <timeout-seconds>
# e.g. upsertIdentityRecord validator "mykey" "My Value" 180
function upsertIdentityRecord() {
    local KM_ACC=$(showAddress $1)
    local IR_KEY=$2
    local IR_VAL=$3
    local TIMEOUT=$4
    ($(isNullOrEmpty $KM_ACC)) && echoErr "ERROR: Account name was NOT defined " && return 1
    ($(isNullOrEmpty $IR_KEY)) && echoErr "ERROR: Key was NOT defined " && return 1
    (! $(isNaturalNumber $TIMEOUT)) && TIMEOUT=180

    if ($(isNullOrEmpty $IR_VAL)) ; then
        sekaid tx customgov delete-identity-records --keys="$IR_KEY" --from=$KM_ACC --keyring-backend=test --home=$SEKAID_HOME --chain-id=$NETWORK_NAME --fees=100ukex --yes --broadcast-mode=async --log_format=json --output=json | txAwait $TIMEOUT
    else
        sekaid tx customgov register-identity-records --infos-json="{\"$IR_KEY\":\"$IR_VAL\"}" --from=$KM_ACC --keyring-backend=test --home=$SEKAID_HOME --chain-id=$NETWORK_NAME --fees=100ukex --yes --broadcast-mode=async --log_format=json --output=json | txAwait $TIMEOUT
    fi
}

# verifyIdentityRecord <account> <verifier-address> <one-or-many-comma-separated-keys/ids> <tip> <timeout>
# e.g. verifyIdentityRecord validator $(showAddress test) "mykey,mykey2" "200ukex" 180
function verifyIdentityRecord() {
    local KM_ACC=$1
    local KM_VER=$(showAddress $2)
    local KM_KEYS=$3
    local KM_TIP=$4
    local TIMEOUT=$5
    ($(isNullOrEmpty $KM_ACC)) && echoErr "ERROR: Account name was NOT defined '$1'" && return 1
    ($(isNullOrEmpty $KM_VER)) && echoErr "ERROR: Verifier address '$2' is invalid" && return 1
    ($(isNullOrEmpty $KM_KEYS)) && echoErr "ERROR: Record keys to verify were NOT specified '$3'" && return 1
    (! $(isNaturalNumber $TIMEOUT)) && TIMEOUT=180

    local FINAL_IR_KEYS=""
    for irkey in $(echo $KM_KEYS | sed "s/,/ /g") ; do
        local irkey_id=$(showIdentityRecord $KM_ACC "${irkey,,}" | jsonParse ".id" 2> /dev/null || echo "")
        ($(isNullOrEmpty $irkey_id)) && echoErr "ERROR: Key '$irkey' is invalid or was NOT found" && return 1
        [ ! -z "$FINAL_IR_KEYS" ] && FINAL_IR_KEYS="${FINAL_IR_KEYS},"
        FINAL_IR_KEYS="${FINAL_IR_KEYS}${irkey_id}"
    done
    ($(isNullOrEmpty $FINAL_IR_KEYS)) && echoErr "ERROR: No valid record keys were found" && return 1

    echoInfo "INFO: Sending request to verify '$FINAL_IR_KEYS'"
    sekaid tx customgov request-identity-record-verify --verifier="$KM_VER" --record-ids="$FINAL_IR_KEYS" --from=$KM_ACC --tip="$KM_TIP" --keyring-backend=test --home=$SEKAID_HOME --chain-id=$NETWORK_NAME --fees=100ukex --yes --broadcast-mode=async --log_format=json --output=json | txAwait $TIMEOUT
}

# showIdentityVerificationRequests <verifier-account/address> <requester-address>
# e.g. showIdentityVerificationRequests validator $(showAddress test)
function showIdentityVerificationRequests() {
    local KM_ACC=$(showAddress $1)
    local KM_REQ=$(showAddress $2)

    ($(isNullOrEmpty $KM_ACC)) && echoErr "ERROR: Account name or address '$1' is invalid" && return 1
    if ($(isNullOrEmpty $KM_REQ)) ; then
        sekaid query customgov identity-record-verify-requests-by-approver $KM_ACC --output=json | jq 2> /dev/null || echo ""
    else
        sekaid query customgov identity-record-verify-requests-by-approver $KM_ACC --output=json | jq ".verify_records | .[] | select(.address==\"$KM_REQ\")" 2> /dev/null || echo ""
    fi
}

# approveIdentityVerificationRequest <account> <id> <timeout>
# e.g. approveIdentityVerificationRequest validator 1 180
function approveIdentityVerificationRequest() {
    local KM_ACC=$1
    local KM_REQ=$2
    local TIMEOUT=$3 && (! $(isNaturalNumber $TIMEOUT)) && TIMEOUT=180
    ($(isNullOrEmpty $KM_ACC)) && echoErr "ERROR: Account name was NOT defined '$1'" && return 1
    (! $(isNaturalNumber $KM_REQ)) && echoErr "ERROR: Request Id must be a valid natural number, but got '$KM_REQ'" && return 1
    
    sekaid tx customgov handle-identity-records-verify-request $KM_REQ --approve="true" --from=$KM_ACC --keyring-backend=test --home=$SEKAID_HOME --chain-id=$NETWORK_NAME --fees=100ukex --yes --broadcast-mode=async --log_format=json --output=json | txAwait $TIMEOUT
}

# rejectIdentityVerificationRequest <account> <id> <timeout>
# e.g. rejectIdentityVerificationRequest validator 1 180
function rejectIdentityVerificationRequest() {
    local KM_ACC=$1
    local KM_REQ=$2
    local TIMEOUT=$3 && (! $(isNaturalNumber $TIMEOUT)) && TIMEOUT=180
    ($(isNullOrEmpty $KM_ACC)) && echoErr "ERROR: Account name was NOT defined " && return 1
    (! $(isNaturalNumber $KM_REQ)) && echoErr "ERROR: Request Id must be a valid natural number, but got '$KM_REQ'" && return 1
    
    sekaid tx customgov handle-identity-records-verify-request $KM_REQ --approve="false" --from=$KM_ACC --keyring-backend=test --home=$SEKAID_HOME --chain-id=$NETWORK_NAME --fees=100ukex --yes --broadcast-mode=async --log_format=json --output=json | txAwait $TIMEOUT
}

