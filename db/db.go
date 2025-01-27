package db

import (
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"log/slog"

	"samuelemusiani/project_90107/config"
)

var sqlFiles embed.FS
var db *sql.DB

func Init(conf *config.Config, sqlFiles embed.FS) (*sql.DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		conf.DB.User, conf.DB.Password, conf.DB.Host, conf.DB.Port, conf.DB.DBName)

	slog.With("connStr", connStr).Debug("Connecting to database")

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	schema, err := sqlFiles.ReadFile("sql/schema.sql")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(string(schema))
	if err != nil {
		return nil, errors.Join(err, errors.New("Error creating schema"))
	}

	triggers, err := sqlFiles.ReadFile("sql/triggers.sql")
	if err != nil {
		return nil, err
	}

	row := db.QueryRow("SELECT tgname FROM pg_trigger WHERE tgname ='trigger_check_max_carte'")
	var tgname string
	err = row.Scan(&tgname)
	if err != nil {
		if err == sql.ErrNoRows {
			_, err = db.Exec(string(triggers))
			if err != nil {
				return nil, errors.Join(err, errors.New("Error creating triggers"))
			}
		} else {
			return nil, err
		}
	}

	return db, nil
}
