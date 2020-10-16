#!/usr/local/bin/bash

set -e

while :; do
  case "$1" in
  -d | --demo)
    mode="demo"
    title="Demo Mode (local testnet)"
    shift
    ;;
  -f | --full)
    mode="full"
    title="Full Node Mode"
    shift
    ;;
  -v | --validator)
    mode="validator"
    title="Validator Mode"
    shift
    ;;
  -h | --help)
    display_help # Call your function
    # no shifting needed here, we're done.
    exit 0
    ;;
  --) # End of all options
    shift
    break
    ;;
  -*)
    echo "Error: Unknown option: $1" >&2
    exit 1
    ;;
  *) # No more options
    break
    ;;
  esac
done

chmod +x $KIRA_WORKDIR/utils.sh
source $KIRA_WORKDIR/utils.sh

printWidth=47
echo -e "-------------------------------------------------"
displayAlign center $printWidth "$title"
displayAlign center $printWidth "$(date '+%d/%m/%Y %H:%M:%S')"
echo -e "|-----------------------------------------------|"
displayAlign left $printWidth " [1] | Quick Setup"
displayAlign left $printWidth " [2] | Advanced Setup"
echo "|-----------------------------------------------|"
displayAlign left $printWidth " [X] | Exit | [B] | Go Back"
echo -e "-------------------------------------------------"

while :; do
  echo -en "Input option: "

  read -n 1 KEY

  echo ""

  case ${KEY,,} in
  1*)
    echo "INFO: Starting Quick Setup..."
    chmod +x $KIRA_WORKDIR/deploy/$mode/setup.sh
    $KIRA_WORKDIR/deploy/$mode/setup.sh -q
    break
    ;;

  2*)
    echo "INFO: Starting Advanced Setup..."
    chmod +x $KIRA_WORKDIR/deploy/$mode/setup.sh
    $KIRA_WORKDIR/deploy/$mode/setup.sh -a
    break
    ;;

  x*)
    exit 0
    ;;

  *)
    echo "Try again."
    ;;
  esac
done
