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
	Salt           string
}

type LogIn struct {
	Login string
	Pass  string
}

type UserInfo struct {
	Id    int64
	LogIn string
}

type ClientInfo struct {
	Id       int64
	Name     string
	RedirUrl string
}

var (
	db       *sql.DB
	salt     string
	NotFound = fmt.Errorf("Not Found")
)

func ApplyConfig(config DatabaseConfig) (err error) {
	dbinfo := fmt.Sprintf("postgres://%s:%s@%s/%s",
		config.Username, config.Pass, config.DatabaseAddres, config.DatabaseName)
	db, err = sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal("ERROR: %s", err.Error())
	}
	db.Ping()
	salt = config.Salt
	return nil
}

func CheckPass(authdata LogIn) (result UserInfo, err error) {
	rows, err := db.Query("SELECT ID, LOGIN FROM AUTH_INFO WHERE LOGIN = $1 AND PASSHASH = crypt($2, $3)", authdata.Login, authdata.Pass, salt)
	if err != nil {
		log.Print(err.Error())
		return
	}
	defer rows.Close()
	if !rows.Next() {
		dbErr := rows.Err()
		if dbErr == nil {
			err = NotFound
			return
		}
		err = fmt.Errorf("ERROR: %s", dbErr.Error())
		return
	}
	rows.Scan(&result.Id, &result.LogIn)
	return
}

func GetClientInfo(clientId int64) (result ClientInfo, err error) {
	rows, err := db.Query("SELECT ID, NAME, REDIR_URL FROM CLIENTS_INFO WHERE ID = $1", clientId)
	if err != nil {
		log.Print(err.Error())
		return
	}
	defer rows.Close()
	if !rows.Next() {
		dbErr := rows.Err()
		if dbErr == nil {
			err = NotFound
			return
		}
		err = fmt.Errorf("ERROR: %s", dbErr.Error())
		return
	}
	rows.Scan(&result.Id, &result.Name, &result.RedirUrl)
	return
}
