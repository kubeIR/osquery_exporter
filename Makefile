version?=0.0.1
build?=alpha
commit=$(shell git rev-parse --short HEAD)
UNAME_S:= $(shell uname -s)
package=github.com/prateeknischal/osqueryexporter/internal
LDFLAGS= -X $(package)/constants.Version=$(version)
LDFLAGS+= -X $(package)/constants.Commit=$(commit)
LDFLAGS+= -X $(package)/constants.Build=$(build)

#ifeq ($(UNAME_S),Linux)
	#LDFLAGS+= -extldflags '-static' -linkmode external
#endif

build:
	go build -ldflags "-s -w $(LDFLAGS)" -o osquery_exporter

test:
	go test -race -v ./...
