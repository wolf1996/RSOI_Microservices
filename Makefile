GOPATH= $(realpath ../../../../)
PIDFILE= $(GOPATH)/pid/stats.pid
LOGFILE= $(GOPATH)/logs/stats.log
PROTOCPLUG=/home/ksg/go/bin/protoc-gen-go

.PHONY: authtoken server


authtoken:
	mkdir -p token
	cp ../auth/token.proto token.proto
	protoc --plugin=$(PROTOCPLUG) --go_out=plugins=grpc:token ./token.proto

server:
	mkdir -p server
	protoc --plugin=$(PROTOCPLUG) --go_out=plugins=grpc:server ./server.proto

clean_protoc:
	rm -rf server
	rm -rf  token

clients: \
	clean_protoc \
    authtoken \
    server

build_source:
	GOPATH=$(GOPATH) go build -o $(GOPATH)/bin/stats main.go

build: \
	clean_protoc \
	clients \
	build_source

start:
	nohup $(GOPATH)/bin/stats > $(LOGFILE) 2>&1 & echo $$!> $(PIDFILE)

stop:
	cat $(PIDFILE) | xargs kill
