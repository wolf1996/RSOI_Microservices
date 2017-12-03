package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/wolf1996/registration/server"
)

type DatabaseConfig struct {
	Username       string
	Pass           string
	DatabaseName   string
	DatabaseAddres string
}

type RegistrationInfo struct {
	Id      int64
	UserId  string
	EventId int64
}

var db *sql.DB

var (
	EmptyResult = fmt.Errorf("Empty result")
	AddError = fmt.Errorf("Addition error")
)

func ApplyConfig(config DatabaseConfig) (err error) {
	dbinfo := fmt.Sprintf("postgres://%s:%s@%s/%s",
		config.Username, config.Pass, config.DatabaseAddres, config.DatabaseName)
	db, err = sql.Open("postgres", dbinfo)
	db.Ping()
	return nil
}

func AddRegistration(userId string, eventId int64) (inf RegistrationInfo, err error) {
	rows, err := db.Query("INSERT INTO RSOI_REGS VALUES  (DEFAULT, $1,$2) RETURNING *;", eventId, userId)
	if err != nil {
		return
	}
	defer rows.Close()
	if !rows.Next() {
		dbErr := rows.Err()
		if dbErr == nil {
			err = AddError
			return
		}
		err = fmt.Errorf("ERROR: %s", dbErr.Error())
		return
	}
	err = rows.Scan(&inf.Id, &inf.EventId, &inf.UserId)
	if err != nil {
		return
	}
	return
}

func GetRegistration(id int64) (inf RegistrationInfo, err error) {
	rows, err := db.Query("SELECT * FROM rsoi_regs WHERE id = $1", id)
	if err != nil {
		return
	}
	defer rows.Close()
	if !rows.Next() {
		dbErr := rows.Err()
		if dbErr == nil {
			err = EmptyResult
			return
		}
		err = fmt.Errorf("ERROR: %s", dbErr.Error())
		return
	}
	err = rows.Scan(&inf.Id, &inf.EventId, &inf.UserId)
	if err != nil {
		return
	}
	return
}

func RemoveRegistration(id int64) (inf RegistrationInfo, err error) {
	rows, err := db.Query("DELETE FROM rsoi_regs WHERE id = $1 RETURNING *;", id)
	if err != nil {
		return
	}
	defer rows.Close()
	if !rows.Next() {
		err = EmptyResult
		return
	}
	err = rows.Scan(&inf.Id, &inf.EventId, &inf.UserId)
	if err != nil {
		return
	}
	return
}

func GetUserRegistrations(id string, pnumber int64, psize int64, stream server.RegistrationService_GetUserRegistrationsServer) (err error) {
	rows, err := db.Query("SELECT * FROM rsoi_regs WHERE user_id = $1 OFFSET $2 LIMIT $3 ;", id, pnumber*psize, psize)
	if err != nil {
		return
	}
	defer rows.Close()
	var inf server.RegistrationInfo
	if !rows.Next() {
		dbErr := rows.Err()
		if dbErr == nil {
			err = EmptyResult
			return
		}
		err = fmt.Errorf("ERROR: %s", dbErr.Error())
		return
	}
	for {
		// проверить чтоб не отваливалось
		err = rows.Scan(&inf.Id, &inf.EventId, &inf.UserId)
		if err != nil {
			return
		}
		stream.Send(&inf)
		if err != nil {
			return
		}
		if !rows.Next() {
			break
		}
	}
	return
}
