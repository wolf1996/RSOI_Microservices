GOPATH= $(realpath ../../../../)
PIDFILE= $(GOPATH)/pid/gateway.pid
LOGFILE= $(GOPATH)/logs/gateway.log

cleanprotocs:
	rm -rf usserver
	rm -rf evserver
	rm -rf regserver
	rm -rf authtoken

userclient:
	mkdir -p usserver
	cp ../user/server.protoc usserver.protoc
	protoc --plugin=/home/ksg/disk_d/GoLang/bin/protoc-gen-go --go_out=plugins=grpc:usserver ./usserver.protoc

eventsclient:
	mkdir -p evserver
	cp ../events/server.protoc evserver.protoc
	protoc --plugin=/home/ksg/disk_d/GoLang/bin/protoc-gen-go --go_out=plugins=grpc:evserver ./evserver.protoc

registrationsclient:
	mkdir -p regserver
	cp ../registration/server.protoc regserver.protoc
	protoc --plugin=/home/ksg/disk_d/GoLang/bin/protoc-gen-go --go_out=plugins=grpc:regserver ./regserver.protoc

authtoken:
	mkdir -p authtoken
	cp ../auth/token.protoc authtoken.protoc
	protoc --plugin=/home/ksg/disk_d/GoLang/bin/protoc-gen-go --go_out=plugins=grpc:authtoken ./authtoken.protoc

clients: \
	cleanprotocs \
    userclient \
    eventsclient \
    registrationsclient \
    authtoken

build_gateway_server:
	GOPATH=$(GOPATH) go build -o $(GOPATH)/bin/gateway main.go

start_gateway_server:
	nohup $(GOPATH)/bin/gateway > $(LOGFILE) 2>&1 & echo $$!> $(PIDFILE)

stop_gateway_server:
	cat $(PIDFILE) | xargs kill

build_gateway_qmanager:
	 $(MAKE) -C queuemanager build_gateway_queuemanager

start_gateway_qmanager:
	$(MAKE) -C queuemanager start_queuemanager_server

stop_gateway_qmanager:
	$(MAKE) -C queuemanager stop_queuemanager_server


build: \
	clients \
	build_gateway_qmanager \
	build_gateway_server \

start: \
	start_gateway_server \
	start_gateway_qmanager

stop: \
	stop_gateway_qmanager \
	stop_gateway_server