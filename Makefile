.PHONY: help fmt fmt-check vet test-core smoke metadata-smoke aidl-smoke check

help:
	@printf '%s\n' 'make fmt            Format Go source' 'make fmt-check      Check Go formatting' 'make vet            Run Go static checks' 'make test-core      Run core tests' 'make smoke          Run CLI fixture smoke' 'make metadata-smoke Read metadata fixture' 'make aidl-smoke     Read AIDL fixture' 'make check          Run all Phase 0 checks'

fmt:
	gofmt -w core cmd

fmt-check:
	@test -z "$$(gofmt -l core cmd)" || (echo 'Go files need gofmt:'; gofmt -l core cmd; exit 1)

test-core:
	go test ./core/...

vet:
	go vet ./core/... ./cmd/...

smoke:
	go run ./cmd/aosp-index scan --profile testdata/source-only/profile.json

metadata-smoke:
	go run ./cmd/aosp-index metadata --root testdata/metadata/out --summary

aidl-smoke:
	go run ./cmd/aosp-index aidl --root testdata/aidl/source --file android/demo/IFoo.aidl --type android.other.IBar --summary

check: fmt-check vet test-core smoke metadata-smoke aidl-smoke
