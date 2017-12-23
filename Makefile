GOPATH= $(realpath ../../../../)
PIDFILE= $(GOPATH)/pid/auth.pid
LOGFILE= $(GOPATH)/logs/auth.log
PROTOCPLUG=/home/ksg/go/bin/protoc-gen-go

server:
	mkdir -p token
	protoc --plugin=$(PROTOCPLUG) --go_out=plugins=grpc:token server.proto token.proto

clients:\
	clean_protoc \
	server
clean_protoc:
	rm -rf token

create_database:
	PGPASSWORD=123456 psql -U postgres -f schema.sql

build_source:
	GOPATH=$(GOPATH) go build -o $(GOPATH)/bin/auth main.go

build:\
	clean_protoc \
	clients \
	build_source

start:
	nohup $(GOPATH)/bin/auth > $(LOGFILE) 2>&1 & echo $$!> $(PIDFILE)

stop:
	cat $(PIDFILE) | xargs kill