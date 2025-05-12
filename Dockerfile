# Development stage with hot reload
FROM golang:1.23 AS dev
WORKDIR /app

# Module download
COPY go.mod go.sum ./
RUN go mod download

# Install Air CLI via official script (sağlıklı ve güncel kurulum)
RUN apt-get update && apt-get install -y curl git bash \
    && rm -rf /var/lib/apt/lists/*

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    | bash -s latest \
    && mv ./bin/air /usr/local/bin/air

# Install migrate CLI for dev
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz \
    | tar xz \
    && mv migrate /usr/local/bin/migrate

# Uygulama kodunu kopyala
COPY . .

EXPOSE 8080

# İlk önce migration’ları uygula, sonra Air ile live-reload başlat
CMD ["sh", "-c", "migrate -path ./migrations -database \"$DATABASE_DSN\" up && air -c .air.toml"]


# Build stage for production binary
FROM golang:1.23 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Production stage
FROM alpine:latest AS prod
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080

# Prod’da migration’lar zaten main() içinde çalıştırılıyor, doğrudan binary’i ayağa kaldır.
CMD ["./main"]