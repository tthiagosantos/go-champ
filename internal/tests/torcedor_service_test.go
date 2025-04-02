package tests

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/tthiagosantos/gochamp/internal/infrastructure/database"
	"github.com/tthiagosantos/gochamp/internal/models"
	"github.com/tthiagosantos/gochamp/internal/repositories"
	"github.com/tthiagosantos/gochamp/internal/services"
	"regexp"
	"testing"
)

func TestSalvarTorcedor_EmailInvalido(t *testing.T) {
	torcedor := models.Torcedor{
		Nome:  "João",
		Email: "invalido-sem-arroba",
		Time:  "Flamengo",
	}

	_, err := services.SalvarTorcedor(torcedor)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "email inválido")
}

func TestSalvarTorcedor_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	database.DB = db

	torcedor := models.Torcedor{
		Nome:  "Maria",
		Email: "maria@example.com",
		Time:  "Corinthians",
	}

	mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO torcedores")).
		WithArgs(torcedor.Nome, torcedor.Email, torcedor.Time).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	repo := repositories.NewTorcedorRepository()
	res, err := repo.Salvar(torcedor)

	assert.NoError(t, err)
	assert.Equal(t, 1, res.ID)
	assert.NoError(t, mock.ExpectationsWereMet())
}
