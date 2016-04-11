build:
	cd cmd/flux && go build

build-linux:
	cd cmd/flux && GOOS=linux GOARCH=amd64 go build

deps:
	go get -t -v ./...

vet:
	go vet ./...

test:
	go test ./...

bench:
	go test ./... -bench=.

.PHONY: vet test bench
