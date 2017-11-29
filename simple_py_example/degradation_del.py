
import requests
import json
from requests.auth import HTTPBasicAuth
import utils.testcase as testcase
import utils.management as management
from time import sleep

def show_info_user(session):
    res = testcase.get_user_info(session)
    print(res.text)
    return res

def show_info_event(session):
    res = testcase.get_event_info(session,1)
    print(res.text)
    return res

def remove_registration(session):
    res = testcase.remove_reg(session,2)
    print(res.text)
    return res

def testscript(session, stop = True, service = "registration"):
    show_info_user(session)
    show_info_event(session)
    if stop :
        management.stop_service(service)
    remove_registration(session)
    show_info_user(session)
    show_info_event(session)
    if stop :
         management.start_service(service)
    show_info_user(session)
    show_info_event(session)

def main():
    users = {"simpleUser": "1", "eventOwner": "1",}
    sessions = testcase.getNamedSessions(users)
    simpleuser = sessions["simpleUser"]
    reg = testcase.registre_user(simpleuser, 1)
    print(reg.text)
    testscript(simpleuser,service="user")
    #testscript(simpleuser, service="event")
    #testscript(simpleuser, stop=False ,service="user")


if __name__ == '__main__':
    main()