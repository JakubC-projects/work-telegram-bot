package telegram

import (
	"github.com/JakubC-projects/work-telegram-bot/src/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func sendRegMessage(chatId int64, messageID int, state models.RegistrationState) tgbotapi.Chattable {
	text := "Successfully sent registration for " + state.Reg.Name

	return tgbotapi.NewEditMessageTextAndMarkup(chatId, messageID, text, tgbotapi.InlineKeyboardMarkup{InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{}})
}
