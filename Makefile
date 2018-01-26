BIN_NAME=gol
BUILD_VERSION=latest

ci: clean devdeps dep lint build test

clean:
	go clean
	rm -f $(BIN_NAME)

devdeps:
	go get github.com/golang/dep/cmd/dep
	go get github.com/golang/lint/golint

dep:
	dep ensure

lint:
	test -z "$$(gofmt -s -d .)" || (gofmt -s -d .; exit 1)
	go vet ./...
	golint -set_exit_status $$(go list ./...)

build:
	go build

test:
	go test ./...

fmt:
	gofmt -s -w .

release:
	git tag $(BUILD_VERSION)
	git push origin $(BUILD_VERSION)

