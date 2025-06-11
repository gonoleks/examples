## help: Display available commands
.PHONY: help
help:
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## audit: Conduct quality checks
.PHONY: audit
audit:
	go mod verify
	go vet ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...

## format: Fix code format issues
.PHONY: format
format:
	go run mvdan.cc/gofumpt@latest -w -l .

## markdown: Find markdown format issues (Requires markdownlint-cli2)
.PHONY: markdown
markdown:
	markdownlint-cli2 "**/*.md" "#vendor"

## lint: Run lint checks
.PHONY: lint
lint:
	golangci-lint run

## tidy: Clean and tidy dependencies
.PHONY: tidy
tidy:
	go mod tidy -v

## generate: Generate code
.PHONY: generate
generate:
	go generate ./...