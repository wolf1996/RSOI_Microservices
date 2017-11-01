authtoken:
	mkdir -p authtoken
	cp ../auth/token.protoc token.protoc
	protoc --plugin=/home/ksg/disk_d/GoLang/bin/protoc-gen-go --go_out=plugins=grpc:authtoken ./token.protoc

clients: \
    authtoken