version?=0.0.1
build?=alpha
commit=$(shell git rev-parse --short HEAD)
UNAME_S:= $(shell uname -s)
package=github.com/prateeknischal/osqueryexporter
LDFLAGS=-X $(package)/constants.version=$(version)
LDFLAGS+= -X $(package)/constants.commit=$(commit)
LDFLAGS+= -X $(package)/constants.build=$(build)

ifeq ($(UNAME_S),Linux)
	LDFLAGS+= -extldflags '-static' -linkmode external
endif

build:
	go build -ldflags "-s -w $(LDFLAGS)" -o osquery_exporter

test:
	go test -race -v ./...
