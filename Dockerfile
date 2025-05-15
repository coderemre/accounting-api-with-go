FROM golang:1.23 AS dev
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN apt-get update && apt-get install -y curl git bash \
    && rm -rf /var/lib/apt/lists/*

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    | bash -s latest \
    && mv ./bin/air /usr/local/bin/air

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz \
    | tar xz \
    && mv migrate /usr/local/bin/migrate

COPY . .

EXPOSE 8080

CMD ["sh", "-c", "migrate -path ./migrations -database \"$DATABASE_DSN\" up && air -c .air.toml"]


FROM golang:1.23 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest AS prod
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080

CMD ["./main"]