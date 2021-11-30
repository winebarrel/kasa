GOOS    := $(shell go env GOOS)
GOARCH  := $(shell go env GOARCH)

.PHONY: all
all: vet build

.PHONY: build
build:
	go build ./cmd/kasa

.PHONY: vet
vet:
	go vet ./...

.PHONY: clean
clean:
	rm -rf kasa kasa.exe
