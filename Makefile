
all: build
	@./bin/api

build:
	@go build -o ./bin/api ./cmd/api/*

.PHONY: all, build
