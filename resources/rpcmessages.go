package resources

import (
	"github.com/wolf1996/gateway/token"
	"context"
	"github.com/golang/protobuf/proto"
	"log"
	"google.golang.org/grpc/metadata"
	"encoding/base64"
)

func TokenToString(tkn token.Token)(btTok []byte, err error ){
	btTok, err = proto.Marshal(&tkn)
	if err != nil {
		log.Print(err.Error())
		return
	}
	return
}

func TokenToContext(tkn token.Token)(ctx context.Context, err error){
	btTok, err := TokenToString(tkn)
	if err != nil {
		return
	}
	md := metadata.Pairs("token", base64.StdEncoding.EncodeToString(btTok[:]))
	ctx = metadata.NewOutgoingContext(context.Background(), md)
	return
}