.PHONY: build
build:
		go build -v ./cmd/vacinex_api/

.PHONY: test
test:
		go test -v -race -timeout 30s ./ ...


.DEFAULT_GOAL := build