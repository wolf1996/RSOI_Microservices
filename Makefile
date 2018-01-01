GOPATH= $(realpath ../../../../)
PIDFILE= $(GOPATH)/pid/events.pid
LOGFILE= $(GOPATH)/logs/events.log
PROTOCPLUG=/home/ksg/go/bin/protoc-gen-go

authtoken:
	mkdir -p authtoken
	cp ../auth/token.proto token.proto
	protoc --plugin=$(PROTOCPLUG) --go_out=plugins=grpc:authtoken ./token.proto

server:
	mkdir -p server
	protoc --plugin=$(PROTOCPLUG) --go_out=plugins=grpc:server ./server.proto


clients:\
	server \
	authtoken

clean_protoc:
	rm -rf server & rm -rf  token

create_database:
	PGPASSWORD=123456 psql -U postgres -f schema.sql

build_source:
	GOPATH=$(GOPATH) go build -o $(GOPATH)/bin/events main.go

build:\
	clean_protoc \
	clients \
	build_source

start:
	nohup $(GOPATH)/bin/events > $(LOGFILE) 2>&1 & echo $$!> $(PIDFILE)

stop:
	cat $(PIDFILE) | xargs kill