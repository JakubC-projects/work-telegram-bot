package main

import (
	"fmt"

	"log/slog"

	"github.com/JakubC-projects/work-telegram-bot/src/config"
	"github.com/JakubC-projects/work-telegram-bot/src/telegram"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.POST("/telegram-update", telegram.TelegramUpdateHttpHandler)

	if config.C.Server.CertFile != "" {
		err := r.RunTLS(fmt.Sprintf(":%d", config.C.Server.Port), config.C.Server.CertFile, config.C.Server.CertKeyFile)
		if err != nil {
			slog.Error("cannot start tls server",
				"err", err)
		}
	} else {
		r.Run(fmt.Sprintf(":%d", config.C.Server.Port))
	}
}
