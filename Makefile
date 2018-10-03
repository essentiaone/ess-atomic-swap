
GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_NAME=ess-atomic-swap

run:	build start

build:
	$(GOBUILD) .

start:
	./$(BINARY_NAME)
