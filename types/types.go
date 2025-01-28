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

type ReturnTeam struct {
	ID              int64     `json:"id"`
	Nome            string    `json:"nome"`
	Logo            []byte    `json:"logo"`
	DataFondazione  time.Time `json:"dataFondazione"`
	StatoGeografico string    `json:"statoGeografico"`
	GPG             string    `json:"gpg"`
	GPV             string    `json:"gpv"`
	EtaMedia        float64   `json:"etaMedia"`
	Coach           string    `json:"coach"`
	Sponsor         string    `json:"sponsor"`
}

type Ingaggio struct {
	Persona int64 `json:"persona"`
	Team    int64 `json:"team"`
	Salario int64 `json:"salario"`
}

type Sponsor struct {
	ID   int64  `json:"id"`
	Nome string `json:"nome"`
}

type Sponsorizza struct {
	Sponsor int64 `json:"sponsor"`
	Team    int64 `json:"team"`
	Budget  int64 `json:"budget"`
}

type Campionato struct {
	ID         int64     `json:"id"`
	Nome       string    `json:"nome"`
	Luogo      string    `json:"luogo"`
	DataInizio time.Time `json:"dataInizio"`
	DataFine   time.Time `json:"dataFine"`
	Tipo       string    `json:"tipo"`
	Montepremi int64     `json:"montepremi"`
}

type Evento struct {
	ID          int64     `json:"id"`
	Nome        string    `json:"nome"`
	Luogo       string    `json:"luogo"`
	Data        time.Time `json:"data"`
	PostiTotali int64     `json:"postiTotali"`
	Campionato  int64     `json:"campionato"`
}

type Partita struct {
	ID        int64  `json:"id"`
	Vincitore int64  `json:"vincitore"`
	Orario    string `json:"orario"`
	Tempo     int64  `json:"tempo"`
	Evento    int64  `json:"evento"`
}

type Gioca struct {
	Giocatore      int64  `json:"giocatore"`
	Partita        int64  `json:"partita"`
	Mazzo          int64  `json:"mazzo"`
	ElisirUsato    int64  `json:"elisirUsato"`
	ElisirSprecato int64  `json:"elisirSprecato"`
	DanniFatti     int64  `json:"danniFatti"`
	TipoTorri      string `json:"tipoTorri"`
}

type Carta struct {
	ID     int64  `json:"id"`
	Nome   string `json:"nome"`
	Elisir int64  `json:"elisir"`
	Danni  int64  `json:"danni"`
}

type Mazzo struct {
	ID int64 `json:"id"`
}

type Formato struct {
	Mazzo int64 `json:"mazzo"`
	Carta int64 `json:"carta"`
}

type Biglietto struct {
	ID     int64 `json:"id"`
	Prezzo int64 `json:"prezzo"`
	Posto  int64 `json:"evento"`
}

type Assiste struct {
	Persona   int64 `json:"persona"`
	Evento    int64 `json:"evento"`
	Biglietto int64 `json:"biglietto"`
}

type Commenta struct {
	Persona int64  `json:"persona"`
	Partita int64  `json:"partita"`
	Lingua  string `json:"lingua"`
}

type Classifica struct {
	Posizione      int64  `json:"posizione"`
	TeamOGiocatore string `json:"teamOGiocatore"`
}

type Statistiche struct {
	ElisirUsato    int64 `json:"elisirUsato"`
	ElisirSprecato int64 `json:"elisirSprecato"`
	DanniFatti     int64 `json:"danniFatti"`
	DanniSubiti    int64 `json:"danniSubiti"`
}

type MazzoVolte struct {
	Mazzo int64 `json:"mazzo"`
	Volte int64 `json:"volte"`
}

type VolteTorri struct {
	TipoTorri string `json:"tipoTorri"`
	Volte     int64  `json:"volte"`
}
