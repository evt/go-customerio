test:
	go test -v --race ./...

lint:
	gofumpt -w -s ./..
	golangci-lint run --fix

check: test lint
