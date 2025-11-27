# BINARY := myapp
# PKG := ./...

.PHONY: run swag

run:
	go run main.go serve --port 8081

swag:
	swag init --generalInfo main.go --output docs