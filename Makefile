
GOCMD=go
GOFMT=gofmt

SERVERBIN=bsph
CLIENTBIN=cli

# our directory relative to $GOPATH/src
BASEDIR=github.com/rleibl/brauspielhaus

# -------------------------------------------------------------------
all: fmt build test

# -------------------------------------------------------------------
#  For local development. Will build locally and run the server
dev: all run

# -------------------------------------------------------------------
#  Build 
# -------------------------------------------------------------------
build: build_server build_client

build_server:
	$(GOCMD) build -o $(SERVERBIN) $(BASEDIR)/cmd/bsph

build_client:
	$(GOCMD) build -o $(CLIENTBIN) $(BASEDIR)/cmd/cli

# run individual tests manually using
#     go test -v gitlab.haufedev.systems/infosec/issue-tracker/backend/disgrc -run TestSearchOptionString
test:
	# $(GOCMD) test -v $(BASEDIR)/...
	$(GOCMD) test $(BASEDIR)/config $(BASEDIR)/models

# -------------------------------------------------------------------
#  Lint 
# -------------------------------------------------------------------
#
# Cannot use -w in development. Files will be open in the editor and
# will overwrite the changes or complain about modified files.
# Also, this step does not fail when there are differences found
fmt:
	$(GOFMT) -d $(GOPATH)src/$(BASEDIR)/

# go get -u golang.org/x/lint/golint
lint:
	$(GOPATH)/bin/golint -set_exit_status $(BASEDIR)/...

run: build_server
	./$(SERVERBIN)

# -------------------------------------------------------------------
clean:
	go clean
	rm -f $(SERVERBIN) $(CLIENTBIN)
