# Choose whatever you want, version >= 1.16
FROM golang:1.20-alpine

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY go.mod ./
RUN go mod download

EXPOSE 8000

CMD ["air", "-c", ".air.toml"]
