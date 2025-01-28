package handler

import (
	"embed"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	"samuelemusiani/project_90107/config"
)

//go:embed dist
var frontend embed.FS

func InitAndServe(conf *config.Config) error {
	router := gin.Default()

	// middleware for static files (frontend)
	router.Use(static.Serve("/", static.EmbedFolder(frontend, "dist")))

	api := router.Group("/api")
	api.GET("/", root)
	api.POST("/giocatore", insertGiocatore)
	api.POST("/team", insertTeam)
	api.POST("/ingaggio", ingaggio)
	api.POST("/coach", insertCoach)
	api.POST("/sponsor", insertSponsor)
	api.POST("/campionato", insertCampionato)
	api.POST("/evento", insertEvento)
	api.POST("/partita", insertPartita)
	api.POST("/mazzo", insertMazzo)
	api.POST("/formato", insertFormato)
	api.POST("/carta", insertCarta)
	api.POST("/biglietto", insertBiglietto)
	api.POST("/assiste", insertAssiste)
	api.POST("/persona", insertPersona)
	api.POST("/commenta", insertCommenta)
	api.PUT("/ingaggio", updateIngaggio)
	api.GET("/classifica/:id_campionato", viewClassifica)
	api.GET("/carta_piu_usata/campionato/:id_campionato", viewCartaPiuUsataCampionato)
	api.GET("/carta_piu_usata/evento/:id_evento", viewCartaPiuUsataEvento)
	api.GET("/carta_piu_usata/partita/:id", viewCartaPiuUsataPartita)
	api.GET("/carte_usate/campionato/:id", viewCarteUsateCampionato)
	api.GET("/carte_usate/evento/:id", viewCarteUsateEvento)
	api.GET("/carte_usate/partita/:id", viewCarteUsatePartita)
	api.GET("/statistiche/campionato/:id_giocatore/:id_campionato", viewStatisticheCampionato)
	api.GET("/statistiche/evento/:id_giocatore/:id_evento", viewStatisticheEvento)
	api.GET("/statistiche/partita/:id_giocatore/:id_partita", viewStatistichePartita)
	api.GET("/team/:nome", viewTeam)
	api.GET("/commentatore/:id_partita", viewCommentatore)
	api.GET("/biglietti/evento/:nome_evento", viewBigliettiEvento)

	// Read index.html into memory
	index, err := frontend.ReadFile("dist/index.html")
	if err != nil {
		return err
	}

	// If no route match is probably a vue route. So we return the index.html
	// and the vue-router takes from here
	router.NoRoute(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.String(), "/api") {
			c.JSON(http.StatusNotFound, "")
			return
		}
		c.Data(http.StatusOK, "text/html", index)
	})

	router.Run(fmt.Sprintf(":%d", conf.Server.Port))
	return nil
}

func root(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}
