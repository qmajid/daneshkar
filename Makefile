# BINARY := myapp
# PKG := ./...

.PHONY: run swag

run:
	go run main.go serve --port 8081

test:
	go test ./...

mock:
	mockgen -source=database/json/interface.go -destination=mocks/recipes_service_mock.go -package=mocks

swag:
	swag init --generalInfo main.go --output docs