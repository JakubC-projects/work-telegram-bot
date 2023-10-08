package log

import (
	"os"

	"log/slog"
)

var L *slog.Logger

func init() {

	var handler slog.Handler
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	serviceName := os.Getenv("K_SERVICE")
	if serviceName != "" {
		handler = slog.NewJSONHandler(os.Stderr, opts)
	} else {
		handler = slog.NewTextHandler(os.Stderr, opts)
	}

	L = slog.New(handler)
}
