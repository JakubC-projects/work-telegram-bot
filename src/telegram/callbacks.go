package telegram

import (
	"context"
	"strings"

	"github.com/JakubC-projects/work-telegram-bot/src/log"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleCallback(ctx context.Context, callback *tgbotapi.CallbackQuery) error {
	command, _, _ := strings.Cut(callback.Data, ":")

	log.L.Debug(callback.Data)

	switch command {
	case "updateReg":
		return updateReg(ctx, callback)
	case "sendReg":
		return sendRegistration(ctx, callback)
	}

	return nil
}
