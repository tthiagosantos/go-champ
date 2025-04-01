package tests

import (
	"testing"

	"github.com/tthiagosantos/gochamp/internal/auth"
)

func TestGeraValidaTokenJWT(t *testing.T) {
	usuario := "admin"
	token, err := auth.GeraTokenJWT(usuario)
	if err != nil {
		t.Fatalf("Erro ao gerar token: %v", err)
	}

	claims, err := auth.ValidaTokenJWT(token)
	if err != nil {
		t.Fatalf("Erro ao validar token: %v", err)
	}

	if claims["sub"] != usuario {
		t.Errorf("Esperava sub = %s, obteve %v", usuario, claims["sub"])
	}
}
