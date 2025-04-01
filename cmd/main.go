package main

import (
	"github.com/tthiagosantos/gochamp/internal/auth"
	"github.com/tthiagosantos/gochamp/internal/handlers"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Rota pública (login)
	r.POST("/auth/login", handlers.Login)

	// Rota pública opcional (cadastro de usuário) - não implementado em detalhes
	// r.POST("/auth/register", handlers.Register)

	protected := r.Group("/")
	protected.Use(auth.JWTAuthMiddleware())

	protected.GET("/campeonatos", handlers.ListarCampeonatos)
	protected.GET("/campeonatos/:id/partidas", handlers.ListarPartidas)

	// Cadastro de torcedores
	protected.POST("/torcedores", handlers.CadastrarTorcedor)

	// Broadcast
	protected.POST("/broadcast", handlers.Broadcast)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	log.Printf("Servidor rodando na porta %s", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}
