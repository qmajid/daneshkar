# BINARY := myapp
# PKG := ./...

.PHONY: run swag test mock

run:
	go run main.go serve --port 8081

test:
	go test ./...

mock:
	mockgen -source=database/json/interface.go -destination=mocks/recipes_service_mock.go -package=mocks

coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o docs/coverage.html


swag:
	swag init --generalInfo main.go --output docs