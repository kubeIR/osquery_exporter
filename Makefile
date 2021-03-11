version?=0.0.1
build?=alpha
commit=$(shell git rev-parse --short HEAD)
UNAME_S:= $(shell uname -s)
package=github.com/prateeknischal/osqueryexporter
LDFLAGS=-X $(package)/ingestd.Version=$(version)
LDFLAGS+= -X $(package)/ingestd.Commit=$(commit)
LDFLAGS+= -X $(package)/ingestd.Build=$(build)

ifeq ($(UNAME_S),Linux)
	LDFLAGS+= -extldflags '-static' -linkmode external
endif

build:
	go build -ldflags "-s -w $(LDFLAGS)" -o osquery_exporter

test:
	go test -race -v ./...
