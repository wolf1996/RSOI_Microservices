package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type DatabaseConfig struct {
	Username       string
	Pass           string
	DatabaseName   string
	DatabaseAddres string
}

type UserInfo struct {
	Id    int64
	Name  string
	Count int64
}

var db *sql.DB

var (
	EmptyResult = fmt.Errorf("Empty result")
	AddError    = fmt.Errorf("Addition error")
)

func ApplyConfig(config DatabaseConfig) (err error) {
	dbinfo := fmt.Sprintf("postgres://%s:%s@%s/%s",
		config.Username, config.Pass, config.DatabaseAddres, config.DatabaseName)
	db, err = sql.Open("postgres", dbinfo)
	db.Ping()
	return nil
}

func IncrementUserEventCounter(id string) (inf UserInfo, err error) {
	rows, err := db.Query("UPDATE USER_INFO SET EVENTS_NUMBER = EVENTS_NUMBER + 1 WHERE username = $1 RETURNING *;", id)
	if err != nil {
		return
	}
	defer rows.Close()
	if !rows.Next() {
		dbErr := rows.Err()
		if dbErr == nil {
			log.Printf("Failed to update %s", id)
			err = EmptyResult
			return
		}
		err = fmt.Errorf("ERROR: %s", dbErr.Error())
		return
	}
	err = rows.Scan(&inf.Id, &inf.Name, &inf.Count)
	if err != nil {
		return
	}
	return
}

func GetUserInfo(login string) (inf UserInfo, err error) {
	rows, err := db.Query("SELECT * FROM USER_INFO WHERE username = $1", login)
	if err != nil {
		return
	}
	defer rows.Close()
	if !rows.Next() {
		dbErr := rows.Err()
		if dbErr == nil {
			log.Printf("Failed to find user info %s", login)
			err = AddError
			return
		}
		err = fmt.Errorf("ERROR: %s", dbErr.Error())
		return
	}
	err = rows.Scan(&inf.Id, &inf.Name, &inf.Count)
	if err != nil {
		return
	}
	return
}

func DecrementUserEventCounter(id string) (inf UserInfo, err error) {
	rows, err := db.Query("UPDATE USER_INFO SET EVENTS_NUMBER = (CASE WHEN EVENTS_NUMBER > 0 THEN (EVENTS_NUMBER - 1) ELSE 0 END) WHERE username = $1 RETURNING *;", id)
	if err != nil {
		return
	}
	defer rows.Close()
	if !rows.Next() {
		dbErr := rows.Err()
		if dbErr == nil {
			log.Printf("Failed to update %s", id)
			err = AddError
			return
		}
		err = fmt.Errorf("ERROR: %s", dbErr.Error())
		return
	}
	err = rows.Scan(&inf.Id, &inf.Name, &inf.Count)
	if err != nil {
		return
	}
	return
}
