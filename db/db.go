package db

import (
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"log/slog"
	"time"

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

	triggers1, err := sqlFiles.ReadFile("sql/triggers1.sql")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(string(triggers1))
	if err != nil {
		return nil, errors.Join(err, errors.New("Error creating triggers1"))
	}

	triggers2, err := sqlFiles.ReadFile("sql/triggers2.sql")
	if err != nil {
		return nil, err
	}

	row := db.QueryRow("SELECT tgname FROM pg_trigger WHERE tgname ='trigger_check_max_carte'")
	var tgname string
	err = row.Scan(&tgname)
	if err != nil {
		if err == sql.ErrNoRows {
			_, err = db.Exec(string(triggers2))
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

func CreateMazzo() (int64, error) {
	sqlFiles, err := sqlFiles.ReadFile("sql/op08_1.sql")
	_, err = db.Query(string(sqlFiles) + ";")

	if err != nil {
		return 0, err
	}

	row := db.QueryRow("SELECT * FROM mazzo ORDER BY id DESC LIMIT 1")
	var id int64
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}

	slog.With("id", id).Debug("Mazzo created")

	return id, nil
}

func InsertFormato(mazzo int64, carta int64) error {
	sqlFiles, err := sqlFiles.ReadFile("sql/op08_2.sql")
	if err != nil {
		return err
	}

	_, err = db.Exec(string(sqlFiles), mazzo, carta)
	return err
}

func InsertBiglietto(b types.Biglietto) (int64, error) {
	sqlFiles, err := sqlFiles.ReadFile("sql/op10_1.sql")
	if err != nil {
		return 0, err
	}

	_, err = db.Exec(string(sqlFiles), b.Prezzo, b.Posto)
	if err != nil {
		return 0, err
	}

	row := db.QueryRow("SELECT lastval()")
	var id int64
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func InsertAssiste(a types.Assiste) error {
	sqlFiles, err := sqlFiles.ReadFile("sql/op10_2.sql")
	if err != nil {
		return err
	}

	_, err = db.Exec(string(sqlFiles), a.Persona, a.Evento, a.Biglietto)
	return err
}

func InsertCommenta(comm types.Commenta) error {
	sqlFiles, err := sqlFiles.ReadFile("sql/op11.sql")
	if err != nil {
		return err
	}

	_, err = db.Exec(string(sqlFiles), comm.Persona, comm.Partita, comm.Lingua)
	return err
}

func ClassificaCampionato(id string) ([]types.Classifica, error) {
	var classifica []types.Classifica

	sqlFiles, err := sqlFiles.ReadFile("sql/op12.sql")
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(string(sqlFiles), id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var tmp types.Classifica
		err = rows.Scan(&tmp.Posizione, &tmp.TeamOGiocatore)
		if err != nil {
			return nil, err
		}
		classifica = append(classifica, tmp)
	}

	return classifica, nil
}

// L'ID all'interno della carta indica il numero di vole che e' stata usata
func CartaPiuUsataCampionato(id string) (carta types.Carta, err error) {
	sqlFiles, err := sqlFiles.ReadFile("sql/op13.sql")
	if err != nil {
		return
	}

	row := db.QueryRow(string(sqlFiles), id)

	err = row.Scan(&carta.ID, &carta.Nome, &carta.Elisir, &carta.Danni)
	return
}

// L'ID all'interno della carta indica il numero di vole che e' stata usata
func CartaPiuUsataEvento(id string) (carta types.Carta, err error) {
	sqlFiles, err := sqlFiles.ReadFile("sql/op14.sql")
	if err != nil {
		return
	}

	row := db.QueryRow(string(sqlFiles), id)

	err = row.Scan(&carta.ID, &carta.Nome, &carta.Elisir, &carta.Danni)
	return
}

func UpdateIngaggio(giocatore string, newteam string, salario int64) error {

	row := db.QueryRow("SELECT id FROM Team WHERE nome = $1", newteam)
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

	sqlFiles, err := sqlFiles.ReadFile("sql/op32.sql")
	if err != nil {
		return err
	}

	_, err = db.Exec(string(sqlFiles), giocatoreID, teamID, salario)
	return err
}

func ViewTeam(team string) (*types.ReturnTeam, error) {
	row := db.QueryRow("SELECT id FROM Team WHERE nome = $1", team)
	var teamID string
	err := row.Scan(&teamID)
	if err != nil {
		return nil, errors.New("Team not found")
	}

	var teamstruct struct {
		ID              int64     `json:"id"`
		Nome            string    `json:"nome"`
		Logo            []byte    `json:"logo"`
		DataFondazione  time.Time `json:"dataFondazione"`
		StatoGeografico string    `json:"statoGeografico"`
		GPG             int64     `json:"gpg"`
		GPV             int64     `json:"gpv"`
		EtaMedia        float64   `json:"etaMedia"`
		Coach           string    `json:"coach"`
		Sponsor         string    `json:"sponsor"`
	}

	sqlFiles, err := sqlFiles.ReadFile("sql/op31.sql")
	if err != nil {
		return nil, err
	}

	row = db.QueryRow(string(sqlFiles), teamID)

	err = row.Scan(&teamstruct.Nome, &teamstruct.Logo, &teamstruct.DataFondazione, &teamstruct.StatoGeografico, &teamstruct.GPG, &teamstruct.GPV, &teamstruct.EtaMedia, &teamstruct.Coach, &teamstruct.Sponsor)
	if err != nil {
		return nil, err
	}

	var gpg, gpv string

	row = db.QueryRow("SELECT username FROM giocatore WHERE id = $1", teamstruct.GPG)
	err = row.Scan(&gpg)
	if err != nil {
		return nil, err
	}

	row = db.QueryRow("SELECT username FROM giocatore WHERE id = $1", teamstruct.GPV)
	err = row.Scan(&gpv)
	if err != nil {
		return nil, err
	}

	return &types.ReturnTeam{
		ID:              teamstruct.ID,
		Nome:            teamstruct.Nome,
		Logo:            teamstruct.Logo,
		DataFondazione:  teamstruct.DataFondazione,
		StatoGeografico: teamstruct.StatoGeografico,
		GPG:             gpg,
		GPV:             gpv,
		EtaMedia:        teamstruct.EtaMedia,
		Coach:           teamstruct.Coach,
		Sponsor:         teamstruct.Sponsor,
	}, nil
}

func ViewCommentatore(partita string) (*types.Persona, error) {
	var commentatore types.Persona

	sqlFiles, err := sqlFiles.ReadFile("sql/op30.sql")
	if err != nil {
		return nil, err
	}
	row := db.QueryRow(string(sqlFiles), partita)

	err = row.Scan(&commentatore.Nome, &commentatore.Cognome, &commentatore.DataNascita, &commentatore.LuogoNascita)
	if err != nil {
		return nil, err
	}

	return &commentatore, nil
}

func ViewBigliettiEvento(evento string) (int64, error) {
	var biglietti int64

	var eventoID int64
	row := db.QueryRow("SELECT id FROM evento WHERE nome = $1", evento)
	err := row.Scan(&eventoID)
	if err != nil {
		return 0, errors.New("Evento not found")
	}

	sqlFiles, err := sqlFiles.ReadFile("sql/op29.sql")
	if err != nil {
		return 0, err
	}
	row = db.QueryRow(string(sqlFiles), eventoID)

	err = row.Scan(&biglietti)
	if err != nil {
		return 0, err
	}

	return biglietti, nil
}
