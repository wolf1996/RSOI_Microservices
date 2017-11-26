package application

import (
	"github.com/wolf1996/events/application/models"
	"testing"
	"time"
	"os"
	"github.com/wolf1996/events/application/models/eventsclient"
)

const databasename = "events_test"
const username  = "eventtest"
const testpass  = "123456"
const databaseaddres = "127.0.0.1"
const testPort = ":8080"

func TestMain(m *testing.M)  {
	conf := Config{testPort, models.DatabaseConfig{username,testpass,databasename,databaseaddres}}
	go StartApplication(conf)
	eventsclient.SetConfigs(eventsclient.Config{databaseaddres+testPort})
	eventsclient.SetConfigs(eventsclient.Config{databaseaddres+testPort})
	time.Sleep(time.Second*30)
	os.Exit(m.Run())
}

func TestInfo (t *testing.T){
	evinf, err := eventsclient.GetEventInfo(1)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	if (evinf.Description != "KISH concert") || (evinf.Owner != "eventOwner"){
		t.Log("wrong info")
		t.Fail()
	}
}

func TestIncDecr(t *testing.T)  {
	evinf, err := eventsclient.GetEventInfo(1)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	evinfI, err := eventsclient.IncrementEventUsers(1)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	evinfD, err := eventsclient.DecrementEventUsers(1)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	if (evinf.PartCount + 1) != evinfI.PartCount {
		t.Log("event increment error")
		t.Fail()
	}
	if evinf.PartCount != evinfD.PartCount {
		t.Log("event decrement error")
		t.Fail()
	}
}