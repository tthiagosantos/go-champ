package handlers

import (
	"github.com/tthiagosantos/gochamp/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BroadcastRequest struct {
	Tipo     string `json:"tipo"`
	Time     string `json:"time"`
	Placar   string `json:"placar,omitempty"`
	Mensagem string `json:"mensagem"`
}

func Broadcast(c *gin.Context) {
	var req BroadcastRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	if req.Time == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Campo 'time' é obrigatório"})
		return
	}
	
	err := services.EnviarBroadcast(req.Tipo, req.Time, req.Placar, req.Mensagem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Broadcast enviado com sucesso"})
}
