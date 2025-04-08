package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/huuloc2026/pos-go.git/config"
	"github.com/huuloc2026/pos-go.git/internal/infrastructure/db"
	"github.com/huuloc2026/pos-go.git/pkg/logger"
	"github.com/spf13/viper"
)

func main() {
	// 1. Load config
	config.InitConfig()

	// // 2. Setup logger
	logger.InitLogger()

	// // 3. Connect database
	dbConn := db.InitDatabase()

	db.SetDB(dbConn)

	// 4. Create Fiber App
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World from Jake Onyx👋!")
	})

	// 6. Start server
	port := viper.GetString("app.port")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Server running at http://localhost:%s", port)
	log.Fatal(app.Listen(":" + port))
}
