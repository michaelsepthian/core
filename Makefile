SHELL=/bin/bash
GO := $(shell command -v go 2> /dev/null)

.DEFAULT_GOAL := run

check-go:
ifndef GO
	$(error "go is not installed! Aborting")
endif

run: check-go
	$(GO) run -ldflags "-X main.buildTime=$(shell date -u '+%Y-%m-%dT%T%z')" main.go

build-production: check-go
	$(GO) build -v -ldflags "-X main.buildTime=$(shell date -u '+%Y-%m-%dT%T%z') -s -w" -o server .
