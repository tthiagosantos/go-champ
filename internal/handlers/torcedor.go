package handlers

import (
	"github.com/tthiagosantos/gochamp/internal/models"
	"github.com/tthiagosantos/gochamp/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CadastrarTorcedor(c *gin.Context) {
	var input models.Torcedor
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	if input.Nome == "" || input.Email == "" || input.Time == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Campos obrigatórios: nome, email, time"})
		return
	}

	novoTorcedor, err := services.SalvarTorcedor(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":       novoTorcedor.ID,
		"nome":     novoTorcedor.Nome,
		"email":    novoTorcedor.Email,
		"time":     novoTorcedor.Time,
		"mensagem": "Cadastro realizado com sucesso",
	})
}
