.PHONY: build clean dist
VERSION = $(shell git describe --tags --always --dirty 2>/dev/null|| echo "unknown")
TRG = linux darwin windows freebsd
ARCH = amd64 386 arm arm64 ppc64le ppc64 mips64 mips64le mips mipsle
GPG_SIGN = metal3d@gmail.com


version:
	@echo $(VERSION)

build:
	@echo "Building..."
	@go build -o dist/ ./cmd/...

dist:
	@echo "Building for all platforms..."
	@for os in $(TRG); do \
		for arch in $(ARCH); do \
			ext=""; \
			if [ "$$os" = "windows" ]; then \
				ext=".exe"; \
			fi; \
			GOOS=$$os GOARCH=$$arch go build -o dist/dropshadow-$$os-$$arch$$ext ./cmd/dropshadow &>/dev/null \
				&& echo "Built for $$os $$arch" || :; \
			strip dist/dropshadow-$$os-$$arch$$ext &>/dev/null || :; \
		done; \
	done

gpg-sign:
	@echo "Signing binaries..."
	@for file in dist/*; do \
		gpg --armor --detach-sign $$file; \
	done


clean:
	@echo "Cleaning..."
	@rm -rf dist/
