package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/smakasaki/typing-trainer/internal/app"
)

func main() {
	// Загружаем переменные окружения из файла .env
	err := godotenv.Load()
	if err != nil {
		log.Printf("No .env file found, proceeding with system environment variables")
	}

	// Инициализируем приложение
	application, err := app.New()
	if err != nil {
		log.Fatalf("Failed to initialize the app: %v", err)
	}

	// Запускаем приложение
	if err := application.Run(); err != nil {
		log.Fatalf("Failed to run the app: %v", err)
	}
}
