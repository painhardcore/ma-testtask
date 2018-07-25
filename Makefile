# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOVET=$(GOCMD) vet
GOGET=$(GOCMD) get
BINFOLDER = ./bin/
BINARY_NAME=searcher
BINARY_UNIX=$(BINARY_NAME)_linux

all: test vet build
build: 
	$(GOBUILD) -o $(BINFOLDER)$(BINARY_NAME) -v
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINFOLDER)$(BINARY_NAME)
	rm -f $(BINFOLDER)$(BINARY_UNIX)
run: vet build
	echo 'https://golang.org\nhttps://en.wikipedia.org/wiki/Go_\(programming_language\)\nhttps://github.com/golang/go\nhttps://medium.com/exploring-code/why-should-you-learn-go-f607681fad65' | $(BINFOLDER)$(BINARY_NAME)
vet:
	$(GOVET) -v ./...

# Cross compilation
build-linux: vet
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINFOLDER)$(BINARY_UNIX) -v