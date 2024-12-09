package main

import (
	"github.com/joho/godotenv"
	"github.com/swaggo/http-swagger/v2"
	"log"
	"net/http"
	"os"
	"song-libary/db"
	_ "song-libary/docs"
	handlers "song-libary/hendlers"
	"song-libary/repository"
	"song-libary/service"
)

func main() {
	log.Println("[INFO] Loading environment variables...")
	if err := godotenv.Load(); err != nil {
		log.Fatalf("[ERROR] Failed to load .env file: %v", err)
	}

	// Загрузка переменных окружения
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	migrationsPath := "./db/migrations"

	dbManager := db.NewDbManager()

	log.Println("[INFO] Initializing database connection...")
	if err := dbManager.InitDB(host, port, user, password, dbname); err != nil {
		log.Fatalf("[ERROR] Failed to initialize database: %v", err)
	}
	log.Printf("[INFO] Connected to database at %s:%s as user %s", host, port, user)

	log.Println("[INFO] Applying migrations...")
	if err := dbManager.ApplyMigrations(migrationsPath); err != nil {
		log.Fatalf("[ERROR] Failed to apply migrations: %v", err)
	}

	log.Println("[INFO] Setting up repositories, services, and handlers...")
	songRepo := repository.NewSongRepositorySqlDbImpl(dbManager.DB)
	songService := service.NewSongService(songRepo)
	songHandler := handlers.NewSongHandler(songService)

	log.Println("[INFO] Registering routes...")
	// Swagger UI доступен по адресу /swagger/index.html
	http.Handle("/swagger/", http.StripPrefix("/swagger", httpSwagger.WrapHandler))
	http.HandleFunc("/songs", songHandler.GetSongsHandler)
	http.HandleFunc("/songs/info", songHandler.InfoHandler)
	http.HandleFunc("/songs/add", songHandler.AddSongHandler)
	http.HandleFunc("/songs/delete", songHandler.DeleteSongHandler)
	http.HandleFunc("/songs/update", songHandler.UpdateSongHandler)
	http.HandleFunc("/songs/text", songHandler.GetSongTextHandler)

	log.Println("[INFO] Starting server on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("[ERROR] Server failed: %v", err)
	}
}
