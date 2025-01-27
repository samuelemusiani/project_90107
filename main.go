package main

import (
	"embed"
	"log/slog"
	"os"

	"samuelemusiani/project_90107/config"
	"samuelemusiani/project_90107/db"
	"samuelemusiani/project_90107/handler"

	_ "github.com/lib/pq"
)

//go:embed sql
var sqlFiles embed.FS

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)

	err := config.Parse()
	if err != nil {
		slog.With("Error", err).Error("Error parsing config")
		os.Exit(1)
	}

	conf := config.Get()
	slog.With("Config", conf).Debug("Config loaded")

	db, err := db.Init(conf, sqlFiles)
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
