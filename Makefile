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

sqldef-update:
	curl -OL https://github.com/k0kubun/sqldef/releases/download/v0.15.6/sqlite3def_darwin_arm64.zip
	sudo tar xf sqlite3def_darwin_arm64.zip -C /usr/local/bin

build:
	go build -ldflags="$(BUILD_LDFLAGS)" -o $(UNIAR_BINARY) cmd/uniar/main.go

air-cmd:
	go mod tidy
	sqlc generate
	@$(MAKE) gen-mock
	go build -o ./tmp/main ./cmd/uniar/main.go

docker-build:
	docker build -t litencatt/uniar:latest  --build-arg GITHUB_COM_TOKEN=$$GITHUB_COM_TOKEN .

db-init:
	rm ~/.uniar/uniar.db

db-migrate:
	sqlite3def ~/.uniar/uniar.db -f sql/schema.sql

db-dump:
	sqlite3 ~/.uniar/uniar.db '.dump "center_skills"' | grep ^INSERT  > sql/seed.sql
	sqlite3 ~/.uniar/uniar.db '.dump "color_types"'   | grep ^INSERT >> sql/seed.sql
	sqlite3 ~/.uniar/uniar.db '.dump "groups"'        | grep ^INSERT >> sql/seed.sql
	sqlite3 ~/.uniar/uniar.db '.dump "lives"'         | grep ^INSERT >> sql/seed.sql
	sqlite3 ~/.uniar/uniar.db '.dump "members"'       | grep ^INSERT >> sql/seed.sql
	sqlite3 ~/.uniar/uniar.db '.dump "music"'         | grep ^INSERT >> sql/seed.sql
	sqlite3 ~/.uniar/uniar.db '.dump "photograph"'    | grep ^INSERT >> sql/seed.sql
	sqlite3 ~/.uniar/uniar.db '.dump "photo_types"'   | grep ^INSERT >> sql/seed.sql
	sqlite3 ~/.uniar/uniar.db '.dump "scenes"'        | grep ^INSERT >> sql/seed.sql
	sqlite3 ~/.uniar/uniar.db '.dump "skills"'        | grep ^INSERT >> sql/seed.sql
	sed -i '' -e "s/INSERT INTO/INSERT OR REPLACE INTO/g" sql/seed.sql

doc: build
	go run cmd/uniar/main.go doc > README.md

gen-mock:
	mockgen -source repository/querier.go -destination repository/querier_mock.go -package repository

prerelease:
	@$(MAKE) db-dump
	go mod tidy
	@$(MAKE) doc
	ghch -w -A --format=markdown -N $(NEXT_VER)
	gocredits -skip-missing -w .

release:
	git tag ${NEXT_VER}
	git push origin main --tag
	goreleaser --rm-dist
