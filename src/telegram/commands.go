package telegram

import (
	"context"
	"strings"

	"github.com/JakubC-projects/work-telegram-bot/src/log"
	"github.com/JakubC-projects/work-telegram-bot/src/workapi"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleCommand(ctx context.Context, m *tgbotapi.Message) error {
	chatId := m.Chat.ID
	command, _, _ := strings.Cut(m.Text, " ")

	switch command {
	case "/start":
		client.Send(newWelcomeMessage(chatId))
		break
	case "/register":
		err := startRegistration(ctx, m)
		if err != nil {
			return err
		}
		break

	case "/result":
		results, err := workapi.GetResults(ctx)
		if err != nil {
			log.L.Error("cannot get results", "err", err)
			client.Send(tgbotapi.NewMessage(chatId, "cannot get results"))
			return nil
		}
		client.Send(newResultsMessage(chatId, results))
		break
	}

	return nil
}
