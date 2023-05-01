FROM golang:1.19.0 as dev

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
RUN go install github.com/golang/mock/mockgen@latest
RUN wget https://github.com/k0kubun/sqldef/releases/latest/download/sqlite3def_linux_amd64.tar.gz && \
      tar xvzf sqlite3def_linux_amd64.tar.gz && \
      mv sqlite3def /go/bin/.

CMD ["air"]
