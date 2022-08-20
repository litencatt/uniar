FROM golang:1.19.0 as dev

ARG GITHUB_COM_TOKEN
RUN git config --global url."https://$GITHUB_COM_TOKEN:x-oauth-basic@github.com/".insteadOf "https://github.com/"

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

CMD ["air"]
