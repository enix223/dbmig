BINARY := dbmig
VERSION ?= latest
os = $(word 1, $@)
PLATFORMS := windows linux darwin

# Release

.PHONY: $(PLATFORMS)
$(PLATFORMS):
	mkdir -p release
	GOOS=$(os) GOARCH=amd64 go build -o release/$(BINARY)-$(VERSION)-$(os)-amd64 main.go

.PHONY: release
release: windows linux darwin

.PHONY: clean
clean:
	rm -rf release
