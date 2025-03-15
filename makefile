.PHONY: build build-unix build-win services-up services-down

build: build-unix build-win

build-unix:
	@go build -o ./dist/report cmd/report/main.go

build-win:
	@GOOS=windows go build -o ./dist/report.exe cmd/report/main.go

services-up:
	@docker compose up -d

services-down:
	@docker compose down

migrate-up:
	@go run cmd/tools/ternup/main.go

migrate-down:
	@go run cmd/tools/terndn/main.go
