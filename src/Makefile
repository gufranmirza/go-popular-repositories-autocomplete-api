export BINARY_NAME = "go-popular"
export CONFIG_FILE = "app-config.json"

# global command
.PHONY: all
all: dependencies build

.PHONY: dependencies
dependencies:
	export GO111MODULE=on
	go mod download

.PHONY: build
build:
	env GOOS=linux GOARCH=amd64 go build -o ../bin/${BINARY_NAME} .
	cp -r ${CONFIG_FILE} ../bin

.PHONY: run
run:
	go run ./main.go serve

.PHONY: start
start:
	../bin/${BINARY_NAME} serve

.PHONY: dataloader
dataloader:
	../bin/${BINARY_NAME} dataloader


api-docs:
	export PATH=$(go env GOPATH)/bin:$PATH
	swag init --output=./web/docs

migrate-up:
	../bin/${BINARY_NAME} migrate --action=up

migrate-down:
	../bin/${BINARY_NAME} migrate --action=down


.PHONY: mock
mock:	
	mockgen -package servermock github.com/gufranmirza/go-popular-repositories-autocomplete-api/web/server Server > web/server/servermock/mock.go
	mockgen -package routermock github.com/gufranmirza/go-popular-repositories-autocomplete-api/web/router Router > web/router/routermock/mock.go
	mockgen -package healthmock github.com/gufranmirza/go-popular-repositories-autocomplete-api/web/services/health Health > web/services/health/healthmock/mock.go
	mockgen -package repositoryservicemock github.com/gufranmirza/go-popular-repositories-autocomplete-api/web/services/v1/repositoryservice RepositoryService > web/services/v1/repositoryservice/repositoryservicemock/mock.go
	mockgen -package middlewaresmock github.com/gufranmirza/go-popular-repositories-autocomplete-api/web/middlewares Middlewares > web/middlewares/middlewaresmock/mock.go
	mockgen -package middlewaresmock github.com/gufranmirza/go-popular-repositories-autocomplete-api/dal/repositorydal RepositoryDal > dal/repositorydal/repositorydalmock/mock.go
	mockgen -package clientsmock github.com/gufranmirza/go-popular-repositories-autocomplete-api/clients Clients > clients/clientsmock/mock.go

packages = \
	./web/server\
	./web/router\
	./web/services/health\
	./web/services/v1/repositoryservice\
	./web/middlewares\
	./dal/repositorydal\
	./clients\


.PHONY: test
test:
	@$(foreach package,$(packages), \
    		set -e; \
    		go test  $(package);)