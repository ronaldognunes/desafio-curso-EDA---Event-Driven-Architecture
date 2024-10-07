FROM golang:1.19

WORKDIR /app

RUN apt-get update && apt-get install -y librdkafka-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN GOOS=linux go build -o api-core ./cmd/walletcore/main.go

EXPOSE 8080

CMD ["./api-core"]