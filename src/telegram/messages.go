package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func newWelcomeMessage(chatId int64) tgbotapi.Chattable {
	return tgbotapi.NewMessage(chatId,
		`Welcome to WORK the work bot`)
}
