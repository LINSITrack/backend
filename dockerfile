FROM golang:1.25-alpine

WORKDIR /app

# Instalar Air para hot reload (versión compatible)
RUN go install github.com/air-verse/air@latest

# Copiar go.mod y go.sum primero (para cache de dependencias)
COPY go.mod go.sum ./
RUN go mod download

# Copiar el resto del código
COPY . .

# Exponer puerto
EXPOSE 8080

# Usar Air para desarrollo con hot reload
CMD ["air"]