.PHONY: help fmt fmt-check vet test-core test-protocol test-jetbrains smoke metadata-smoke aidl-smoke jni-smoke rpc-smoke commit-message-check check

help:
	@printf '%s\n' 'make fmt                  Format Go source' 'make fmt-check            Check Go formatting' 'make vet                  Run Go static checks' 'make test-core            Run core tests' 'make test-protocol        Run protocol tests' 'make smoke                Run CLI fixture smoke' 'make metadata-smoke       Read metadata fixture' 'make aidl-smoke             Read AIDL fixture' 'make jni-smoke              Read JNI fixture' 'make rpc-smoke             Run JSON-RPC daemon smoke' 'make commit-message-check  Check message width and Change-Id' 'make check                 Run all checks'

fmt:
	gofmt -w core protocol cmd

fmt-check:
	@test -z "$$(gofmt -l core protocol cmd)" || (echo 'Go files need gofmt:'; gofmt -l core protocol cmd; exit 1)

test-core:
	go test ./core/...

test-protocol:
	go test ./protocol/...

test-jetbrains:
	cd clients/jetbrains && env JAVA_HOME="$${JAVA_HOME_21:-$${JAVA_HOME}}" ./gradlew --no-daemon $${INTELLIJ_PLATFORM_PATH:+-PintellijPlatformPath="$${INTELLIJ_PLATFORM_PATH}"} test buildPlugin

vet:
	go vet ./core/... ./protocol/... ./cmd/...

smoke:
	go run ./cmd/aosp-index scan --profile testdata/source-only/profile.json

metadata-smoke:
	go run ./cmd/aosp-index metadata --root testdata/metadata/out --summary

aidl-smoke:
	go run ./cmd/aosp-index aidl --root testdata/aidl/source --file android/demo/IFoo.aidl --type android.demo.IFoo --metadata-root testdata/metadata/out --path-map /srv/aosp/out/soong/generated=testdata/aidl/generated --path-map /srv/aosp/out/soong/.intermediates/frameworks/base/framework-minus-apex/gen=testdata/aidl/generated --cache-root .aosp-ide/cache/smoke --summary

jni-smoke:
	go run ./cmd/aosp-index jni --java-root testdata/jni/source/java --cpp-root testdata/jni/source/cpp --class android.security.Scrypt --method nativeScrypt --summary

rpc-smoke:
	go test ./cmd/aosp-indexd -run '^TestDaemonSmoke$$' -count=1

commit-message-check:
	bash scripts/check-commit-messages.sh

check: fmt-check vet test-core test-protocol smoke metadata-smoke aidl-smoke jni-smoke rpc-smoke commit-message-check
