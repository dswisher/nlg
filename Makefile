
# Basic go commands
GOCMD  := go
GOTEST := $(GOCMD) test

# Other tools
BINDIR := $(GOPATH)/bin
GOLINT := $(BINDIR)/golint

$(GOLINT):
	go get github.com/golang/lint/golint


# Specifics about this project
PKGS := $(shell go list ./... | grep -v /vendor)


# Testing
.PHONY: test
test: lint
	$(GOTEST) -v $(PKGS)

# Linting
.PHONY: lint
lint: $(GOLINT)
	$(GOLINT) $(PKGS)
