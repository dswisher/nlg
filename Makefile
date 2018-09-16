
# Basic go commands
GOCMD   := go
GOBUILD := $(GOCMD) build
GOTEST  := $(GOCMD) test

# Other tools
BINDIR  := $(GOPATH)/bin
GOLINT  := $(BINDIR)/golint

$(GOLINT):
	go get github.com/golang/lint/golint


# Specifics about this project
PKGS := $(shell go list ./... | grep -v /vendor)

# Building
.PHONY: build
build:
	$(GOBUILD) $(PKGS)

# Testing
.PHONY: test
test: lint
	$(GOTEST) -v -cover $(PKGS)

# Linting
.PHONY: lint
lint: $(GOLINT)
	$(GOLINT) $(PKGS)
