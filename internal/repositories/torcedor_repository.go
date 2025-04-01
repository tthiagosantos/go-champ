package repositories

import "github.com/tthiagosantos/gochamp/internal/models"

type TorcedorRepository interface {
	Salvar(t models.Torcedor) (models.Torcedor, error)
	BuscarPorTime(time string) ([]models.Torcedor, error)
}
