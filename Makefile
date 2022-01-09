GOOS    := $(shell go env GOOS)
GOARCH  := $(shell go env GOARCH)

.PHONY: all
all: vet test build

.PHONY: build
build:
	go build ./cmd/kasa

.PHONY: vet
vet:
	go vet ./...

.PHONY: test
vet:
	go test ./...

.PHONY: clean
clean:
	rm -rf kasa kasa.exe
