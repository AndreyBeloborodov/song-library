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
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	if err = DB.Ping(); err != nil {
		return fmt.Errorf("database is unreachable: %w", err)
	}

	log.Println("Database connection established")
	return nil
}

// ApplyMigrations применяет миграции
func ApplyMigrations(migrationsDir string) error {
	goose.SetDialect("postgres")

	if err := goose.Up(DB, migrationsDir); err != nil {
		return err
	}

	log.Println("Migrations applied successfully")
	return nil
}
