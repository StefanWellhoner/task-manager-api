# The name of your application
APP_NAME = task-forge

# The go compiler to use
GO = go

build:
	$(GO) build -o $(APP_NAME) main.go

run: build
	@./$(APP_NAME)

test:
	$(GO) test -v ./...

coverage:
	$(GO) test --race -coverprofile=coverage.out -covermode=atomic ./...

clean:
	rm -f $(APP_NAME) coverage.out

.PHONY: build run test clean coverage