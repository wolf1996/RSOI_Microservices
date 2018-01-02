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
	RetriesHashName string
}

var (
	red *redis.Client
	hashName string
	retriesHashName string
)

func ApplyConfig(config Config) (err error) {
	log.Printf("connect to redis with addreds: %s,\n  db %d,\n HashName %s," +
		" RetriesHashName %s\n", config.Addres, config.Db, config.HashName, config.RetriesHashName)
	red = redis.NewClient(&redis.Options{
		Addr:     config.Addres,
		Password: "", // no password set
		DB:       config.Db,  // use default DB
	})
	hashName = config.HashName
	retriesHashName = config.RetriesHashName
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


func AddMessageStorage(msgid shared.MessageId,msg, tp string, retries int) (err error){
	key, err := GetKey(msgid)
	if err != nil{
		return
	}
	vl := GetVal(msg,tp)

	log.Printf("Message key =  %s \n val = %s", key[:], vl)
	plpn := red.TxPipeline()
	cnt := plpn.ZAdd(retriesHashName,redis.Z{Score:float64(retries),Member:string(key[:])})
	if err = cnt.Err(); err != nil {
		return
	}

	res := plpn.HSet(hashName,string(key[:]) ,vl)
	if err = res.Err(); err != nil {
		return
	}
	_, err = plpn.Exec()
	return
}

func CleanList() (err error) {
	retriedRes := red.ZRangeByScore(retriesHashName,redis.ZRangeBy{Min:"-inf",Max:"(1"})
	if err = retriedRes.Err(); err != nil{
		return
	}
	retried,err := retriedRes.Result()
	if err != nil {
		return
	}
	for _, name := range retried {
		red.ZRem(retriesHashName, name)
		red.HDel(hashName, name)
	}
	log.Printf("Deleted %v", retried)
	retries := red.ZRange(retriesHashName, 0, -1)
	if err = retries.Err(); err != nil{
		return
	}
	retrDecr,err := retries.Result()
	if err != nil {
		return
	}
	for _, name := range retrDecr {
		red.ZIncrBy(retriesHashName,-1, name)
	}
	log.Printf("Decremented %v", retrDecr)
	return
}

func ReLoadNRefresh()(msglist map[string]string,err error){
	err = CleanList()
	if err != nil {
		return
	}
	result := red.HGetAll(hashName)
	if err = result.Err(); err != nil{
		return
	}
	msglist,err = result.Result()
	if err != nil {
		return
	}
	return
}


func RemoveMessage(msgid shared.MessageId)(err error)  {
	key, err := GetKey(msgid)
	if err != nil {
		return
	}
	log.Printf("Remove Message key =  %s \n", key[:])
	plpn := red.TxPipeline()
	cnt := plpn.ZRem(retriesHashName, key)
	if err = cnt.Err(); err != nil {
		return
	}

	res := plpn.HDel(hashName, key)
	if err = res.Err(); err != nil {
		return
	}
	_, err = plpn.Exec()
	return
}