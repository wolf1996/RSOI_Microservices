package application


import (
	"github.com/wolf1996/user/application/models"
	"testing"
	"github.com/wolf1996/user/application/userclient"
	"os"
	"time"
)

const databasename = "users_test"
const username  = "rstest"
const testpass  = "123456"
const databaseaddres = "127.0.0.1"
const testPort  = ":8080"

func TestMain(m *testing.M){
	conf := Config{testPort,models.DatabaseConfig{username, testpass, databasename, databaseaddres}}
	go StartApplication(conf)
	userclient.SetConfigs(userclient.Config{databaseaddres+testPort})
	time.Sleep(time.Second*30)
	os.Exit(m.Run())
}

func TestFst(t *testing.T)  {
	uinf, err := userclient.GetUserInfo("simpleUser")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	if uinf.Name != "simpleUser" || uinf.Count != 0{
		t.Log("Error of getting GetUserinfo")
		t.Fail()
	}
}

func TestCounters(t *testing.T)  {
	uinf, err := userclient.GetUserInfo("simpleUser")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	ui2, err := userclient.IncrementEventsCounter("simpleUser")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	ui3, err := userclient.DecrementEventsCounter("simpleUser")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	if (uinf.Count != ui3.Count) || ((uinf.Count+1) != ui2.Count){
		t.Log("failed to counter")
		t.Fail()
	}
}