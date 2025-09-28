FROM golang:1.23 AS dev

WORKDIR /app

RUN go install github.com/air-verse/air@v1.52.3
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
RUN go install github.com/golang/mock/mockgen@latest
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
RUN wget https://github.com/k0kubun/sqldef/releases/latest/download/sqlite3def_linux_amd64.tar.gz && \
      tar xvzf sqlite3def_linux_amd64.tar.gz && \
      mv sqlite3def /go/bin/.

CMD ["air"]
