## ----------------------------------------------------------------------
## Author: Siddhartha Chowdhuri
## This makefile is used to test code locally
## Please report any bug or error to the author
## ----------------------------------------------------------------------
tests:
	@echo Running tests locally
	go get -v -t -d ./...
	go test -v ./... --coverprofile coverage.out

server:
	@echo Spinning up back-end servers
	go run cmd/servers/main.go -s=5

all: tests server
help:
	@echo ' 						   '
	@echo  '  Usage: make [target] ... '
	@echo ' 						   '
	@echo  '  tests         - Runs all the unit tests'
	@echo  '  server        - Runs the back-end servers'
	@echo  '  all           - Runs tests and the entire pipeline'
	@echo  ''
