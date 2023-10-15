curl -X POST \
    --data-urlencode "url=$SERVICE_URL/telegram-update" \
    https://api.telegram.org/bot$TELEGRAM_API_KEY/setWebhook
