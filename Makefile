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

GOFLAGS = -ldflags "-X main.version=$(VERSION)"

build:
	@echo "Building..."
	go build $(GOFLAGS) ./cmd/dropshadow

dist:
# if -dirty is present, we don't build
ifeq ($(findstring dirty,$(VERSION)),dirty)
	@echo "Dirty version, not building"
else
	@echo "Building for all platforms..."
	@for os in $(TRG); do \
		for arch in $(ARCH); do \
			ext=""; \
			if [ "$$os" = "windows" ]; then \
				ext=".exe"; \
			fi; \
			GOOS=$$os GOARCH=$$arch go build $(GOFLAGS) -o dist/dropshadow-$$os-$$arch$$ext ./cmd/dropshadow &>/dev/null \
				&& echo "Built for $$os $$arch" || :; \
			strip dist/dropshadow-$$os-$$arch$$ext &>/dev/null || :; \
		done; \
	done
endif

gpg-sign:
	@echo "Signing binaries..."
	@for file in dist/*; do \
		gpg --armor --detach-sign $$file; \
	done


clean:
	@echo "Cleaning..."
	@rm -rf dist/
