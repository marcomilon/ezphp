GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GORUN=$(GOCMD) run
BUILDDIR=ezphp
BINARY_NAME=ezphp
BINARY_UNIX=$(BINARY_NAME)-linux
BINARY_WIN=$(BINARY_NAME).exe
PHPDIR=php-7.0.0
RELEASEDIR=release
RELEASEFILE=ezphp.zip
PUBLICDIR=public

all: build-linux
	
run:
	$(GORUN) $(BINARY_NAME).go

clean:
	$(GOCLEAN)
	rm -f $(BINARY_UNIX)
	rm -rf $(PHPDIR)
	rm -rf $(RELEASEDIR)
	rm -rf $(BUILDDIR)
	rm -rf $(PUBLICDIR)
	
build-linux: clean
	$(GOBUILD) -o $(BINARY_UNIX)
	
build-win: clean
	GOOS=windows GOARCH=386 $(GOBUILD) -o $(BUILDDIR)/$(BINARY_WIN)
	
release: clean build-win
	mkdir $(RELEASEDIR)
	zip -r $(RELEASEDIR)/$(RELEASEFILE) $(BUILDDIR)
	
	