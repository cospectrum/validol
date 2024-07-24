.PHONY: ci build test lint fmt gofmt goimports install_goimports

ci: build test lint

build:
	go build

test:
	go clean -testcache && go test ./... -v

lint:
	golangci-lint run ./...

fmt: goimports gofmt 

gofmt:
	gofmt -w -s .

goimports: install_goimports
	goimports -w .

install_goimports:
	which goimports || go install golang.org/x/tools/cmd/goimports@latest
