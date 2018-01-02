GOPATH= $(realpath ../../../../)
PIDFILE= $(GOPATH)/pid/user.pid
LOGFILE= $(GOPATH)/logs/user.log

.PHONY: authtoken server


clean_protoc:
	rm -rf server
	rm -rf  token

clients: \
	clean_protoc \
    authtoken \
    server

create_database:
	PGPASSWORD=123456 psql -U postgres -f schema.sql
	
server:
	mkdir -p server
	protoc --plugin=$(PROTOCPLUG) --go_out=plugins=grpc:server ./server.proto

authtoken:
	mkdir -p authtoken
	cp ../auth/token.proto token.proto
	protoc --plugin=$(PROTOCPLUG) --go_out=plugins=grpc:authtoken ./token.proto

build_source:
	GOPATH=$(GOPATH) go build -o $(GOPATH)/bin/user main.go

build: \
	clients \
	build_source


start:
	nohup $(GOPATH)/bin/user > $(LOGFILE) 2>&1 & echo $$!> $(PIDFILE)

stop:
	cat $(PIDFILE) | xargs kill