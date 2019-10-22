BASE_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
BIN ?= $(BASE_DIR)/bin

.PHONY: build-api-server
build-api-server:
	cd $(BASE_DIR)/server/api &&  go build -o $(BIN)/api-server

.PHONY: build-stream-server
build-stream-server:
	cd $(BASE_DIR)/server/stream &&  go build -o $(BIN)/stream-server

.PHONY: build-scheduler-server
build-scheduler-server:
	cd $(BASE_DIR)/server/scheduler &&  go build -o $(BIN)/scheduler-server