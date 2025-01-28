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

	c.JSON(204, gin.H{
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

	c.JSON(204, gin.H{
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

	c.JSON(204, gin.H{
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

	c.JSON(204, gin.H{
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

	c.JSON(204, gin.H{
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

	c.JSON(204, gin.H{
		"message": "Evento inserted",
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
