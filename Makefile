EXECUTABLE=./Bin/captura_artefatos
WINDOWS=$(EXECUTABLE)_windows_amd64.exe
LINUX=$(EXECUTABLE)_linux_amd64
DARWIN=$(EXECUTABLE)_darwin_amd64.dmg
VERSION=$(shell git describe --tags --always --long --dirty)

.PHONY: all test clean

all: clean test build

test:
	./scripts/test_unit.sh

build: windows linux darwin
	@echo version: $(VERSION)

windows: $(WINDOWS)

linux: $(LINUX)

darwin: $(DARWIN)

$(WINDOWS):
	env GOOS=windows GOARCH=amd64 go build -v -o $(WINDOWS) -ldflags="-s -w -X main.version=$(VERSION)" ./Main/Main.go

$(LINUX):
	env GOOS=linux GOARCH=amd64 go build -v -o $(LINUX) -ldflags="-s -w -X main.version=$(VERSION)" ./Main/Main.go

$(DARWIN):
	env GOOS=darwin GOARCH=amd64 go build -v -o $(DARWIN) -ldflags="-s -w -X main.version=$(VERSION)" ./Main/Main.go

clean:
	rm -rf $(WINDOWS) $(LINUX) $(DARWIN)
