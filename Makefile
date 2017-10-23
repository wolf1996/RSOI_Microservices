userclient:
	mkdir server
	cp ../user/server.protoc server.protoc
	protoc --plugin=/home/ksg/disk_d/GoLang/bin/protoc-gen-go --go_out=plugins=grpc:server ./server.protoc