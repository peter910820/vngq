package main

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/joho/godotenv"
	"github.com/lmittmann/tint"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	w := os.Stderr
	logger := slog.New(
		tint.NewHandler(w, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Stamp,
		}),
	)
	slog.SetDefault(logger)
}

func main() {

	app := fiber.New(fiber.Config{})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("ALLOW_ORIGINS")},
		AllowMethods:     []string{"POST", "GET", "DELETE"},
		AllowCredentials: true,
	}))

	// Serve built frontend assets from public/dist.
	app.Get("/*", static.New("./web/dist", static.Config{
		NotFoundHandler: func(c fiber.Ctx) error {
			return c.SendFile("./web/dist/index.html")
		},
	}))

	addr := fmt.Sprintf("127.0.0.1:%s", os.Getenv("PORT"))
	if err := app.Listen(addr); err != nil {
		slog.Error("start fiber server error", "address", addr, "error", err)
		os.Exit(1)
	}
}
