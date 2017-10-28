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

clients: \
    userclient \
    eventsclient \
    registrationsclient
