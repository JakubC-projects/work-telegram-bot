package telegram

import (
	"context"

	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleUserAction(ctx context.Context, u tgbotapi.Update) error {
	callback := u.CallbackData()
	slog.Info("callback data", "callback", callback)
	// if callback != "" {
	// 	command, payload, _ := strings.Cut(callback, "-")
	// 	switch command {
	// 	case models.CommandStartChangeOrg:
	// 		return startChangeOrg(ctx, user)
	// 	}
	// }

	// msg, err := telegram.SendMenuMessage(user)
	// if err != nil {
	// 	return fmt.Errorf("cannot send menu message: %w", err)
	// }
	// user.LastMessageId = msg.MessageID

	// if err = db.SaveUser(ctx, user); err != nil {
	// 	return fmt.Errorf("cannot update user: %w", err)
	// }
	return nil
}
