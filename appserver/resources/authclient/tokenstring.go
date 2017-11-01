package authclient

import (
	"github.com/wolf1996/gateway/authtoken"
)

func NewTokenInt(id int64) authtoken.Token {
	return authtoken.Token{id}
}