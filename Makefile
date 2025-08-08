SRC_DIR := .
BUILD_DIR := build
CONFIG_DIR := config

MAIN_PKG := $(SRC_DIR)/server

.PHONY: build
build: dependencies
	@go build -o=$(BUILD_DIR)/apiserver $(MAIN_PKG)

.PHONY: dev
dev: build
	@API_ENV=$(CONFIG_DIR)/dev.yml $(BUILD_DIR)/apiserver

.PHONY: prod
prod: build
	@API_ENV=$(CONFIG_DIR)/prod.yml $(BUILD_DIR)/apiserver

.PHONY: dependencies
dependencies:
	@go get ./...
