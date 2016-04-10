deps:
	go get -t -v ./...

vet:
	go vet ./...

test:
	go test ./...

bench:
	go test ./... -bench=.

.PHONY: vet test bench
