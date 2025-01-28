package db

import (
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"log/slog"

	"samuelemusiani/project_90107/config"
	"samuelemusiani/project_90107/types"
)

var sqlFiles embed.FS
var db *sql.DB

func Init(conf *config.Config, sqlFS embed.FS) (*sql.DB, error) {
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

	sqlFiles = sqlFS

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

func InsertPersona(p types.Persona) (int64, error) {
	sqlFiles, err := sqlFiles.ReadFile("sql/op01_1.sql")
	if err != nil {
		return 0, err
	}

	row := db.QueryRow("SELECT id FROM persona WHERE nome = $1 AND cognome = $2", p.Nome, p.Cognome)
	var tmp int64
	err = row.Scan(&tmp)
	if err == nil {
		return 0, errors.New("Persona already exists")
	}

	slog.With("Persona", p).Debug("Inserting persona")
	_, err = db.Exec(string(sqlFiles), p.Nome, p.Cognome, p.DataNascita, p.LuogoNascita)
	if err != nil {
		return 0, err
	}

	row = db.QueryRow("SELECT id FROM persona WHERE nome = $1 AND cognome = $2 AND data_nascita = $3 AND luogo_nascita = $4", p.Nome, p.Cognome, p.DataNascita, p.LuogoNascita)
	var id int64
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, err
}

func InsertGiocatore(p types.Giocatore) error {
	sqlFiles, err := sqlFiles.ReadFile("sql/op01_2.sql")
	if err != nil {
		return err
	}

	_, err = db.Exec(string(sqlFiles), p.ID, p.Username, p.DIG)
	return err
}

func InsertTeam(t types.Team) error {
	sqlFiles, err := sqlFiles.ReadFile("sql/op02_1.sql")
	if err != nil {
		return err
	}

	_, err = db.Exec(string(sqlFiles), t.Nome, t.Logo, t.DataFondazione, t.StatoGeografico)
	return err
}

func Ingaggia(team string, giocatore string, salario int64) error {

	row := db.QueryRow("SELECT id FROM Team WHERE nome = $1", team)
	var teamID string
	err := row.Scan(&teamID)
	if err != nil {
		return errors.New("Team not found")
	}

	row = db.QueryRow("SELECT id FROM Giocatore WHERE username = $1", giocatore)
	var giocatoreID string
	err = row.Scan(&giocatoreID)
	if err != nil {
		return errors.New("Giocatore not found")
	}

	sqlFiles, err := sqlFiles.ReadFile("sql/op02_2.sql")
	if err != nil {
		return err
	}

	_, err = db.Exec(string(sqlFiles), teamID, giocatoreID, salario)
	return err
}

func InsertCoach(ing types.Ingaggio) error {
	sqlFiles, err := sqlFiles.ReadFile("sql/op03.sql")
	if err != nil {
		return err
	}

	_, err = db.Exec(string(sqlFiles), ing.Persona, ing.Team, ing.Salario)
	return err
}

func InsertSponsor(s types.Sponsor) (int64, error) {
	sqlFiles, err := sqlFiles.ReadFile("sql/op04_1.sql")
	if err != nil {
		return 0, err
	}

	row := db.QueryRow("SELECT id FROM sponsor WHERE nome = $1", s.Nome)
	var tmp int64
	err = row.Scan(&tmp)
	if err == nil {
		return 0, errors.New("Sponsor already exists")
	}

	_, err = db.Exec(string(sqlFiles), s.Nome)
	if err != nil {
		return 0, err
	}

	row = db.QueryRow("SELECT id FROM sponsor WHERE nome = $1", s.Nome)
	var id int64
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, err
}

func InsertSponsorizza(s types.Sponsorizza) error {
	sqlFiles, err := sqlFiles.ReadFile("sql/op04_2.sql")
	if err != nil {
		return err
	}

	_, err = db.Exec(string(sqlFiles), s.Sponsor, s.Team, s.Budget)
	return err
}

func InsertCampionato(c types.Campionato) error {
	sqlFiles, err := sqlFiles.ReadFile("sql/op05.sql")
	if err != nil {
		return err
	}

	_, err = db.Exec(string(sqlFiles), c.Nome, c.Luogo, c.DataInizio, c.DataFine, c.Tipo, c.Montepremi)
	return err
}

func InsertEvento(e types.Evento) error {
	sqlFiles, err := sqlFiles.ReadFile("sql/op06.sql")
	if err != nil {
		return err
	}

	_, err = db.Exec(string(sqlFiles), e.Nome, e.Luogo, e.Data, e.PostiTotali, e.Campionato)
	return err
}

func InsertPartita(p types.Partita) (int64, error) {
	sqlFiles, err := sqlFiles.ReadFile("sql/op07_1.sql")
	if err != nil {
		return 0, err
	}

	_, err = db.Exec(string(sqlFiles), p.Vincitore, p.Orario, p.Tempo, p.Evento)
	if err != nil {
		return 0, err
	}

	row := db.QueryRow("SELECT id FROM partita WHERE vincitore = $1 AND orario = $2 AND tempo = $3 AND evento = $4", p.Vincitore, p.Orario, p.Tempo, p.Evento)

	var id int64
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func InsertGioca(g types.Gioca) error {
	sqlFiles, err := sqlFiles.ReadFile("sql/op07_2.sql")
	if err != nil {
		return err
	}

	_, err = db.Exec(string(sqlFiles), g.Giocatore, g.Partita, g.Mazzo, g.ElisirUsato, g.ElisirSprerato, g.DanniFatti, g.TipoTorri)
	return err
}

func InsertCarta(g types.Carta) error {
	sqlFiles, err := sqlFiles.ReadFile("sql/op09.sql")
	if err != nil {
		return err
	}

	_, err = db.Exec(string(sqlFiles), g.Nome, g.Elisir, g.Danni)
	return err
}
