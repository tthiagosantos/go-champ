package database

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func InitPostgres() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar .env")
	}
	dsn := os.Getenv("DATABASE_URL")
	log.Printf("DATABASE_URL: %s", dsn)
	if dsn == "" {
		log.Fatal("DATABASE_URL não configurada")
	}

	DB, err = sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("Erro ao abrir conexão com banco: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Erro ao pingar banco de dados: %v", err)
	}

	fmt.Println("✅ Banco de dados conectado com sucesso!")
}
