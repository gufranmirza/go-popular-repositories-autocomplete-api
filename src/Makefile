api-docs:
	export PATH=$(go env GOPATH)/bin:$PATH
	swag init --output=./web/docs

migrate-up:
	go run main.go migrate --action=up

migrate-down:
	go run main.go migrate --action=down