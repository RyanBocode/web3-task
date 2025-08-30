package main

import (
	"log"
	"os"

	"blog/internal/config"
	"blog/internal/database"
	"blog/internal/router"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load() // 本地开发可加载 .env

	cfg := config.Load()

	if err := database.Init(cfg); err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	r := router.Setup(cfg)

	addr := ":" + cfg.Port
	log.Printf("server listening on %s (mode=%s, db=%s)\n", addr, os.Getenv("GIN_MODE"), cfg.DBDriver)
	if err := r.Run(addr); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
