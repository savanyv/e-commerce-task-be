package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/savanyv/e-commerce-task-be/config"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitPostgres(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.PG.HostDB,
		cfg.PG.PortDB,
		cfg.PG.UserDB,
		cfg.PG.PassDB,
		cfg.PG.NameDB,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("failed to open database connection to Postgres: ", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatal("failed to ping postgres: ", err)
		return nil, err
	}

	DB = db

	log.Println("Connected to postgres Successfully")
	return db, nil
}
