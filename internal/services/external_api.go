package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tthiagosantos/gochamp/internal/models"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var baseURL = os.Getenv("BASE_URL_API")

var apiToken = os.Getenv("API_KEY")

type CompetitionResponse struct {
	Competitions []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Area struct {
			Name string `json:"name"`
		} `json:"area"`
		Code        string `json:"code"`
		Plan        string `json:"plan"`
		LastUpdated string `json:"lastUpdated"`
	} `json:"competitions"`
}

type MatchesResponse struct {
	Matches []struct {
		Matchday int `json:"matchday"`
		HomeTeam struct {
			Name string `json:"name"`
		} `json:"homeTeam"`
		AwayTeam struct {
			Name string `json:"name"`
		} `json:"awayTeam"`
		Score struct {
			FullTime struct {
				Home int `json:"home"`
				Away int `json:"away"`
			} `json:"fullTime"`
		} `json:"score"`
	} `json:"matches"`
}

func BuscarCampeonatos() ([]models.Campeonato, error) {
	url := baseURL + "/competitions/"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Auth-Token", apiToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("falha ao buscar competições na API externa")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var compResp CompetitionResponse
	if err := json.Unmarshal(body, &compResp); err != nil {
		return nil, err
	}

	var campeonatos []models.Campeonato
	for _, c := range compResp.Competitions {
		camp := models.Campeonato{
			ID:        fmt.Sprintf("%d", c.ID),
			Nome:      c.Name,
			Temporada: "2025",
		}
		campeonatos = append(campeonatos, camp)
	}

	return campeonatos, nil
}

func BuscarPartidas(idCampeonato string, equipe string, rodada string) ([]models.PartidaResponse, error) {
	url := baseURL + "/competitions/" + idCampeonato + "/matches"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Auth-Token", apiToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("falha ao buscar partidas na API externa, status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var mResp MatchesResponse
	if err := json.Unmarshal(body, &mResp); err != nil {
		return nil, err
	}

	var rodadaInt int
	if rodada != "" {
		if rodadaInt, err = strconv.Atoi(rodada); err != nil {
			return nil, errors.New("rodada inválida")
		}
	}

	partidasPorRodada := make(map[int][]models.Partida)

	for _, m := range mResp.Matches {
		p := models.Partida{
			Rodada:   m.Matchday,
			TimeCasa: m.HomeTeam.Name,
			TimeFora: m.AwayTeam.Name,
			Placar: fmt.Sprintf("%d-%d",
				m.Score.FullTime.Home,
				m.Score.FullTime.Away),
		}

		if equipe != "" {
			eqLower := strings.ToLower(equipe)
			casaLower := strings.ToLower(p.TimeCasa)
			foraLower := strings.ToLower(p.TimeFora)
			if !strings.Contains(casaLower, eqLower) && !strings.Contains(foraLower, eqLower) {
				continue
			}
		}

		if rodadaInt != 0 && p.Rodada != rodadaInt {
			continue
		}

		partidasPorRodada[p.Rodada] = append(partidasPorRodada[p.Rodada], p)
	}

	var resultado []models.PartidaResponse
	for rodadaNum, lista := range partidasPorRodada {
		resultado = append(resultado, models.PartidaResponse{
			Rodada:   rodadaNum,
			Partidas: lista,
		})
	}

	return resultado, nil
}
