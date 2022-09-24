PKG = github.com/litencatt/uniar
COMMIT = $$(git describe --tags --always)
OSNAME=${shell uname -s}
ifeq ($(OSNAME),Darwin)
	DATE = $$(gdate --utc '+%Y-%m-%d_%H:%M:%S')
else
	DATE = $$(date --utc '+%Y-%m-%d_%H:%M:%S')
endif

export GO111MODULE=on

BUILD_LDFLAGS = -X $(PKG).commit=$(COMMIT) -X $(PKG).date=$(DATE)
UNIAR_BINARY ?= ./uniar

build:
	go build -ldflags="$(BUILD_LDFLAGS)" -o $(UNIAR_BINARY) cmd/uniar/main.go

air-cmd:
	go mod tidy
	sqlc generate
	go build -o ./tmp/main ./cmd/uniar/main.go

docker-build:
	docker build -t litencatt/uniar:latest  --build-arg GITHUB_COM_TOKEN=$$GITHUB_COM_TOKEN .

db-migrate:
	sqlite3def uniar.db -f sql/schema.sql
