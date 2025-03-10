.PHONY: build build-unix build-win migrate-up migrate-down

build: build-unix build-win

build-unix:
	@go build -o ./dist/report cmd/report/main.go

build-win:
	@GOOS=windows go build -o ./dist/report.exe cmd/report/main.go

migrate-up:
	@docker compose up -d
	@go generate ./...

migrate-down:
	@docker compose down
