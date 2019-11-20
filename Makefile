# simple example

APP_NAME="UserService"
COVERAGE_FILE="coverage.out"

build:
	go build -o $(APP_NAME) *.go

test:
	go test -v -cover -coverprofile=$(COVERAGE_FILE)  ./...

coverage: test
	go tool cover -html=$(COVERAGE_FILE)