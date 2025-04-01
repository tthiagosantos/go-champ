package services

import (
	"fmt"
	"github.com/tthiagosantos/gochamp/internal/models"
	"github.com/tthiagosantos/gochamp/internal/repositories"
	"net/mail"
)

var repo = repositories.NewTorcedorRepository()

func SalvarTorcedor(t models.Torcedor) (models.Torcedor, error) {
	if _, err := mail.ParseAddress(t.Email); err != nil {
		return models.Torcedor{}, fmt.Errorf("email inválido: %s", t.Email)
	}

	return repo.Salvar(t)
}
