PROJECTNAME = netscan
BIN = netscan

# if the os is Windows then we should generate an exe
ifeq ($(OS),Windows_NT) 
    BIN = netscan.exe
endif

default: fmt

.PHONY: build
build:
	@go build -o bin/$(BIN)

.PHONY: run
run:
	@go run main.go

.PHONY: test
test: clean
	go test -json -v ./... | gotestfmt

.PHONY: clean
clean:
	@rm -rf coverage.out dist/ $(PROJECTNAME)

.PHONY: vet
vet:
	go vet -x ./

.PHONY: fmt
fmt:
	go fmt -x ./...