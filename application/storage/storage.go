package storage

import(
	"github.com/go-redis/redis"
	"time"
)

type Config struct {
	Addres   string
}

var (
	red *redis.Client
)

func ApplyConfig(config Config) (err error) {
	red = redis.NewClient(&redis.Options{
		Addr:     config.Addres,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	err = red.Ping().Err()
	return
}

func AddTokenStorage(token string, expTime time.Time) (err error){
	res := red.Set(token, "",0)
	if err = res.Err(); err != nil {
		return
	}
	res2 := red.ExpireAt(token,expTime)
	if err = res2.Err(); err != nil {
		return
	}
	return
}


func CheckTokenStorage(token string)(exists bool,err error){
	res := red.Exists(token)
	if err = res.Err(); err != nil {
		return
	}
	exs, err  := res.Result()
	if err != nil {
		return
	}
	exists = (exs != 0)
	return
}

func RemoveToken(token string)(err error)  {
	res := red.Del(token)
	err = res.Err()
	return
}