.PHONY: help fmt fmt-check vet test-core test-protocol test-jetbrains test-vscode package package-jetbrains-thin package-jetbrains-full package-vscode-local install-vscode smoke metadata-smoke aidl-smoke jni-smoke pack-smoke pack-validate-smoke rpc-smoke pack-preflight export-pack commit-message-check check

help:
	@printf '%s\n' 'make fmt                    Format Go source' 'make fmt-check              Check Go formatting' 'make vet                     Run Go static checks' 'make test-core               Run core tests' 'make test-protocol           Run protocol tests' 'make test-jetbrains          Run JetBrains tests' 'make test-vscode             Run VS Code tests' 'make package                 Build all release artifacts' 'make package-jetbrains-thin  Build development ZIP without engines' 'make package-jetbrains-full  Build release ZIP with six engines' 'make package-vscode-local    Build current-platform VSIX' 'make install-vscode          Install current-platform VSIX' 'make smoke                   Run CLI fixture smoke' 'make metadata-smoke          Read metadata fixture' 'make aidl-smoke              Read AIDL fixture' 'make jni-smoke               Read JNI fixture' 'make rpc-smoke               Run JSON-RPC daemon smoke' 'make commit-message-check    Check message width and Change-Id' 'make check                   Run all checks'

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

test-vscode:
	cd clients/vscode && npm ci && npm run check

package:
	bash packaging/build.sh

package-jetbrains-thin:
	bash packaging/package-jetbrains-thin.sh

package-jetbrains-full:
	bash packaging/build-engines.sh dist/engine >/dev/null
	bash packaging/package-jetbrains-full.sh dist/engine dist/jetbrains

package-vscode-local:
	bash packaging/package-vscode-local.sh

install-vscode: package-vscode-local
	$${VSCODE_BIN:-code} --install-extension $$(ls -t dist/vscode/aosp-ide-vscode-*.vsix | head -1) --force

verify-package:
	bash packaging/verify.sh

vet:
	go vet ./core/... ./protocol/... ./cmd/...

smoke:
	go run ./cmd/aosp-index scan --profile testdata/source-only/profile.json

metadata-smoke:
	go run ./cmd/aosp-index metadata --root testdata/metadata/out --summary

aidl-smoke:
	go run ./cmd/aosp-index aidl --root testdata/aidl/source --file android/demo/IFoo.aidl --type android.demo.IFoo --out-root testdata/metadata/out --path-map /srv/aosp/out/soong/generated=testdata/aidl/generated --path-map /srv/aosp/out/soong/.intermediates/frameworks/base/framework-minus-apex/gen=testdata/aidl/generated --cache-root .aosp-ide/cache/smoke --summary

jni-smoke:
	go run ./cmd/aosp-index jni --java-root testdata/jni/source/java --cpp-root testdata/jni/source/cpp --class android.security.Scrypt --method nativeScrypt --summary

pack-smoke:
	go test ./cmd/aosp-index -run '^TestExportPack' -count=1

pack-validate-smoke:
	go test ./cmd/aosp-index -run '^TestExportPack' -count=1

rpc-smoke:
	go test ./cmd/aosp-index -run '^TestServeJSONRPCSmoke$$' -count=1

export-pack:
	bash scripts/export-aosp-pack.sh

pack-preflight:
	go run ./cmd/aosp-index pack preflight --out-root "$${OUT_DIR}"

commit-message-check:
	bash scripts/check-commit-messages.sh

check: fmt-check vet test-core test-protocol test-vscode smoke metadata-smoke aidl-smoke jni-smoke pack-smoke pack-validate-smoke rpc-smoke commit-message-check
