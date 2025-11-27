# daneshkar
Daneshkar golang advance course

go install go.uber.org/mock/mockgen@latest
mockgen -source=database/json/interface.go -destination=mocks/recipes_service_mock.go -package=mocks