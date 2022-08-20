FROM golang:1.19.0 as dev

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

CMD ["air"]
