package handlers

import (
	"github.com/tthiagosantos/gochamp/internal/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Usuario string `json:"usuario"`
	Senha   string `json:"senha"`
}

func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}
	
	if req.Usuario == "admin" && req.Senha == "123456" {
		token, err := auth.GeraTokenJWT(req.Usuario)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
	}
}
