FROM golang:1.23 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
# Copia o .env para dentro da imagem
COPY .env .env

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o /app/main ./cmd

FROM alpine:3.19
WORKDIR /app

COPY --from=builder /app/main /app/main

# Copia o .env do estágio de build
COPY --from=builder /app/.env /app/.env

RUN chmod +x /app/main
RUN apk add --no-cache netcat-openbsd

CMD sh -c "until nc -z go_db_champ 5432; do echo '🔄 Aguardando PostgreSQL...'; sleep 1; done; echo '🚀 Subindo app'; /app/main"