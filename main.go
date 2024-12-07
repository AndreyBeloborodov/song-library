package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"song-libary/db"
	handlers "song-libary/hendlers"
	"song-libary/repository"
	"song-libary/service"
)

func main() {
	// Загружаем переменные окружения из .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Загрузка переменных окружения
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	migrationsPath := "./db/migrations"

	// Инициализация базы данных
	if err := db.InitDB(host, port, user, password, dbname); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	log.Printf("Connecting to database at %s:%s as user %s\n", host, port, user)

	// Применение миграций
	if err := db.ApplyMigrations(migrationsPath); err != nil {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	// Создаем репозиторий, сервис и хендлер
	songRepo := repository.NewSongRepository(db.DB)
	songService := service.NewSongService(songRepo)
	songHandler := handlers.NewSongHandler(songService)

	// Регистрация маршрутов
	http.HandleFunc("/songs", songHandler.AddSongHandler)

	log.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
