#!/bin/bash
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
source $SCRIPT_DIR/../.env

echo $PORT
# Start NGROK in background
ngrok http $PORT > /dev/null &

# # Wait for ngrok to be available
while ! nc -z localhost 4040; do
  sleep 0.2 # wait Ngrok to be available
done

while [[ "$NGROK_REMOTE_URL" != *"http"* ]]; do
    NGROK_REMOTE_URL="$(curl -s http://localhost:4040/api/tunnels | jq ".tunnels[0].public_url")"
    sleep 0.2
done

# # Trim double quotes from variable
NGROK_REMOTE_URL=$(echo ${NGROK_REMOTE_URL} | tr -d '"')

echo $NGROK_REMOTE_URL

curl -X POST \
    --data-urlencode "url=$NGROK_REMOTE_URL/telegram-update" \
    https://api.telegram.org/bot$TELEGRAM_API_KEY/setWebhook

sleep infinity