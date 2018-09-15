BINARY := ranchars
VERSION := 0.1.0
FULL_VERSION := $(VERSION).$(CIRCLE_BUILD_NUM)

.PHONY: windows
windows:
	mkdir -p release
	GOOS=windows GOARCH=amd64 go build -o release/$(BINARY)-v$(FULL_VERSION)-windows-amd64/$(BINARY).exe

.PHONY: linux
linux:
	mkdir -p release
	GOOS=linux GOARCH=amd64 go build -o release/$(BINARY)-v$(FULL_VERSION)-linux-amd64/$(BINARY)

.PHONY: darwin
darwin:
	mkdir -p release
	GOOS=darwin GOARCH=amd64 go build -o release/$(BINARY)-v$(FULL_VERSION)-darwin-amd64/$(BINARY)

.PHONY: release
release: windows linux darwin
