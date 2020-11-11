# Auto generated binary variables helper managed by https://github.com/bwplotka/bingo v0.2.2. DO NOT EDIT.
# All tools are designed to be build inside $GOBIN.
GOPATH ?= $(shell go env GOPATH)
GOBIN  ?= $(firstword $(subst :, ,${GOPATH}))/bin
GO     ?= $(shell which go)

# Bellow generated variables ensure that every time a tool under each variable is invoked, the correct version
# will be used; reinstalling only if needed.
# For example for abigen variable:
#
# In your main Makefile (for non array binaries):
#
#include .bingo/Variables.mk # Assuming -dir was set to .bingo .
#
#command: $(ABIGEN)
#	@echo "Running abigen"
#	@$(ABIGEN) <flags/args..>
#
ABIGEN := $(GOBIN)/abigen-v1.9.23
$(ABIGEN): .bingo/abigen.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/abigen-v1.9.23"
	@cd .bingo && $(GO) build -modfile=abigen.mod -o=$(GOBIN)/abigen-v1.9.23 "github.com/ethereum/go-ethereum/cmd/abigen"

FAILLINT := $(GOBIN)/faillint-v1.5.0
$(FAILLINT): .bingo/faillint.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/faillint-v1.5.0"
	@cd .bingo && $(GO) build -modfile=faillint.mod -o=$(GOBIN)/faillint-v1.5.0 "github.com/fatih/faillint"

GOIMPORTS := $(GOBIN)/goimports-v0.0.0-20200925163029-4c77dbd9e533
$(GOIMPORTS): .bingo/goimports.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/goimports-v0.0.0-20200925163029-4c77dbd9e533"
	@cd .bingo && $(GO) build -modfile=goimports.mod -o=$(GOBIN)/goimports-v0.0.0-20200925163029-4c77dbd9e533 "golang.org/x/tools/cmd/goimports"

GOLANGCI_LINT := $(GOBIN)/golangci-lint-v1.31.0
$(GOLANGCI_LINT): .bingo/golangci-lint.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/golangci-lint-v1.31.0"
	@cd .bingo && $(GO) build -modfile=golangci-lint.mod -o=$(GOBIN)/golangci-lint-v1.31.0 "github.com/golangci/golangci-lint/cmd/golangci-lint"

