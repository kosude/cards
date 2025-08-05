SRC_DIR := .
BUILD_DIR := build

MAIN_PKG := $(SRC_DIR)/cmd/server

.PHONY: build
build:
	@go build -o=$(BUILD_DIR)/apiserver $(MAIN_PKG)

.PHONY: run
run: build
	@$(BUILD_DIR)/apiserver

.PHONY: dependencies
dependencies:
	@go get ./...
