build: vet fmt
	go $@ .

linux: vet fmt
	GOOS=$@ GOARCH=amd64 go build .

darwin: vet fmt
	GOOS=$@ GOARCH=386 go build .

%:
	go $@ ./...


.PHONY: linux darwin build
