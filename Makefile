GOOS := linux

all: vet fmt build

vet:
	go vet .

fmt:
	go fmt .

build:
	go build .

linux:
	GOOS=$(GOOS) go build .
