package repositories

import (
	"github.com/tthiagosantos/gochamp/internal/infrastructure/database"
	"github.com/tthiagosantos/gochamp/internal/models"
)

type torcedorRepository struct{}

func NewTorcedorRepository() TorcedorRepository {
	return &torcedorRepository{}
}

func (r *torcedorRepository) Salvar(t models.Torcedor) (models.Torcedor, error) {
	query := `INSERT INTO torcedores (nome, email, time) VALUES ($1, $2, $3) RETURNING id`
	err := database.DB.QueryRow(query, t.Nome, t.Email, t.Time).Scan(&t.ID)
	if err != nil {
		return t, err
	}
	return t, nil
}

func (r *torcedorRepository) BuscarPorTime(time string) ([]models.Torcedor, error) {
	query := `SELECT id, nome, email, time FROM torcedores WHERE time = $1`
	rows, err := database.DB.Query(query, time)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var torcedores []models.Torcedor
	for rows.Next() {
		var t models.Torcedor
		if err := rows.Scan(&t.ID, &t.Nome, &t.Email, &t.Time); err != nil {
			return nil, err
		}
		torcedores = append(torcedores, t)
	}
	return torcedores, nil
}
