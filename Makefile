CWD=$(shell pwd)
GOPATH := $(CWD)

prep:
	if test -d pkg; then rm -rf pkg; fi

self:   prep
	if test -d src/github.com/whosonfirst/go-whosonfirst-updated-v2; then rm -rf src/github.com/whosonfirst/go-whosonfirst-updated-v2; fi
	mkdir -p src/github.com/whosonfirst/go-whosonfirst-updated-v2
	cp -r copy src/github.com/whosonfirst/go-whosonfirst-updated-v2/
	cp -r flags src/github.com/whosonfirst/go-whosonfirst-updated-v2/
	cp -r process src/github.com/whosonfirst/go-whosonfirst-updated-v2/
	cp *.go src/github.com/whosonfirst/go-whosonfirst-updated-v2/
	cp -r vendor/* src/

rmdeps:
	if test -d src; then rm -rf src; fi 

build:	fmt bin

deps:
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-webhookd"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-redis"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-readwrite/..."
	@GOPATH=$(GOPATH) go get -u "gopkg.in/redis.v1"
	rm -rf src/github.com/whosonfirst/go-webhookd/vendor/github.com/golang

vendor-deps: rmdeps deps
	if test ! -d vendor; then mkdir vendor; fi
	if test -d vendor; then rm -rf vendor; fi
	cp -r src vendor
	find vendor -name '.git' -print -type d -exec rm -rf {} +
	rm -rf src

fmt:
	go fmt cmd/*.go
	go fmt copy/*.go
	go fmt flags/*.go
	go fmt process/*.go
	go fmt *.go

bin: 	rmdeps self
	@GOPATH=$(shell pwd) go build -o bin/wof-updated cmd/wof-updated.go
	@GOPATH=$(shell pwd) go build -o bin/wof-updated-replay cmd/wof-updated-replay.go

debug:
	if test ! -f docker/webhookd-config.json; then echo "missing docker/webhookd-config.json file, maybe make one from the example config?"; exit 1; fi
	bin/wof-updated -pubsubd -pubsub-host 0.0.0.0 -webhookd -webhookd-config docker/webhookd-config.json

docker-build:
	if test ! -f docker/webhookd-config.json; then echo "missing docker/webhookd-config.json file, maybe make one from the example config?"; exit 1; fi
	make bin
	docker build -t wof-updated .

docker-run:
	docker run -it -p 6379:6379 -p 8080:8080 wof-updated

docker-debug:
	@make docker-build
	@make docker-run
