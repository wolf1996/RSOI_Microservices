package storage

import(
	"github.com/go-redis/redis"
	"github.com/wolf1996/stats/shared"
	"encoding/json"
	"log"
)

type Config struct {
	Addres   string
	Db int
	HashName string
}

var (
	red *redis.Client
	hashName string
)

func ApplyConfig(config Config) (err error) {
	log.Printf("connect to redis with addreds: %s, db %d, HashName %s", config.Addres, config.Db, config.HashName)
	red = redis.NewClient(&redis.Options{
		Addr:     config.Addres,
		Password: "", // no password set
		DB:       config.Db,  // use default DB
	})
	hashName = config.HashName
	err = red.Ping().Err()
	return
}

func GetVal(msg, tp  string) string {
	return msg
}

func ParseVal(val string)(msg string, tp string) {
	return val, ""
}

func ParseKey(msgid string) (msg shared.MessageId , err error){
	err = json.Unmarshal([]byte(msgid), &msg)
	if err != nil{
		return
	}
	return
}


func GetKey(msgid shared.MessageId) (key string , err error){
	bkey, err := json.Marshal(&msgid)
	if err != nil{
		return
	}
	key = string(bkey[:len(bkey)])
	return
}


func AddMessageStorage(msgid shared.MessageId,msg, tp string) (err error){
	key, err := GetKey(msgid)
	if err != nil{
		return
	}
	vl := GetVal(msg,tp)

	log.Printf("Message key =  %s \n val = %s", key[:], vl)

	res := red.HSet(hashName,string(key[:]) ,vl)
	if err = res.Err(); err != nil {
		return
	}
	return
}


func GetAllMessages()(msglist map[string]string,err error){
	result := red.HGetAll(hashName)
	if err = result.Err(); err != nil{
		return
	}
	msglist,err = result.Result()
	return
}

func RemoveMessage(msgid shared.MessageId)(err error)  {
	key, err := GetKey(msgid)
	if err != nil {
		return
	}
	log.Printf("Remove Message key =  %s \n", key[:])
	res := red.HDel(hashName,key)
	err = res.Err()
	return
}