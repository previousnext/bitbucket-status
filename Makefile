#!/usr/bin/make -f

export CGO_ENABLED=0

PROJECT=github.com/previousnext/bitbucket-status

# Builds the project
build:
	gox -os='linux darwin' -arch='amd64' -output='bin/bitbucket-status_{{.OS}}_{{.Arch}}' -ldflags='-extldflags "-static"' $(PROJECT)

# Run all lint checking with exit codes for CI
lint:
	golint -set_exit_status `go list ./... | grep -v /vendor/`

# Run tests with coverage reporting
test:
	go test -cover ./...

IMAGE=previousnext/bitbucket-status
VERSION=$(shell git describe --tags --always)

# Releases the project Docker Hub
release:
	docker build -t ${IMAGE}:${VERSION} .
	docker push ${IMAGE}:${VERSION}

.PHONY: build lint test release
