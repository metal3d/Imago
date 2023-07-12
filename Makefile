.PHONY: build clean dist

VERSION=$(shell \
	git describe --exact-match --tags --dirty 2>/dev/null || \
	printf "%s-%s%s" \
		$(shell git branch --show-current) \
		$(shell git log -n1 --pretty=%h) \
		$(shell git diff --quiet || echo "-dirty") \
)

TRG = linux darwin windows freebsd
ARCH = amd64 386 arm arm64 ppc64le ppc64 mips64 mips64le mips mipsle
GPG_SIGN = metal3d@gmail.com

GOFLAGS = "-ldflags=-s -w -X main.version=$(VERSION)"
BINARY = imago

build:
	@echo "Building..."
	CGO_ENABLED=0 go build $(GOFLAGS) ./cmd/$(BINARY)

dist:
ifeq ($(findstring dirty,$(VERSION)),dirty)
	@echo "Dirty version, not building"
else
	@echo "Building for all platforms..."
	rm -rf dist
	mkdir -p dist
	echo "Version: $(VERSION)"
	echo "Platforms: $(TRG)"
	@for os in $(TRG); do \
		echo "Building for $$os..."; \
		for arch in $(ARCH); do \
			echo "Building for $$os $$arch..."; \
			ext=""; \
			if [ "$$os" = "windows" ]; then \
				ext=".exe"; \
			fi; \
			CGO_ENABLED=0 GOOS=$$os GOARCH=$$arch \
				go build $(GOFLAGS) -o dist/$(BINARY)-$$os-$$arch$$ext ./cmd/$(BINARY) &>/dev/null && \
			echo "Built for $$os $$arch" || :; \
			strip dist/$(BINARY)-$$os-$$arch$$ext &>/dev/null || :; \
		done; \
	done
endif

gpg-sign:
	@echo "Signing binaries with signer $(GPG_SIGN)..."
	@for file in dist/*; do \
		if [ "$${file##*.}" = "asc" ]; then \
			continue; \
		fi; \
		gpg --armor --detach-sign --local-user $(GPG_SIGN) $$file; \
	done

clean:
	@echo "Cleaning..."
	@rm -rf dist/
