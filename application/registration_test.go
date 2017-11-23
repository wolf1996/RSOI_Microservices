package application

import (
	"testing"
	"time"
	"os"
	"github.com/wolf1996/registration/application/registrationclient"
	"github.com/wolf1996/registration/application/models"
	"github.com/wolf1996/gateway/authtoken"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/metadata"
	"encoding/base64"
)

const databasename = "reg_test"
const username  = "regrstest"
const testpass  = "123456"
const databaseaddres = "127.0.0.1"
const testPort = ":8080"

func TestMain(m *testing.M)  {
	conf := Config{testPort, models.DatabaseConfig{username,testpass,databasename,databaseaddres}}
	go StartApplication(conf)
	registrationclient.SetConfigs(registrationclient.Config{databaseaddres+testPort})
	registrationclient.SetConfigs(registrationclient.Config{databaseaddres+testPort})
	time.Sleep(time.Second*30)
	os.Exit(m.Run())
}

func TestGetRegInfo(t *testing.T) {
	reginfo, err := registrationclient.GetRegistrationInfo(1)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
		return
	}
	if (reginfo.EventId != 1) || (reginfo.UserId != "simpleUser"){
		t.Log("invalid event")
		t.Fail()
	}
}

func TestDelRegs(t *testing.T){
	token := authtoken.Token{"simpleUser"}
	btTok,err := proto.Marshal(&token)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	strTok := base64.StdEncoding.EncodeToString(btTok)
	md := metadata.Pairs("token", strTok)
	inf, err := registrationclient.RemoveRegistration(1,md)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	if (inf.UserId != "simpleUser") || (inf.EventId != 1){
		t.Log("Error of getting info")
		t.Fail()
	}
	_, err = registrationclient.GetRegistrationInfo(1)
	if err == nil {
		t.Log(err.Error())
		t.Fail()
		return
	}
}

func TestAddStreaming(t *testing.T)  {
	for i := 1; i <= 70; i++{
		regdata, err := registrationclient.AddRegistration("simpleUser", int64(i))
		if err != nil {
			t.Log(err.Error())
			t.Fail()
		}
		if (regdata.UserId != "simpleUser") || (regdata.EventId != int64(i)){
			t.Log("invalid addition")
			t.Fail()
		}
	}
	reg, err := registrationclient.GetRegistrations("simpleUser",1,50)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	if len(reg) != 50 {
		t.Log("invalid full page length:",len(reg))
		t.Fail()
	}
	for ind,rec := range(reg){
		if (rec.UserId != "simpleUser") || (rec.EventId != int64(ind+1)){
			t.Log("invalid addition", rec.UserId, rec.EventId)
			t.Fail()
		}
	}
	reg, err = registrationclient.GetRegistrations("simpleUser",2,50)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	if len(reg) != 20 {
		t.Log("invalid part page length")
		t.Fail()
	}
	for ind,rec := range(reg){
		if (rec.UserId != "simpleUser") || (rec.EventId != int64(ind+51)){
			t.Log("invalid addition", rec.UserId, rec.EventId)
			t.Fail()
		}
	}
}
