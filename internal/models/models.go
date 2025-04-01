package models

type Campeonato struct {
	ID        string `json:"id"`
	Nome      string `json:"nome"`
	Temporada string `json:"temporada"`
}

type Partida struct {
	Rodada   int    `json:"rodada"`
	TimeCasa string `json:"time_casa"`
	TimeFora string `json:"time_fora"`
	Placar   string `json:"placar"`
}

type PartidaResponse struct {
	Rodada   int       `json:"rodada"`
	Partidas []Partida `json:"partidas"`
}

type Torcedor struct {
	ID    int    `json:"id,omitempty"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Time  string `json:"time"`
}
