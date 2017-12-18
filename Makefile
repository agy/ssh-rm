GOARCH ?= amd64


build: vet fmt
	go $@ .

linux: vet fmt
	GOOS=$@ GOARCH=$(GOARCH) go build .

darwin: vet fmt
	GOOS=$@ GOARCH=$(GOARCH) go build .

%:
	go $@ ./...


.PHONY: linux darwin build
