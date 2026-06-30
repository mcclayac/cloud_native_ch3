.DEFAULT_GOAL := build

.PHONY:fmt vet build clean
build: clean
		go clean
fmt:
		go fmt ./...

vet: fmt
		go vet ./...

test: vet
		go test -v ./...

build: test
		go build -o cloud_native_ch3_app