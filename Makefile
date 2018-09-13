GOCMD=go
GOFILES=$(shell find . -type f -name '*.go' -not -path "./vendor/*")

GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GORUN=$(GOCMD) run

APP_NAME=main
BINARY_WIN=.exe

BUILDDIR=dist

RELEASE_NAME=ezphp
RELEASEDIR=$(BUILDDIR)/$(RELEASE_NAME)
RELEASEFILE=$(RELEASE_NAME).zip

all: release
	
setup:
	mkdir -p $(RELEASEDIR)
	
run:
	$(GORUN) $(APP_NAME).go

clean:
	$(GOCLEAN)
	rm -rf $(RELEASE_NAME) $(BUILDDIR)

format:
	goimports -w -d $(GOFILES)
	
build-linux: clean
	$(GOBUILD) -o $(APP_NAME).go -o $(RELEASEDIR)/$(RELEASE_NAME)
	
build-win: clean setup
	GOOS=windows GOARCH=386 $(GOBUILD) -o $(APP_NAME).go -o $(RELEASEDIR)/$(RELEASE_NAME)$(BINARY_WIN)
	
release: clean setup build-win
	cd $(BUILDDIR); zip -r $(RELEASE_NAME).zip $(RELEASE_NAME)