# Makefile para RecipeHub

.PHONY: help run build clean test dev 

help:
	@echo "Usage: make <target>"
	@echo "Targets:"
	@echo "  run - Executar em modo desenvolvimento"
	@echo "  dev - Executar em modo desenvolvimento com hot reload (requer air)"
	@echo "  build - Compilar para produção"
	@echo "  clean - Limpar arquivos compilados"
	@echo "  test - Executar testes"
	@echo "  deps - Instalar dependências"
	@echo "  run-prod - Executar em modo produção"
	@echo "  docker-build - Compilar para produção em container Docker"
	@echo "  docker-run - Executar container Docker"

# Executar em modo desenvolvimento
run:
	go run cmd/server/main.go

# Compilar para produção
build:
	go build -o bin/server cmd/server/main.go

# Limpar arquivos compilados
clean:
	rm -rf bin/

# Executar testes
test:
	go test ./...

# Instalar dependências
deps:
	go mod tidy
	go mod download

# Docker (opcional)
docker-build:
	docker build -f deployments/docker/Dockerfile -t iriscafe50/recipehub .

docker-run:
	docker run -p 8080:8080 recipehub

docker-stop:
	docker stop recipehub

docker-remove:
	docker rm recipehub

docker-logs:
	docker logs recipehub
