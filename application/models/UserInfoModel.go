package models

import (
	_ "github.com/lib/pq"
	"database/sql"
	"fmt"
)

type UserDatabaseConfig struct {
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

func ApplyConfig(config UserDatabaseConfig) (err error) {
	dbinfo := fmt.Sprintf("postgres://%s:%s@%s/%s",
		config.Username, config.Pass, config.DatabaseAddres,config.DatabaseName)
	db, err = sql.Open("postgres", dbinfo)
	db.Ping()
	return nil
}


func GetUserInfo(login string) (inf UserInfo, err error) {
	rows, err := db.Query("SELECT username, EVENTS_NUMBER FROM USER_INFO WHERE username = $1", login)
	if err!=nil{
		return
	}
	defer rows.Close()
	if !rows.Next() {
		err= fmt.Errorf("Нет такого пользователя %s",login)
		return
	}
	err = rows.Scan(&inf.Name, &inf.Count)
	if err != nil {
		return
	}
	return
}
