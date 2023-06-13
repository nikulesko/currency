build:
	go build -o cmd/main main.go

run:
	docker-compose up -d
	go run cmd/main.go

format:
	@gofmt -l -w .