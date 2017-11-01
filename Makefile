authtoken:
	mkdir -p authtoken
	cp ../auth/token.protoc token.protoc
	protoc --plugin=/home/ksg/disk_d/GoLang/bin/protoc-gen-go --go_out=plugins=grpc:authtoken ./token.protoc

server:
	mkdir -p server
	protoc --plugin=/home/ksg/disk_d/GoLang/bin/protoc-gen-go --go_out=plugins=grpc:server ./server.protoc


clean_protoc:
	rm -rf server
	rm -rf  token

clients: \
	clean_protoc \
    authtoken \
    server