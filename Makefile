.PHONY: all build clean default help init test format check-license
default: help

build: build-mongoman-server

build-mongoman-server: clean
	bash build.sh

gomod:
	go mod tidy
	go mod vendor

gen:


lint:
	golangci-lint run

clean:
	rm target/* -rf

