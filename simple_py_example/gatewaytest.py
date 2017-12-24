"""
1
"""
import requests
import json
from requests.auth import HTTPBasicAuth
import utils.testcase as testcase

def test_all(test, sessions):
    for name, user in sessions.items():
        print("###############################")
        print(name)
        test(user)
        print("###############################")

def events_test(session):
    print("events_test")
    resp = testcase.events(session, 2)
    print(resp.text)

def hello_test(session):
    print("hello test")
    resp = testcase.get_user_hello(session)
    print(resp.text)

def user_info_test(session):
    print("user info test")
    resp = testcase.get_user_info(session)
    print(resp.text)

def registration(session):
    print("registration test")
    resp = testcase.registre_user(session, 1)
    print(resp.text)
    return resp


def removetest(session, id = 0):
    print("remove test")
    resp = testcase.remove_reg(session, id)
    print(resp.text)


def registrationslist(session, page = None):
    print('registration list test')
    resp = testcase.registrations(session, page)
    print(resp.text)

def s_user_reg(sessions):
    sess = sessions['simpleUser']
    user_info_test(sess)
    rsp = registration(sess)
    jsonResp = json.loads(rsp.text)
    print(jsonResp['id'])
    user_info_test(sess)
    registrationslist(sess)
    registrationslist(sess, 3)
    removetest(sess, int(jsonResp['id']))
    user_info_test(sess)
    

def non_reg_tes():
    res = testcase.get_event_info(requests.session, 1)  
    print(res)
    print(res.content) 
    res = testcase.get_reg_info(requests.session, 2)  
    print(res)
    print(res.content) 

def main():
    users = {"simpleUser": "123123", "eventOwner": "123456",}
    sessions = testcase.getNamedSessions(users)
    test_all(hello_test, sessions)
    test_all(user_info_test, sessions) 
    test_all(events_test, sessions)
    print('########################')
    s_user_reg(sessions)
    print('########################')



if __name__ == '__main__':
    main()