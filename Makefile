.PHONY: ci build test lint fmt gofmt install_gofumpt nilaway install_nilaway

ci: build test lint

build:
	go build

test:
	go clean -testcache && go test ./... -v

lint:
	golangci-lint run ./...

fmt: gofmt

gofmt:
	gofumpt -l -w .

install_gofumpt:
	which gofumpt || go install mvdan.cc/gofumpt@latest

nilaway: install_nilaway
	nilaway ./...

install_nilaway:
	which nilaway || go install go.uber.org/nilaway/cmd/nilaway@latest
