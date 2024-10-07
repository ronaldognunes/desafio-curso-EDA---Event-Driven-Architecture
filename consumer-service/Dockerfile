FROM golang:1.21

WORKDIR /app

RUN apt update && \
    apt install -y wget netcat-traditional && \
    wget -q -O /usr/bin/wait-for https://raw.githubusercontent.com/eficode/wait-for/v2.2.3/wait-for && \
    chmod +x /usr/bin/wait-for

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o api-consumer ./cmd/main.go

EXPOSE 3003


CMD ["./api-consumer"]