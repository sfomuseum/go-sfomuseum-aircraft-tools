CWD=$(shell pwd)
GOPATH := $(CWD)

prep:
	if test -d pkg; then rm -rf pkg; fi

self:   prep rmdeps
	if test -d src; then rm -rf src; fi
	mkdir -p src/github.com/sfomuseum/go-sfomuseum-aircraft-tools
	cp -r template src/github.com/sfomuseum/go-sfomuseum-aircraft-tools/
	cp -r vendor/* src/

rmdeps:
	if test -d src; then rm -rf src; fi 

build:	fmt bin

deps:
	@GOPATH=$(GOPATH) go get -u "github.com/sfomuseum/go-sfomuseum-aircraft"
	@GOPATH=$(GOPATH) go get -u "github.com/sfomuseum/go-sfomuseum-geojson"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-index"
	mv src/github.com/sfomuseum/go-sfomuseum-geojson/vendor/github.com/whosonfirst/go-whosonfirst-geojson-v2 src/github.com/whosonfirst
	mv src/github.com/sfomuseum/go-sfomuseum-geojson/vendor/github.com/whosonfirst/go-whosonfirst-flags src/github.com/whosonfirst
	mv src/github.com/sfomuseum/go-sfomuseum-geojson/vendor/github.com/whosonfirst/go-whosonfirst-placetypes src/github.com/whosonfirst
	rm -rf src/github.com/whosonfirst/go-whosonfirst-index/vendor/github.com/whosonfirst/go-whosonfirst-geojson-v2
	rm -rf src/github.com/whosonfirst/go-whosonfirst-index/vendor/github.com/whosonfirst/go-whosonfirst-flags

vendor-deps: rmdeps deps
	if test ! -d vendor; then mkdir vendor; fi
	if test -d vendor; then rm -rf vendor; fi
	cp -r src vendor
	find vendor -name '.git' -print -type d -exec rm -rf {} +
	rm -rf src

fmt:
	# go fmt *.go
	go fmt cmd/*.go
	go fmt template/*.go

bin: 	self
	rm -rf bin/*
	@GOPATH=$(GOPATH) go build -o bin/build-sfomuseum-data cmd/build-sfomuseum-data.go
	@GOPATH=$(GOPATH) go build -o bin/build-icao-data cmd/build-icao-data.go
