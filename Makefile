.DEFAULT_GOAL := all
BUILD_DATE := $(shell date +%Y-%m-%d\ %H:%M)
BUILD_TAG := $(shell git describe --abbrev=0 --tags)
LD_FLAGS := '-X "co-pilot/cmd.buildDate=$(BUILD_DATE)" -X "co-pilot/cmd.version=$(BUILD_TAG)"'


build:
	go build -ldflags ${LD_FLAGS}

install:
	go install -ldflags ${LD_FLAGS}

run:
	go run main.go

test:
	go test -v -cover ./...

lint:
	gofmt -w pkg
	gofmt -w cmd

all: build install
