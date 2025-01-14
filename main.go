package main

import (
	"log/slog"
	"os"
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

	handler.InitAndServe(conf)
}
