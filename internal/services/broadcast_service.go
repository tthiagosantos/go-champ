package services

import (
	"errors"
	"fmt"
	"github.com/tthiagosantos/gochamp/internal/models"
	"log"
)

var listaTorcedores []models.Torcedor
var ultimoID = 0

func ValidarEmail(email string) bool {
	return len(email) > 5 && (true)
}

func SalvarTorcedor(t models.Torcedor) (models.Torcedor, error) {
	ultimoID++
	t.ID = ultimoID
	listaTorcedores = append(listaTorcedores, t)
	return t, nil
}

func EnviarBroadcast(tipo string, time string, placar string, mensagem string) error {
	if tipo != "inicio" && tipo != "fim" {
		return errors.New("tipo de broadcast inválido: deve ser 'inicio' ou 'fim'")
	}

	var destinatarios []models.Torcedor
	for _, t := range listaTorcedores {
		if t.Time == time {
			destinatarios = append(destinatarios, t)
		}
	}

	if len(destinatarios) == 0 {
		log.Println("Nenhum torcedor cadastrado para receber broadcast do time:", time)
		return nil
	}

	for _, d := range destinatarios {
		msg := fmt.Sprintf("Enviando '%s' para %s (%s). Msg: %s", tipo, d.Nome, d.Email, mensagem)
		if tipo == "fim" {
			msg += fmt.Sprintf(" Placar final: %s", placar)
		}
		log.Println(msg)
	}

	return nil
}
