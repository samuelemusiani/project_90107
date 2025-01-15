package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	_ "github.com/lib/pq"
	"samuelemusiani/project_90107/config"
	"samuelemusiani/project_90107/handler"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)

	err := config.Parse()
	if err != nil {
		slog.With("Error", err).Error("Error parsing config")
		os.Exit(1)
	}

	conf := config.Get()
	slog.With("Config", conf).Debug("Config loaded")

	db, err := initDB(conf)
	if err != nil {
		slog.With("Error", err).Error("Error initializing database")
		os.Exit(1)
	}

	err = handler.InitAndServe(conf, db)
	if err != nil {
		slog.With("Error", err).Error("Error initializing server")
		os.Exit(1)
	}
}

func initDB(conf *config.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		conf.DB.User, conf.DB.Password, conf.DB.Host, conf.DB.Port, conf.DB.DBName)

	slog.With("connStr", connStr).Debug("Connecting to database")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	err = initUsers(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func initUsers(db *sql.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS users(
      username VARCHAR(20) PRIMARY KEY
    )
  `
	_, err := db.Exec(query)
	return err
}
