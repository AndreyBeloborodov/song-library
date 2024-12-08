package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // PostgreSQL драйвер
	"github.com/pressly/goose/v3"
	"log"
)

var DB *sql.DB

// InitDB инициализирует соединение с базой данных
func InitDB(host, port, user, password, dbname string) error {
	log.Println("[INFO] Initializing database connection...")

	// Формируем DSN (Data Source Name)
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)
	log.Printf("[DEBUG] DSN: %s", dsn)

	// Подключение к базе данных
	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Printf("[ERROR] Failed to open database connection: %v", err)
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// Проверка доступности базы данных
	if err = DB.Ping(); err != nil {
		log.Printf("[ERROR] Database is unreachable: %v", err)
		return fmt.Errorf("database is unreachable: %w", err)
	}

	log.Println("[INFO] Database connection established")
	return nil
}

// ApplyMigrations применяет миграции
func ApplyMigrations(migrationsDir string) error {
	log.Printf("[INFO] Applying migrations from directory: %s", migrationsDir)
	goose.SetDialect("postgres")

	// Применение миграций
	if err := goose.Up(DB, migrationsDir); err != nil {
		log.Printf("[ERROR] Failed to apply migrations: %v", err)
		return err
	}

	log.Println("[INFO] Migrations applied successfully")
	return nil
}
