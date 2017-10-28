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


type RegistrationInfo struct {
	Id      int64
	UserId  int64
	EventId int64
}

var db *sql.DB


func ApplyConfig(config DatabaseConfig) (err error) {
	dbinfo := fmt.Sprintf("postgres://%s:%s@%s/%s",
		config.Username, config.Pass, config.DatabaseAddres,config.DatabaseName)
	db, err = sql.Open("postgres", dbinfo)
	db.Ping()
	return nil
}


func GetRegistration(id int64) (inf RegistrationInfo, err error) {
	rows, err := db.Query("SELECT * FROM rsoi_regs WHERE id = $1", id)
	if err!=nil{
		log.Print(err.Error())
		return
	}
	defer rows.Close()
	if !rows.Next() {
		err= fmt.Errorf("ERROR: Нет такой регистрации %s",id)
		log.Print(err.Error())
		return
	}
	err = rows.Scan(&inf.Id, &inf.UserId, &inf.EventId)
	if err != nil {
		log.Print(err.Error())
		return
	}
	return
}
