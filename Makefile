VERSION?=$(shell git describe --tags --dirty)
UNAME_S:= $(shell uname -s)
LDFLAGS= -X main.versionString=$(VERSION)

build:
	go build -ldflags "-s -w $(LDFLAGS)" -o osquery_exporter

test:
	go test -race -v ./...

docker:
	docker build -t osquery_exporter:latest \
		--build-arg VERSION=${version} \
		--build-arg OSQUERY_VERSION=4.7.0-1 \
		--build-arg OSQUERY_CONFIG='{}' \
		--build-arg OSQUERY_FLAGS='' .

