package tests

import (
	"github.com/tthiagosantos/gochamp/internal/services"
	"testing"
)

func TestBuscarCampeonatos(t *testing.T) {
	_, err := services.BuscarCampeonatos()
	if err != nil {
		t.Fatalf("Erro ao buscar campeonatos: %v", err)
	}

}
