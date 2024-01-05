# Makefile for building the Go application

# Application name
APP_NAME := duplicateremover

# Go related variables
GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin
GOFILES := $(wildcard *.go)

# Go environment variables
export GO111MODULE=on

# Build variables
BUILD_ENV := CGO_ENABLED=0 GOOS=windows GOARCH=amd64

.PHONY: all build clean

all: build

build:
	@echo "Building for amd64 architecture..."
	$(BUILD_ENV) go build -o $(GOBIN)/$(APP_NAME).exe $(GOFILES)

clean:
	@echo "Cleaning up..."
	@rm -f $(GOBIN)/$(APP_NAME)
