package handlers

import (
	"github.com/tthiagosantos/gochamp/internal/models"
	"github.com/tthiagosantos/gochamp/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CadastrarTorcedor(c *gin.Context) {
	var t models.Torcedor
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	if t.Nome == "" || t.Email == "" || t.Time == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nome, email e time são obrigatórios"})
		return
	}

	if !services.ValidarEmail(t.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email inválido"})
		return
	}

	novoTorcedor, err := services.SalvarTorcedor(t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       novoTorcedor.ID,
		"nome":     novoTorcedor.Nome,
		"email":    novoTorcedor.Email,
		"time":     novoTorcedor.Time,
		"mensagem": "Cadastro realizado com sucesso",
	})
}
