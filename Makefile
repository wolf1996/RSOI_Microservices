GOPATH= $(realpath ../../../../)
PIDFILE= $(GOPATH)/pid/registration.pid
LOGFILE= $(GOPATH)/logs/registration.log
PROTOCPLUG=/home/ksg/go/bin/protoc-gen-go

authtoken:
	mkdir -p authtoken
	cp ../auth/token.protoc token.protoc
	protoc --plugin=$(PROTOCPLUG) --go_out=plugins=grpc:authtoken ./token.protoc

create_database:
	PGPASSWORD=123456 psql -U postgres -f schema.sql

server:
	mkdir -p server
	protoc --plugin=$(PROTOCPLUG) --go_out=plugins=grpc:server ./server.protoc

clean_protoc:
	rm -rf server
	rm -rf  token

clients: \
	clean_protoc \
    authtoken \
    server

build_source:
	GOPATH=$(GOPATH) go build -o $(GOPATH)/bin/registration main.go

build: \
	clean_protoc \
	clients \
	build_source

start:
	nohup $(GOPATH)/bin/registration > $(LOGFILE) 2>&1 & echo $$!> $(PIDFILE)

stop:
	cat $(PIDFILE) | xargs kill
