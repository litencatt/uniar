air-cmd:
	go mod tidy
	sqlc generate
	go build -o ./tmp/main ./cmd/uniar/main.go

build:
	docker build -t litencatt/uniar:latest  --build-arg GITHUB_COM_TOKEN=$$GITHUB_COM_TOKEN .

migrate:
	sqlite3def uniar.db -f sql/schema.sql
