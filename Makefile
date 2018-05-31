GOCMD=go
GOIMPORT=goimports
GOFILES=$(shell find . -type f -name '*.go' -not -path "./vendor/*")

GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GORUN=$(GOCMD) run

APP_NAME=ezphp
CMD_INSTALLER_NAME=installer
CMD_SERVER_NAME=server

CMD_INSTALLER_DIR=cmd/installer/
CMD_SERVER_DIR=cmd/server/

BUILDDIR=dist
BINARY_WIN=.exe

RELEASEDIR=$(BUILDDIR)/$(APP_NAME)
RELEASEFILE=$(APP_NAME).zip

all: release
	
setup:
	mkdir -p $(RELEASEDIR)
	
run:
	$(GORUN) $(APP_NAME).go

clean:
	$(GOCLEAN)
	rm -rf $(APP_NAME) $(CMD_INSTALLER_NAME) $(CMD_SERVER_NAME) $(BUILDDIR)

format:
	goimports -w -d $(GOFILES)
	
build-linux: clean
#	$(GOBUILD) -o $(CMD_INSTALLER_DIR)$(CMD_INSTALLER_NAME).go -o $(RELEASEDIR)/$(CMD_INSTALLER_NAME)
#	$(GOBUILD) -o $(CMD_SERVER_DIR)$(CMD_SERVER_NAME).go -o $(RELEASEDIR)/$(CMD_SERVER_NAME)
	$(GOBUILD) -o $(APP_NAME).go -o $(RELEASEDIR)/$(APP_NAME)
	
build-win: clean setup
#	GOOS=windows GOARCH=386 $(GOBUILD) -o $(CMD_INSTALLER_DIR)$(CMD_INSTALLER_NAME).go -o $(RELEASEDIR)/$(CMD_INSTALLER_NAME)$(BINARY_WIN)
#	GOOS=windows GOARCH=386 $(GOBUILD) -o $(CMD_SERVER_DIR)$(CMD_SERVER_NAME).go -o $(RELEASEDIR)/$(CMD_SERVER_NAME)$(BINARY_WIN)
	GOOS=windows GOARCH=386 $(GOBUILD) -o $(APP_NAME).go -o $(RELEASEDIR)/$(APP_NAME)$(BINARY_WIN)
	
release: clean setup build-win
	cd $(BUILDDIR); zip -r $(APP_NAME).zip $(APP_NAME)