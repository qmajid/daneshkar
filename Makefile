# BINARY := myapp
# PKG := ./...

.PHONY: run swag

run:
	go run main.go

swag:
	swag init --generalInfo main.go --output docs