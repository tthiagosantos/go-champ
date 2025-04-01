package handlers

import (
	"github.com/tthiagosantos/gochamp/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListarCampeonatos(c *gin.Context) {
	campeonatos, err := services.BuscarCampeonatos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, campeonatos)
}
