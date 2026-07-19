BUILD_FLAGS := -ldflags="-w -s"
BUILD_DIR := build
BIN_NAME := aether

.PHONY: build

build:
	go build -o $(BUILD_DIR)/$(BIN_NAME) $(BUILD_FLAGS)
