GO = go
GOFLAGS =
INSTALLDIR = "$(APPDATA)\Elgato\StreamDeck\Plugins\dev.samwho.streamdeck.cpu.sdPlugin"

.PHONY: test install build logs

build:
	$(GO) build $(GOFLAGS)

test:
	$(GO) run $(GOFLAGS) main.go -port 12345 -pluginUUID 213 -registerEvent test -info "{\"application\":{\"language\":\"en\",\"platform\":\"mac\",\"version\":\"4.1.0\"},\"plugin\":{\"version\":\"1.1\"},\"devicePixelRatio\":2,\"devices\":[{\"id\":\"55F16B35884A859CCE4FFA1FC8D3DE5B\",\"name\":\"Device Name\",\"size\":{\"columns\":5,\"rows\":3},\"type\":0},{\"id\":\"B8F04425B95855CF417199BCB97CD2BB\",\"name\":\"Another Device\",\"size\":{\"columns\":3,\"rows\":2},\"type\":1}]}"

install: build
	rm -rf $(INSTALLDIR)
	mkdir $(INSTALLDIR)
	cp *.json $(INSTALLDIR)
	cp *.html $(INSTALLDIR)
	cp *.css $(INSTALLDIR)
	cp *.exe $(INSTALLDIR)

logs:
	tail -f "$(TMP)"/streamdeck-cpu.log*