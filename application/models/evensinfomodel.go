package models

import (
	"github.com/wolf1996/events/server"
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

var (
	EmptyResult = fmt.Errorf("Empty result")
	AddError = fmt.Errorf("Addition error")
)


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
		dbErr := rows.Err()
		if dbErr == nil {
			err = AddError
			return
		}
		err = fmt.Errorf("ERROR: %s", dbErr.Error())
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
		dbErr := rows.Err()
		if dbErr == nil {
			err = AddError
			return
		}
		err = fmt.Errorf("ERROR: %s", dbErr.Error())
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
		dbErr := rows.Err()
		if dbErr == nil {
			err = EmptyResult
			return
		}
		err = fmt.Errorf("ERROR: %s", dbErr.Error())
		return
	}
	err = rows.Scan(&info.Id, &info.Owner, &info.PartCount, &info.Description)
	if err != nil {
		return
	}
	return
}

func GetEvents(userId string, pageNumber int64, pageSize int64, stream server.EventService_GetEventsServer)(err error){
	var params []interface{}
	ind := 1
	qr := "SELECT * FROM events_info "
	if userId != ""{
		qr+= fmt.Sprintf(" WHERE userId = $%d ", ind)
		ind += 1		
		params = append(params, userId)
		log.Print("found param")
	}
	qr += fmt.Sprintf(" OFFSET $%d LIMIT $%d ; ", ind, ind+1)
	params = append(params, pageNumber*pageSize)	
	params = append(params, pageSize)
	log.Print(qr)		
	rows, err := db.Query(qr, params...)
	if err != nil {
		return
	}
	defer rows.Close()
	var info server.EventInfo
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
		rows.Scan(&info.Id, &info.Name, &info.Participants, &info.Description)
		if err != nil {
			return
		}
		stream.Send(&info)
		if err != nil {
			return
		}
		if !rows.Next() {
			break
		}
	}
	return
}