# The name of your application
APP_NAME = task-forge

# variables
SCRIPTS_DIR = ./scripts

build:
	@echo "Building $(APP_NAME)..."
	@${SCRIPTS_DIR}/build.sh $(APP_NAME)
	@echo "Done"

run: 
	@echo "Running $(APP_NAME)..."
	@${SCRIPTS_DIR}/run.sh

test:
	@echo "Running tests..."
	@${SCRIPTS_DIR}/test.sh

coverage:
	@echo "Running coverage..."
	@${SCRIPTS_DIR}/coverage.sh

clean:
	@echo "Cleaning $(APP_NAME)..."
	@${SCRIPTS_DIR}/clean.sh $(APP_NAME)

.PHONY: build run test clean coverage