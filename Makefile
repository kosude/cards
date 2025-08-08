SRC_DIR := .
BUILD_DIR := build
CONFIG_DIR := config

MAIN_PKG := $(SRC_DIR)/server

GO := go
AIR := air # As in https://github.com/air-verse/air -- needed for `make dev`

.PHONY: build
build: dependencies
	@$(GO) build -o=$(BUILD_DIR)/apiserver $(MAIN_PKG)

.PHONY: dev
dev: build
	@API_ENV=$(CONFIG_DIR)/dev.yml $(AIR) \
		--build.cmd "make build" \
		--build.bin "$(BUILD_DIR)/apiserver"

.PHONY: prod
prod: build
	@API_ENV=$(CONFIG_DIR)/prod.yml $(BUILD_DIR)/apiserver

.PHONY: dependencies
dependencies:
	@$(GO) get ./...
