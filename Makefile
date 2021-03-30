BINARY := dbmig
VERSION := 1.0.0
os = $(word 1, $@)
PLATFORMS := windows linux darwin

# Release

.PHONY: $(PLATFORMS)
$(PLATFORMS):
	mkdir -p release
	GOOS=$(os) GOARCH=amd64 go build -o release/$(BINARY)-$(os)-$(VERSION)-amd64 main.go

.PHONY: install
install: darwin
	cp release/$(BINARY)-darwin-amd64 ${GOPATH}/bin/$(BINARY)

.PHONY: release
release: windows linux darwin
	mv release/$(BINARY)-windows-$(VERSION)-amd64 release/$(BINARY)-windows-$(VERSION)-amd64.exe

.PHONY: clean
clean:
	rm -rf release
