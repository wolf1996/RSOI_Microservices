package authclient

import (
	"github.com/wolf1996/gateway/authtoken"
)

func NewTokenInt(id string) authtoken.Token {
	return authtoken.Token{id}
}