package tests

import (
	"github.com/tthiagosantos/gochamp/internal/services"
	"testing"
)

func TestBuscarPartidas(t *testing.T) {
	_, err := services.BuscarPartidas("2021", "", "")
	if err != nil {
		t.Fatalf("Erro ao buscar partidas: %v", err)
	}
}
