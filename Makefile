BIN_NAME=gol

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
	$(eval BUILD_VERSION_WITHOUT_V := $(shell echo $(BUILD_VERSION) | tr -d 'v'))
	sed -i '' -E 's/Version string =.+/Version string = "$(BUILD_VERSION_WITHOUT_V)"/' version.go
	git add version.go
	git commit -m 'release $(BUILD_VERSION)'
	git tag $(BUILD_VERSION)
	git push origin $(BUILD_VERSION)
	@echo "===> please publish new release"
	open https://github.com/matsu-chara/gol/releases/new
