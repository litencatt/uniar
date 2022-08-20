air-cmd:
	go mod tidy
	sqlc generate
	go build -o ./tmp/main .
build:
	docker build -t litencatt/unisonair:latest .