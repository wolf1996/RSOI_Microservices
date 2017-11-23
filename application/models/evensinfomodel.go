package models

import (
	_ "github.com/lib/pq"
	"database/sql"
	"fmt"
	"log"
)

type EventInfo struct {
	Id              int64
	Owner           string
	PartCount       int64
	Description     string
}

type DatabaseConfig struct {
	Username       string
	Pass           string
	DatabaseName   string
	DatabaseAddres string
}

var db *sql.DB


func ApplyConfig(config DatabaseConfig) (err error) {
	dbinfo := fmt.Sprintf("postgres://%s:%s@%s/%s",
		config.Username, config.Pass, config.DatabaseAddres,config.DatabaseName)
	db, err = sql.Open("postgres", dbinfo)
	db.Ping()
	return nil
}

func IncrementEventUserCounter(id int64) (info EventInfo, err error) {
	rows, err := db.Query("UPDATE events_info SET PARTICIPANT_COUNT = PARTICIPANT_COUNT + 1 WHERE id = $1 RETURNING *;", id)
	if err!=nil{
		log.Print(err.Error())
		return
	}
	defer rows.Close()
	if !rows.Next() {
		err= fmt.Errorf("ERROR: Нет такого ивента %d",id)
		log.Print(err.Error())
		return
	}
	err = rows.Scan(&info.Id, &info.Owner, &info.PartCount, &info.Description)
	if err != nil {
		log.Print(err.Error())
		return
	}
	return
}


func DecrementEventUserCounter(id int64) (info EventInfo, err error) {
	rows, err := db.Query("UPDATE events_info SET PARTICIPANT_COUNT =  (CASE WHEN PARTICIPANT_COUNT > 0 THEN (PARTICIPANT_COUNT - 1)ELSE 0 END) WHERE id = $1 RETURNING *;", id)
	if err!=nil{
		log.Print(err.Error())
		return
	}
	defer rows.Close()
	if !rows.Next() {
		err= fmt.Errorf("ERROR: Нет такого ивента %d",id)
		log.Print(err.Error())
		return
	}
	err = rows.Scan(&info.Id, &info.Owner, &info.PartCount, &info.Description)
	if err != nil {
		log.Print(err.Error())
		return
	}
	return
}

func GetEventInfo(id int64)(info EventInfo, err error){
	rows, err := db.Query("SELECT * FROM events_info WHERE id = $1", id)
	if err!=nil{
		log.Print(err.Error())
		return
	}
	defer rows.Close()
	if !rows.Next() {
		err= fmt.Errorf("ERROR: Нет такого ивента %d",id)
		log.Print(err.Error())
		return
	}
	err = rows.Scan(&info.Id, &info.Owner, &info.PartCount, &info.Description)
	if err != nil {
		log.Print(err.Error())
		return
	}
	return
}