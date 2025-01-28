package handler

import (
	"log/slog"
	"samuelemusiani/project_90107/db"
	"samuelemusiani/project_90107/types"
	"time"

	"github.com/gin-gonic/gin"
)

func insertGiocatore(c *gin.Context) {
	var fullGiocatore struct {
		Nome         string    `json:"nome"`
		Cognome      string    `json:"cognome"`
		DataNascita  time.Time `json:"dataNascita"`
		LuogoNascita string    `json:"luogoNascita"`
		Username     string    `json:"username"`
		DIG          time.Time `json:"dig"`
	}

	if err := c.BindJSON(&fullGiocatore); err != nil {
		slog.With("err", err).Error("Binding JSON")
		c.JSON(400, gin.H{
			"error": "Username not found"})
		return
	}

	id, err := db.InsertPersona(types.Persona{
		Nome:         fullGiocatore.Nome,
		Cognome:      fullGiocatore.Cognome,
		DataNascita:  fullGiocatore.DataNascita,
		LuogoNascita: fullGiocatore.LuogoNascita,
	})
	if err != nil {
		slog.With("err", err).Error("Inserting persona")
		c.JSON(500, gin.H{
			"error": "Error inserting persona",
		})
		return
	}

	err = db.InsertGiocatore(types.Giocatore{
		ID:       id,
		Username: fullGiocatore.Username,
		DIG:      fullGiocatore.DIG,
	})

	if err != nil {
		slog.With("err", err).Error("Inserting giocatore")
		c.JSON(500, gin.H{
			"error": "Error inserting giocatore",
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "Giocatore inserted",
	})
}

func insertTeam(c *gin.Context) {
	var team types.Team
	if err := c.BindJSON(&team); err != nil {
		slog.With("err", err).Error("Binding JSON")
		c.JSON(400, gin.H{
			"error": "Invalid JSON"})
		return
	}

	err := db.InsertTeam(team)
	if err != nil {
		slog.With("err", err).Error("Inserting team")
		c.JSON(500, gin.H{
			"error": "Error inserting team",
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "Team inserted",
	})
}

func ingaggio(c *gin.Context) {
	var ingaggio struct {
		Team      string `json:"team"`
		Giocatore string `json:"giocatore"`
		Salario   int64  `json:"salario"`
	}

	if err := c.BindJSON(&ingaggio); err != nil {
		slog.With("err", err).Error("Binding JSON")
		c.JSON(400, gin.H{
			"error": "Invalid JSON"})
		return
	}

	err := db.Ingaggia(ingaggio.Team, ingaggio.Giocatore, ingaggio.Salario)
	if err != nil {
		slog.With("err", err).Error("Ingaggio")
		c.JSON(500, gin.H{
			"error": "Error ingaggio",
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "Ingaggio effettuato",
	})
}

func insertCoach(c *gin.Context) {
	var fullCoach struct {
		Nome         string    `json:"nome"`
		Cognome      string    `json:"cognome"`
		DataNascita  time.Time `json:"dataNascita"`
		LuogoNascita string    `json:"luogoNascita"`
		Team         string    `json:"team"`
		Salario      int64     `json:"salario"`
	}

	if err := c.BindJSON(&fullCoach); err != nil {
		slog.With("err", err).Error("Binding JSON")
		c.JSON(400, gin.H{
			"error": "Username not found"})
		return
	}

	id, err := db.InsertPersona(types.Persona{
		Nome:         fullCoach.Nome,
		Cognome:      fullCoach.Cognome,
		DataNascita:  fullCoach.DataNascita,
		LuogoNascita: fullCoach.LuogoNascita,
	})
	if err != nil {
		slog.With("err", err).Error("Inserting persona")
		c.JSON(500, gin.H{
			"error": "Error inserting persona",
		})
		return
	}

	teamID, err := db.GetTeamIDByName(fullCoach.Team)
	if err != nil {
		slog.With("err", err).Error("Getting team ID")
		c.JSON(500, gin.H{
			"error": "Error getting team ID",
		})
		return
	}

	err = db.InsertCoach(types.Ingaggio{
		Persona: id,
		Team:    teamID,
		Salario: fullCoach.Salario,
	})

	if err != nil {
		slog.With("err", err).Error("Inserting coach")
		c.JSON(500, gin.H{
			"error": "Error inserting coach",
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "Ingaggio effettuato",
	})
}

func insertSponsor(c *gin.Context) {
	var fullSponsor struct {
		Nome   string `json:"nome"`
		Team   string `json:"team"`
		Budget int64  `json:"budget"`
	}

	if err := c.BindJSON(&fullSponsor); err != nil {
		slog.With("err", err).Error("Binding JSON")
		c.JSON(400, gin.H{
			"error": "Invalid JSON"})
		return
	}

	id, err := db.InsertSponsor(types.Sponsor{
		Nome: fullSponsor.Nome,
	})
	if err != nil {
		slog.With("err", err).Error("Inserting sponsor")
		c.JSON(500, gin.H{
			"error": "Error inserting sponsor",
		})
		return
	}

	teamID, err := db.GetTeamIDByName(fullSponsor.Team)

	err = db.InsertSponsorizza(types.Sponsorizza{
		Sponsor: id,
		Team:    teamID,
		Budget:  fullSponsor.Budget,
	})

	if err != nil {
		slog.With("err", err).Error("Inserting sponsorizza")
		c.JSON(500, gin.H{
			"error": "Error inserting sponsorizza",
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "Sponsor inserted",
	})
}

func insertCampionato(c *gin.Context) {
	var campionato types.Campionato
	if err := c.BindJSON(&campionato); err != nil {
		slog.With("err", err).Error("Binding JSON")
		c.JSON(400, gin.H{
			"error": "Invalid JSON"})
		return
	}

	err := db.InsertCampionato(campionato)
	if err != nil {
		slog.With("err", err).Error("Inserting campionato")
		c.JSON(500, gin.H{
			"error": "Error inserting campionato",
		})
		return
	}

	c.JSON(204, gin.H{
		"message": "Campionato inserted",
	})
}

func insertEvento(c *gin.Context) {
	var fullEvento struct {
		Nome        string    `json:"nome"`
		Luogo       string    `json:"luogo"`
		Data        time.Time `json:"data"`
		PostiTotali int64     `json:"postiTotali"`
		Campionato  string    `json:"campionato"`
	}

	if err := c.BindJSON(&fullEvento); err != nil {
		slog.With("err", err).Error("Binding JSON")
		c.JSON(400, gin.H{
			"error": "Invalid JSON"})
		return
	}

	id, err := db.GetCampionatoIDByName(fullEvento.Campionato)
	if err != nil {
		slog.With("err", err).Error("Getting campionato ID")
		c.JSON(500, gin.H{
			"error": "Error getting campionato ID",
		})
		return
	}

	err = db.InsertEvento(types.Evento{
		Nome:        fullEvento.Nome,
		Luogo:       fullEvento.Luogo,
		Data:        fullEvento.Data,
		PostiTotali: fullEvento.PostiTotali,
		Campionato:  id,
	})

	if err != nil {
		slog.With("err", err).Error("Inserting evento")
		c.JSON(500, gin.H{
			"error": "Error inserting evento",
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "Evento inserted",
	})
}

func insertPartita(c *gin.Context) {
	var partita struct {
		Vincitore       bool   `json:"vincitore"` // true = giocatore1, false = giocatore2
		Orario          string `json:"orario"`
		Tempo           int64  `json:"tempo"`
		Evento          string `json:"evento"`
		Giocatore1      string `json:"giocatore1"`
		Giocatore2      string `json:"giocatore2"`
		Mazzo1          int64  `json:"mazzo1"`
		Mazzo2          int64  `json:"mazzo2"`
		ElisirUsato1    int64  `json:"elisirUsato1"`
		ElisirUsato2    int64  `json:"elisirUsato2"`
		ElisirSprecato1 int64  `json:"elisirSprecato1"`
		ElisirSprecato2 int64  `json:"elisirSprecato2"`
		DanniFatti1     int64  `json:"danniFatti1"`
		Dannifatti2     int64  `json:"danniFatti2"`
		TipoTorri1      string `json:"tipoTorri1"`
		TipoTorri2      string `json:"tipoTorri2"`
	}

	if err := c.BindJSON(&partita); err != nil {
		slog.With("err", err).Error("Binding JSON")
		c.JSON(400, gin.H{
			"error": "Invalid JSON"})
		return
	}

	id1, err := db.GetGiocatoreIDByUsername(partita.Giocatore1)
	if err != nil {
		slog.With("err", err).Error("Getting giocatore ID")
		c.JSON(500, gin.H{
			"error": "Error getting giocatore ID",
		})
		return
	}

	id2, err := db.GetGiocatoreIDByUsername(partita.Giocatore2)
	if err != nil {
		slog.With("err", err).Error("Getting giocatore ID")
		c.JSON(500, gin.H{
			"error": "Error getting giocatore ID",
		})
		return
	}

	eventoID, err := db.GetEventoIDByName(partita.Evento)
	if err != nil {
		slog.With("err", err).Error("Getting evento ID")
		c.JSON(500, gin.H{
			"error": "Error getting evento ID",
		})
		return
	}

	partitaID, err := db.InsertPartita(types.Partita{
		Vincitore: id1,
		Orario:    partita.Orario,
		Tempo:     partita.Tempo,
		Evento:    eventoID,
	})
	if err != nil {
		slog.With("err", err).Error("Inserting partita")
		c.JSON(500, gin.H{
			"error": "Error inserting partita",
		})
		return
	}

	err = db.InsertGioca(types.Gioca{
		Giocatore:      id1,
		Partita:        partitaID,
		Mazzo:          partita.Mazzo1,
		ElisirUsato:    partita.ElisirUsato1,
		ElisirSprecato: partita.ElisirSprecato1,
		DanniFatti:     partita.DanniFatti1,
		TipoTorri:      partita.TipoTorri1,
	})

	if err != nil {
		slog.With("err", err).Error("Inserting gioca")
		c.JSON(500, gin.H{
			"error": "Error inserting gioca",
		})
		return
	}

	err = db.InsertGioca(types.Gioca{
		Giocatore:      id2,
		Partita:        partitaID,
		Mazzo:          partita.Mazzo2,
		ElisirUsato:    partita.ElisirUsato2,
		ElisirSprecato: partita.ElisirSprecato2,
		DanniFatti:     partita.Dannifatti2,
		TipoTorri:      partita.TipoTorri2,
	})

	if err != nil {
		slog.With("err", err).Error("Inserting gioca")
		c.JSON(500, gin.H{
			"error": "Error inserting gioca",
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "Partita inserted",
	})
}

func insertMazzo(c *gin.Context) {
	id, err := db.CreateMazzo()
	if err != nil {
		slog.With("err", err).Error("Creating mazzo")
		c.JSON(500, gin.H{
			"error": "Error creating mazzo",
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "Mazzo created",
		"id":      id,
	})
}

func insertCarta(c *gin.Context) {
	var carta types.Carta
	if err := c.BindJSON(&carta); err != nil {
		slog.With("err", err).Error("Binding JSON")
		c.JSON(400, gin.H{
			"error": "Invalid JSON"})
		return
	}

	err := db.InsertCarta(carta)
	if err != nil {
		slog.With("err", err).Error("Inserting carta")
		c.JSON(500, gin.H{
			"error": "Error inserting carta",
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "Carta inserted",
	})
}

func insertFormato(c *gin.Context) {
	var formato struct {
		Mazzo int64  `json:"mazzo"`
		Carta string `json:"carta"`
	}
	if err := c.BindJSON(&formato); err != nil {
		slog.With("err", err).Error("Binding JSON")
		c.JSON(400, gin.H{
			"error": "Invalid JSON"})
		return
	}

	cartaID, err := db.GetCartaIDByName(formato.Carta)
	if err != nil {
		slog.With("err", err).Error("Getting carta ID")
		c.JSON(500, gin.H{
			"error": "Error getting carta ID",
		})
		return
	}

	err = db.InsertFormato(formato.Mazzo, cartaID)
	if err != nil {
		slog.With("err", err).Error("Inserting formato")
		c.JSON(500, gin.H{
			"error": "Error inserting formato",
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "Formato inserted",
	})
}

func insertBiglietto(c *gin.Context) {
	var biglietto types.Biglietto
	if err := c.BindJSON(&biglietto); err != nil {
		slog.With("err", err).Error("Binding JSON")
		c.JSON(400, gin.H{
			"error": "Invalid JSON"})
		return
	}

	id, err := db.InsertBiglietto(biglietto)
	if err != nil {
		slog.With("err", err).Error("Inserting biglietto")
		c.JSON(500, gin.H{
			"error": "Error inserting biglietto",
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "Biglietto inserted",
		"id":      id,
	})
}

func insertAssiste(c *gin.Context) {
	var assiste types.Assiste
	if err := c.BindJSON(&assiste); err != nil {
		slog.With("err", err).Error("Binding JSON")
		c.JSON(400, gin.H{
			"error": "Invalid JSON"})
		return
	}

	err := db.InsertAssiste(assiste)
	if err != nil {
		slog.With("err", err).Error("Inserting assiste")
		c.JSON(500, gin.H{
			"error": "Error inserting assiste",
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "Assiste inserted",
	})
}

func insertPersona(c *gin.Context) {
	var p types.Persona
	if err := c.BindJSON(&p); err != nil {
		slog.With("err", err).Error("Binding JSON")
		c.JSON(400, gin.H{
			"error": "INvalid JSON"})
		return
	}

	id, err := db.InsertPersona(p)
	if err != nil {
		slog.With("err", err).Error("Inserting persona")
		c.JSON(500, gin.H{
			"error": "Error inserting persona",
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "persona inserted",
		"id":      id,
	})
}

func insertCommenta(c *gin.Context) {
	var comm types.Commenta
	if err := c.BindJSON(&comm); err != nil {
		slog.With("err", err).Error("Biding JSON")
		c.JSON(400, gin.H{
			"error": "Invalid JSON",
		})
	}

	err := db.InsertCommenta(comm)
	if err != nil {
		slog.With("err", err).Error("Inserting commenta")
		c.JSON(500, gin.H{
			"error": "Error inserting commenta",
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "commenta inserted",
	})
}

func updateIngaggio(c *gin.Context) {
	var ingaggio struct {
		Team      string `json:"team"`
		Giocatore string `json:"giocatore"`
		Salario   int64  `json:"salario"`
	}

	if err := c.BindJSON(&ingaggio); err != nil {
		slog.With("err", err).Error("Binding JSON")
		c.JSON(400, gin.H{
			"error": "Invalid JSON"})
		return
	}

	err := db.UpdateIngaggio(ingaggio.Giocatore, ingaggio.Team, ingaggio.Salario)
	if err != nil {
		slog.With("err", err).Error("Update ingaggio")
		c.JSON(500, gin.H{
			"error": "Error upadate ingaggio",
		})
		return
	}

	c.JSON(204, gin.H{
		"message": "Update ingaggio effettuato",
	})
}

func viewTeam(c *gin.Context) {
	team := c.Param("nome")
	t, err := db.ViewTeam(team)
	if err != nil {
		slog.With("err", err).Error("View team")
		c.JSON(500, gin.H{
			"error": "Error view team",
		})
		return
	}

	c.JSON(200, t)
}

func viewCommentatore(c *gin.Context) {
	idPartita := c.Param("id_partita")
	t, err := db.ViewCommentatore(idPartita)
	if err != nil {
		slog.With("err", err).Error("View commentatore")
		c.JSON(500, gin.H{
			"error": "Error view commentatore",
		})
		return
	}

	c.JSON(200, t)
}

func viewBigliettiEvento(c *gin.Context) {
	nomeEvento := c.Param("nome_evento")
	t, err := db.ViewBigliettiEvento(nomeEvento)
	if err != nil {
		slog.With("err", err).Error("View biglietti")
		c.JSON(500, gin.H{
			"error": "Error view biglietti",
		})
		return
	}

	var ret struct {
		Numero int64 `json:"numero"`
	}
	ret.Numero = t

	c.JSON(200, ret)
}

func viewClassifica(c *gin.Context) {
	idCampionato := c.Param("id_campionato")
	cass, err := db.ClassificaCampionato(idCampionato)
	if err != nil {
		slog.With("err", err).Error("View classifica campionato")
		c.JSON(500, gin.H{
			"error": "Error view classifica campionato",
		})
		return
	}

	c.JSON(200, cass)
}

func viewCartaPiuUsataCampionato(c *gin.Context) {
	idCampionato := c.Param("id_campionato")
	carta, err := db.CartaPiuUsataCampionato(idCampionato)
	if err != nil {
		slog.With("err", err).Error("View carta piu' usata campionato")
		c.JSON(500, gin.H{
			"error": "Error view carta piu' usata campionato",
		})
		return
	}

	c.JSON(200, carta)
}

func viewCartaPiuUsataEvento(c *gin.Context) {
	idEvento := c.Param("id_evento")
	carta, err := db.CartaPiuUsataEvento(idEvento)
	if err != nil {
		slog.With("err", err).Error("View carta piu' usata evento")
		c.JSON(500, gin.H{
			"error": "Error view carta piu' usata evento",
		})
		return
	}

	c.JSON(200, carta)
}

func viewCartaPiuUsataPartita(c *gin.Context) {
	id := c.Param("id")
	carta, err := db.CartaPiuUsataPartita(id)
	if err != nil {
		slog.With("err", err).Error("View carta piu' usata partita")
		c.JSON(500, gin.H{
			"error": "Error view carta piu' usata partita",
		})
		return
	}

	c.JSON(200, carta)
}

func viewCarteUsateCampionato(c *gin.Context) {
	id := c.Param("id")
	carte, err := db.CarteUsateCampionato(id)
	if err != nil {
		slog.With("err", err).Error("View carte usate campionato")
		c.JSON(500, gin.H{
			"error": "Error view carte usate campionato",
		})
		return
	}

	c.JSON(200, carte)
}

func viewCarteUsateEvento(c *gin.Context) {
	id := c.Param("id")
	carte, err := db.CarteUsateEvento(id)
	if err != nil {
		slog.With("err", err).Error("View carte usate evento")
		c.JSON(500, gin.H{
			"error": "Error view carte usate evento",
		})
		return
	}

	c.JSON(200, carte)
}

func viewCarteUsatePartita(c *gin.Context) {
	id := c.Param("id")
	carte, err := db.CarteUsatePartita(id)
	if err != nil {
		slog.With("err", err).Error("View carte usate partita")
		c.JSON(500, gin.H{
			"error": "Error view carte usate partita",
		})
		return
	}

	c.JSON(200, carte)
}

func viewStatisticheCampionato(c *gin.Context) {
	idGiocatore := c.Param("id_giocatore")
	idCampionato := c.Param("id_campionato")
	stat, err := db.StatisticheCampionato(idGiocatore, idCampionato)
	if err != nil {
		slog.With("err", err).Error("View statistiche campionato")
		c.JSON(500, gin.H{
			"error": "Error view statistiche campionato",
		})
		return
	}

	c.JSON(200, stat)
}

func viewStatisticheEvento(c *gin.Context) {
	idGiocatore := c.Param("id_giocatore")
	idEvento := c.Param("id_evento")
	stat, err := db.StatisticheEvento(idGiocatore, idEvento)
	if err != nil {
		slog.With("err", err).Error("View statistiche evento")
		c.JSON(500, gin.H{
			"error": "Error view statistiche evento",
		})
		return
	}

	c.JSON(200, stat)
}

func viewStatistichePartita(c *gin.Context) {
	idGiocatore := c.Param("id_giocatore")
	idPartita := c.Param("id_partita")
	stat, err := db.StatistichePartita(idGiocatore, idPartita)
	if err != nil {
		slog.With("err", err).Error("View statistiche partita")
		c.JSON(500, gin.H{
			"error": "Error view statistiche partita",
		})
		return
	}

	c.JSON(200, stat)
}

func viewMazzoPiuUsatoCampionato(c *gin.Context) {
	id := c.Param("id")
	mazzo, err := db.MazzoPiuUsatoCampionato(id)
	if err != nil {
		slog.With("err", err).Error("View mazzo piu' usato campionato")
		c.JSON(500, gin.H{
			"error": "Error view mazzo piu' usato campionato",
		})
		return
	}

	c.JSON(200, mazzo)
}

func viewMazzoPiuUsatoEvento(c *gin.Context) {
	id := c.Param("id")
	mazzo, err := db.MazzoPiuUsatoEvento(id)
	if err != nil {
		slog.With("err", err).Error("View mazzo piu' usato evento")
		c.JSON(500, gin.H{
			"error": "Error view mazzo piu' usato evento",
		})
		return
	}

	c.JSON(200, mazzo)
}

func viewMazzoMiglioreCampionato(c *gin.Context) {
	id := c.Param("id")
	mazzo, err := db.MazzoMiglioreCampionato(id)
	if err != nil {
		slog.With("err", err).Error("View mazzo migliore campionato")
		c.JSON(500, gin.H{
			"error": "Error view mazzo migliore campionato",
		})
		return
	}

	c.JSON(200, mazzo)
}

func viewMazzoMiglioreEvento(c *gin.Context) {
	id := c.Param("id")
	mazzo, err := db.MazzoMiglioreEvento(id)
	if err != nil {
		slog.With("err", err).Error("View mazzo migliore evento")
		c.JSON(500, gin.H{
			"error": "Error view mazzo migliore evento",
		})
		return
	}

	c.JSON(200, mazzo)
}

func viewTipoTorriCampionato(c *gin.Context) {
	id := c.Param("id")
	tipo, err := db.TipoTorriCampionato(id)
	if err != nil {
		slog.With("err", err).Error("View tipo torri campionato")
		c.JSON(500, gin.H{
			"error": "Error view tipo torri campionato",
		})
		return
	}

	c.JSON(200, tipo)
}

func viewTipoTorriEvento(c *gin.Context) {
	id := c.Param("id")
	tipo, err := db.TipoTorriEvento(id)
	if err != nil {
		slog.With("err", err).Error("View tipo torri evento")
		c.JSON(500, gin.H{
			"error": "Error view tipo torri evento",
		})
		return
	}

	c.JSON(200, tipo)
}
