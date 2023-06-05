build:
	go build -o cmd/main main.go

run:
	docker-compose up && go run main.go

fmt:
	@gofmt -l -w .