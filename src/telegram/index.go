package telegram

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func TelegramUpdateHttpHandler(c *gin.Context) {
	var update tgbotapi.Update
	err := c.BindJSON(&update)
	if err != nil {
		slog.Error("Unexpected telegram update body",
			"err", err)
		c.AbortWithError(http.StatusBadRequest, err)
	}
	err = HandleUpdate(c.Request.Context(), update)

	if err != nil {
		slog.Error("Error while handling telegram update ",
			"err", err)
		c.AbortWithError(http.StatusBadRequest, err)
	}

	c.Status(200)
}

func HandleUpdate(ctx context.Context, u tgbotapi.Update) error {
	slog.DebugContext(ctx, "Received message",
		"update", u)
	return nil
}
