BASE_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
BIN ?= $(BASE_DIR)/bin

.PHONY: build-api-server
build-api-server:
	cd $(BASE_DIR)/server/api &&  go build -o $(BIN)/api-server

