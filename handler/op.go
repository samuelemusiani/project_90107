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
