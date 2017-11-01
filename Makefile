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
