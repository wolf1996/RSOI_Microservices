package client

import (
	"github.com/wolf1996/frontend/application/client/gatewayview"
	"net/http"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
)

var addres string

type Config struct {
	Backend string
} 

func ApplyConfig(config Config)  {
	addres = config.Backend
}

func GetEvents(pageNum int64, pageSize int64,cookies []*http.Cookie) (gw []gatewayview.EventInfo, resp *http.Response, err error){
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/events/%d",addres,pageNum), nil)
	q := req.URL.Query()
	q.Add("pagesize", fmt.Sprintf("%d", pageSize))
	req.URL.RawQuery = q.Encode()
	for _, i := range cookies{
		req.AddCookie(i)
	}
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

func RegistreMe(eventId int64,cookies []*http.Cookie)( resp *http.Response, err error){
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/events/%d/register",addres,eventId), nil)
	for _, i := range cookies{
		req.AddCookie(i)
	}
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


func UserInfo(cookies []*http.Cookie)(uinf gatewayview.UserInfo, resp *http.Response, err error){
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/user/info",addres), nil)
	for _, i := range cookies{
		req.AddCookie(i)
	}
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

func GetUserRegistrations(pageNum int64, pageSize int64, cookies []*http.Cookie)(regs []gatewayview.RegistrationInfo, resp *http.Response, err error){
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/user/registrations/%d",addres,pageNum), nil)
	q := req.URL.Query()
	q.Add("pagesize", fmt.Sprintf("%d", pageSize))
	req.URL.RawQuery = q.Encode()
	for _, i := range cookies{
		req.AddCookie(i)
	}
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

func RemoveReg(regId int64, cookies []*http.Cookie)( resp *http.Response, err error){
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/registrations/%d/remove",addres,regId), nil)
	for _, i := range cookies{
		req.AddCookie(i)
	}
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

func EventInfo(evId int64, cookies []*http.Cookie)(einf gatewayview.EventInfo, resp *http.Response, err error){
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/event/%d",addres,evId), nil)
	for _, i := range cookies{
		req.AddCookie(i)
	}
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

func LogIn(login, password string, cookies []*http.Cookie) (resp *http.Response, err error) {
	logform := string("{ \"login\": \"%s\", \"pass\": \"%s\"}")
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/login",addres), strings.NewReader(fmt.Sprintf(logform, login, password)))
	for _, i := range cookies{
		req.AddCookie(i)
	}
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	return
}

func GetAccessButton(id int64, cookies []*http.Cookie) (cinf gatewayview.ClientInfo, resp *http.Response, err error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/get_access/%d",addres,id), nil)
	for _, i := range cookies{
		req.AddCookie(i)
	}
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
	json.Unmarshal(bdy,&cinf)
	return
}

func GivAccess(id int64, urlRed string, cookies []*http.Cookie) (cinf gatewayview.RedirectInfo, resp *http.Response, err error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/get_access/%d",addres,id), nil)
	for _, i := range cookies{
		req.AddCookie(i)
	}
	q := req.URL.Query()
	q.Add("redirect_url", urlRed)
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
	err = json.Unmarshal(bdy,&cinf)
	return
}
