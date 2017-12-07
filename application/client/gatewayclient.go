package client

import (
	"github.com/wolf1996/frontend/application/client/gatewayview"
	"net/http"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
)

var addres string

type Config struct {
	Backend string
} 

func ApplyConfig(config Config)  {
	addres = config.Backend
}

func GetEvents(pageNum int64, pageSize int64) (gw []gatewayview.EventInfo, resp *http.Response, err error){
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/events/%d",addres,pageNum), nil)
	q := req.URL.Query()
	q.Add("pagesize", fmt.Sprintf("%d", pageSize))
	req.URL.RawQuery = q.Encode()
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	bdy, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	log.Print(string(bdy))
	json.Unmarshal(bdy,&gw)
	return
}

func RegistreMe(eventId int64)( resp *http.Response, err error){
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/events/%d/register",addres,eventId), nil)
	req.SetBasicAuth("simpleUser","1")
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	bdy, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	log.Print(string(bdy))
	return
}


func UserInfo()(uinf gatewayview.UserInfo, resp *http.Response, err error){
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/user/info",addres), nil)
	req.SetBasicAuth("simpleUser","1")
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	bdy, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	log.Print(string(bdy))
	json.Unmarshal(bdy,&uinf)
	return
}

func GetUserRegistrations(pageNum int64, pageSize int64)(regs []gatewayview.RegistrationInfo, resp *http.Response, err error){
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/user/registrations/%d",addres,pageNum), nil)
	q := req.URL.Query()
	q.Add("pagesize", fmt.Sprintf("%d", pageSize))
	req.URL.RawQuery = q.Encode()
	req.SetBasicAuth("simpleUser","1")
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	bdy, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	log.Print(string(bdy))
	err = json.Unmarshal(bdy,&regs)
	return
}

func RemoveReg(regId int64)( resp *http.Response, err error){
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/registrations/%d/remove",addres,regId), nil)
	req.SetBasicAuth("simpleUser","1")
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	bdy, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	log.Print(string(bdy))
	return
}

func EventInfo(evId int64)(einf gatewayview.EventInfo, resp *http.Response, err error){
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/event/%d",addres,evId), nil)
	req.SetBasicAuth("simpleUser","1")
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	bdy, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	log.Print(string(bdy))
	json.Unmarshal(bdy,&einf)
	return
}
