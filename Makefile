
GOCMD=go
GOBUILD=$(GOCMD) build

build:
	$(GOBUILD) -o bin/cli -v cmd/cli/main.go
	$(GOBUILD) -o bin/srv -v cmd/srv/main.go

clean:
	rm -f bin/cli
	rm -f bin/srv

