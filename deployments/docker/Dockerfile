FROM golang:1.25.1-alpine AS builder

# Instalar dependências para CGO e SQLite
RUN apk add --no-cache gcc musl-dev sqlite-dev

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

# Habilitar CGO para SQLite
ENV CGO_ENABLED=1
RUN go build -o server cmd/server/main.go

FROM alpine:latest

# Instalar SQLite no container final
RUN apk add --no-cache sqlite

COPY --from=builder /app/server /usr/local/bin/server
COPY --from=builder /app/.env /app/.env

# Criar diretório para dados
RUN mkdir -p /app/data

WORKDIR /app

EXPOSE 8080

CMD ["server"]