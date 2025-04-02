# Executar projeto go champ

```
GoChamp/
├── cmd/
│   └── main.go
├── internal/
│   ├── auth/
│   │   ├── jwt.go
│   │   └── middleware.go
│   ├── config/
│   │   └── config.go
│   ├── handlers/
│   │   ├── auth.go
│   │   ├── campeonato.go
│   │   ├── partida.go
│   │   ├── torcedor.go
│   │   └── broadcast.go
│   ├── models/
│   │   └── models.go
│   ├── services/
│   │   ├── external_api.go
│   │   └── broadcast_service.go
│   └── tests/
│       ├── auth_test.go
│       ├── campeonato_test.go
│       └── partida_test.go
├── go.mod
├── go.sum
├── Dockerfile
└── README.md
```
```
SSH: git clone git@github.com:tthiagosantos/go-champ.git
```

```
docker compose up -d
```

```
make migrate-up
```

```
docker-compose logs -f go_champ
```

## Documentacao
```
- Dentro da pasta doc/api.http

### LOGIN
POST http://localhost:8080/auth/login
Content-Type: application/json

{
  "usuario": "admin",
  "senha": "123456"
}

### Listar Campeonatos (JWT obrigatório)
GET http://localhost:8080/campeonatos
Content-Type: application/json
Authorization: Bearer <TOKEN>

### Listar Partidas (JWT obrigatório)
GET http://localhost:8080/campeonatos/2021/partidas?equipe=Flamengo&rodada=3
Content-Type: application/json
Authorization: Bearer <TOKEN>


###Cadastrar Torcedor (JWT obrigatório)
POST http://localhost:8080/torcedores
Content-Type: application/json
Authorization: Bearer <TOKEN>

{
  "nome": "João Silva",
  "email": "joao@example.com",
  "time": "Flamengo"
}

###Broadcast (JWT obrigatório)
POST http://localhost:8080/broadcast
Content-Type: application/json
Authorization: Bearer <TOKEN>

{
  "tipo": "inicio",
  "time": "Flamengo",
  "mensagem": "O jogo do Flamengo vai começar!"
}

###Broadcast (JWT obrigatório)
POST http://localhost:8080/broadcast
Content-Type: application/json
Authorization: Bearer <TOKEN>

{
  "tipo": "fim",
  "time": "Flamengo",
  "placar": "2-1",
  "mensagem": "Acabou o jogo!"
} 
```

```
🛠️ Tecnologias e Ferramentas recomendadas:
	•	Gin como framework web.
	•	PostgreSQL como banco de dados.
	•	JWT para autenticação.
	•	Docker para containerização.
	•	sql-migrate ou golang-migrate para migração de banco.
	•	Testify para testes unitários e integração.
	•	Clean Architecture para a estruturação da lógica do projeto.
```

### Evelucao do projeto

	•	Integrar com banco de dados real (PostgreSQL) caso queira persistir usuários, torcedores e logs.
	•	Melhorar os testes, especialmente os de integração (usando httptest e rotas reais).
	•	Implementar logs estruturados, tratamento de erros avançado, etc.
	•	Ajustar o Docker e use docker-compose caso precise de banco adicional.


