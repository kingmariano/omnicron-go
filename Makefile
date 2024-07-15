# Makefile for Go project

# Variables
GOCMD = go
GOTEST = $(GOCMD) test
GOFMT = $(GOCMD) fmt

test:
	$(GOTEST) -v ./...

fmt:
	$(GOFMT) ./...
