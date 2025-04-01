package handlers

import (
	"github.com/tthiagosantos/gochamp/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListarPartidas(c *gin.Context) {
	idCampeonato := c.Param("id")
	if idCampeonato == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do campeonato é obrigatório"})
		return
	}

	equipe := c.Query("equipe")
	rodada := c.Query("rodada") // string, converta pra int se precisar

	partidas, err := services.BuscarPartidas(idCampeonato, equipe, rodada)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, partidas)
}
