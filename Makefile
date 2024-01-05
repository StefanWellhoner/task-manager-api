.PHONY: build run test clean

# The name of your application
APP_NAME = task-manager-api

# The go compiler to use
GO = go

build:
	$(GO) build -o $(APP_NAME) main.go

run: build
	@./$(APP_NAME)

test:
	$(GO) test -v ./...

clean:
	rm -f $(APP_NAME)
