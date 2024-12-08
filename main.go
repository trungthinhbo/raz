//go:generate tailwindcss -i ./static/css/input.css -o ./static/css/style.css --minify
package main

import (
	"context"
	"embed"
	"log/slog"
	"os"
	"os/signal"
	"raz/internal/app"

	"github.com/joho/godotenv"
)

//go:embed templates/*.html static migrations/*.sql
var files embed.FS

func main() {
	godotenv.Load()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	app := app.New(logger, app.Config{}, files)

	if err := app.Start(ctx); err != nil {
		logger.Error("failed to start app", slog.Any("error", err))
	}
}
