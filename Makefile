projectname?=netscan

default: fmt

.PHONY: build
build:
	@go build -o bin/$(projectname)

.PHONY: run
run:
	@go run main.go

.PHONY: test
test: clean
	go test -json -v ./... | gotestfmt

.PHONY: clean
clean:
	@rm -rf coverage.out dist/ $(projectname)

.PHONY: vet
vet:
	go vet -x ./

.PHONY: fmt
fmt:
	go fmt -x ./...