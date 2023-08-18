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
test:
	go test -v ./...

.PHONY: clean
clean:
	rm -rf kasa kasa.exe

.PHONY: readme
readme: build
	go run tool/gen_readme.go

.PHONY: lint
lint:
	golangci-lint run
