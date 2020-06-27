.PHONY: build
build:
		go build -v ./cmd/vaccinex_api/

test:
		go test -v -race -timeout 30s ./ ...

.DEFAULT_GOAL := build