FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o /app/main ./cmd

FROM scratch
COPY --from=builder /app/main /main
ENTRYPOINT ["/main"]