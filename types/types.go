package types

import "time"

type Persona struct {
	ID           int64     `json:"id"`
	Nome         string    `json:"nome"`
	Cognome      string    `json:"cognome"`
	DataNascita  time.Time `json:"dataNascita"`
	LuogoNascita string    `json:"luogoNascita"`
}

type Giocatore struct {
	ID       int64     `json:"id"`
	Username string    `json:"username"`
	DIG      time.Time `json:"dig"`
}

type Team struct {
	ID              int64     `json:"id"`
	Nome            string    `json:"nome"`
	Logo            []byte    `json:"logo"`
	DataFondazione  time.Time `json:"dataFondazione"`
	StatoGeografico string    `json:"statoGeografico"`
	GPG             int64     `json:"gpg"`
	GPV             int64     `json:"gpv"`
	EtaMedia        float64   `json:"etaMedia"`
}
