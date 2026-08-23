.PHONY: help fmt fmt-check vet test-core check

help:
	@printf '%s\n' 'make fmt        Format Go source' 'make fmt-check  Check Go formatting' 'make vet        Run Go static checks' 'make test-core  Run core tests' 'make check      Run core checks'

fmt:
	gofmt -w core cmd

fmt-check:
	@test -z "$$(gofmt -l core cmd)" || (echo 'Go files need gofmt:'; gofmt -l core cmd; exit 1)

test-core:
	go test ./core/...

vet:
	go vet ./core/...

check: fmt-check vet test-core
