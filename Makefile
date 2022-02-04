setup:
	go install github.com/laher/goxc@latest
	go install github.com/tcnksm/ghr@latest
	go get -d -t ./...

lint: setup
	go vet ./...

build:
	go build

.PHONY: setup lint build
