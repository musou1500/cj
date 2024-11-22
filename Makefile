VERSION=0.0.1
BUILD_TARGETS= \
	build-darwin-amd64 \
	build-darwin-arm64 \
	build-linux-amd64 \
	build-linux-arm64 \
	build-windows-amd64 \
	build-windows-arm64

RELEASE_TARGETS= \
	release-darwin-amd64 \
	release-darwin-arm64 \
	release-linux-amd64 \
	release-linux-arm64 \
	release-windows-amd64 \
	release-windows-arm64

RELEASE_DIR=releases
ARTIFACTS_DIR=$(RELEASE_DIR)/artifacts/$(VERSION)
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)

# artifact: release/$(VERSION)/$(APPNAME)-$(GOOS)-$(GOARCH).tar.gz
# binary: release/$(VERSION)/$(APPNAME)-$(GOOS)-$(GOARCH)/csv2json

APPNAME=cj

release: $(RELEASE_TARGETS)
all: $(BUILD_TARGETS)

build-windows-amd64:
	@$(MAKE) build GOOS=windows GOARCH=amd64 SUFFIX=.exe

release-windows-amd64: build-windows-amd64
	@$(MAKE) release-zip GOOS=windows GOARCH=amd64

build-windows-arm64:
	@$(MAKE) build GOOS=windows GOARCH=arm64 SUFFIX=.exe

release-windows-arm64: build-windows-arm64
	@$(MAKE) release-zip GOOS=windows GOARCH=arm64


build-linux-amd64:
	@$(MAKE) build GOOS=linux GOARCH=amd64

release-linux-amd64: build-linux-amd64
	@$(MAKE) release-tgz GOOS=linux GOARCH=amd64

build-linux-arm64:
	@$(MAKE) build GOOS=linux GOARCH=arm64

release-linux-arm64: build-linux-arm64
	@$(MAKE) release-tgz GOOS=linux GOARCH=arm64

build-darwin-amd64:
	@$(MAKE) build GOOS=darwin GOARCH=amd64

release-darwin-amd64: build-darwin-amd64
	@$(MAKE) release-tgz GOOS=darwin GOARCH=amd64

build-darwin-arm64:
	@$(MAKE) build GOOS=darwin GOARCH=arm64

release-darwin-arm64: build-darwin-arm64
	@$(MAKE) release-tgz GOOS=darwin GOARCH=arm64

$(ARTIFACTS_DIR):
	mkdir -p $(ARTIFACTS_DIR)

release-tgz: $(ARTIFACTS_DIR)
	tar -czf $(ARTIFACTS_DIR)/$(APPNAME)-$(GOOS)-$(GOARCH).tar.gz -C $(RELEASE_DIR) $(APPNAME)-$(GOOS)-$(GOARCH)

release-zip: $(ARTIFACTS_DIR)
	cd $(RELEASE_DIR) && zip -r9 $(CURDIR)/$(ARTIFACTS_DIR)/$(APPNAME)-$(GOOS)-$(GOARCH).zip $(APPNAME)-$(GOOS)-$(GOARCH)

build:
	go build -o $(RELEASE_DIR)/$(APPNAME)-$(GOOS)-$(GOARCH)/$(APPNAME)$(SUFFIX) .