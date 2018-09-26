GOCMD=go
GOFILES=$(shell find . -type f -name '*.go' -not -path "./vendor/*")

GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean

APP_NAME=main
BINARY_WIN=.exe

BUILDDIR=dist

RELEASE_NAME=ezphp
RELEASEDIR=$(BUILDDIR)/$(RELEASE_NAME)
RELEASEFILE=$(RELEASE_NAME).zip

all: release
	
setup:
	mkdir -p $(RELEASEDIR)

clean:
	$(GOCLEAN)
	rm -rf $(RELEASE_NAME) $(BUILDDIR)

format:
	goimports -w -d $(GOFILES)
	
build-linux: clean
	$(GOBUILD) -o $(APP_NAME).go -o $(RELEASEDIR)/$(RELEASE_NAME)
	
build-win: format clean setup
	GOOS=windows GOARCH=386 $(GOBUILD) -ldflags "-s -w" -o $(APP_NAME).go -o $(RELEASEDIR)/$(RELEASE_NAME)$(BINARY_WIN)
	
release: build-win