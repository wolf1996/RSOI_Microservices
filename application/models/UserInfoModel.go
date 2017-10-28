package models

import (
	_ "github.com/lib/pq"
	"database/sql"
	"fmt"
	"log"
)

type DatabaseConfig struct {
	Username       string
	Pass           string
	DatabaseName   string
	DatabaseAddres string
}


var db *sql.DB

type UserInfo struct {
	Name string
	Count int64
}

func ApplyConfig(config DatabaseConfig) (err error) {
	dbinfo := fmt.Sprintf("postgres://%s:%s@%s/%s",
		config.Username, config.Pass, config.DatabaseAddres,config.DatabaseName)
	db, err = sql.Open("postgres", dbinfo)
	db.Ping()
	return nil
}


func GetUserInfo(login string) (inf UserInfo, err error) {
	rows, err := db.Query("SELECT username, EVENTS_NUMBER FROM USER_INFO WHERE username = $1", login)
	if err!=nil{
		log.Print(err.Error())
		return
	}
	defer rows.Close()
	if !rows.Next() {
		err= fmt.Errorf("ERROR: Нет такого пользователя %s",login)
		log.Print(err.Error())
		return
	}
	err = rows.Scan(&inf.Name, &inf.Count)
	if err != nil {
		log.Print(err.Error())
		return
	}
	return
}
