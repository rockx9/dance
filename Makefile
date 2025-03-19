# Go parameters
export PATH:=${PATH}:${GOPATH}/bin
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
BINARY_DIR=outputs
SWAG=swag init
LINUX_AMD64= GOOS=linux GOARCH=amd64

run-dev:
	$(SWAG)
	$(GORUN) main.go
run-test:
build: