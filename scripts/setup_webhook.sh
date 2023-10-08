#!/bin/bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
source $SCRIPT_DIR/../.env

curl -X POST \
    --data-urlencode "url=https://4ec9-193-160-15-74.ngrok.io/telegram-update" \
    https://api.telegram.org/bot$TELEGRAM_API_KEY/setWebhook
