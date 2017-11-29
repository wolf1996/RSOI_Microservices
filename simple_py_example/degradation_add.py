
import requests
import json
from requests.auth import HTTPBasicAuth
import utils.testcase as testcase
import utils.management as management
import subprocess
from time import sleep

def show_info_user(session):
    res = testcase.get_user_info(session)
    print(res.text)
    return res

def show_info_event(session):
    res = testcase.get_event_info(session,1)
    print(res.text)
    return res

def registration(session):
    res = testcase.registre_user(session,1)
    print(res.text)
    return res

def testscript(session, stop = True, service = "registration"):
    show_info_user(session)
    show_info_event(session)
    if stop :
        management.stop_service(service)
    registration(session)
    show_info_user(session)
    show_info_event(session)
    if stop :
         management.start_service(service)

def main():
    users = {"simpleUser": "1", "eventOwner": "1",}
    sessions = testcase.getNamedSessions(users)
    simpleuser = sessions["simpleUser"]
    testscript(simpleuser)
    testscript(simpleuser, service="event")
    testscript(simpleuser, service="user")
    testscript(simpleuser, stop=False)


if __name__ == '__main__':
    main()