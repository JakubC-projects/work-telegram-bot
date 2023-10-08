package log

import (
	"os"

	"log/slog"
)

func Configure() {

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

	slog.SetDefault(slog.New(handler))
}
