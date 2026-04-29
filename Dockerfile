# Estágio de compilação
FROM golang:latest AS builder

WORKDIR /app

# Copia os arquivos de dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia o restante do código
COPY . .

# Compila o binário (estático para rodar no alpine)
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api/main.go

# Estágio final (imagem leve para rodar)
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .

# Expõe a porta da API
EXPOSE 8080

# Comando para rodar a aplicação
CMD ["./main"]