# Makefile for Go project

# Variables
GOCMD = go
GOFMT =  gofmt
GOTEST = $(GOCMD) test

test:
	$(GOTEST) -v ./...

fmt:
	$(GOFMT) -s -w .
