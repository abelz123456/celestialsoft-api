DIR = ./...

run:
	@echo "> swag init"
	@swag init
	@echo "> air"
	@air

runtest:
	@echo "> go test -v -coverprofile cover.out $(DIR)"
	@go test -v -coverprofile cover.out $(DIR)
	@echo "> go tool cover -html=cover.out -o cover.html"
	@go tool cover -html=cover.out -o cover.html
	@echo "> open cover.html"
	@open cover.html

migration:
	@echo "> go run cmd/db_migration.go"
	@go run cmd/db_migration.go