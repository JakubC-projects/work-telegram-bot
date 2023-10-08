package telegram

import (
	"os"

	"github.com/JakubC-projects/work-telegram-bot/src/config"
	"github.com/JakubC-projects/work-telegram-bot/src/log"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var client *tgbotapi.BotAPI

func init() {
	var err error

	client, err = tgbotapi.NewBotAPI(config.C.TelegramApiKey)
	if err != nil {
		log.L.Error("cannot create telegram bot api",
			"error", err)
		os.Exit(1)
	}
}
