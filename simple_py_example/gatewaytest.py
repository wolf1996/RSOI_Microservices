"""
1
"""
import requests
import json
from requests.auth import HTTPBasicAuth

USERLIST = {"simpleUser": "1", "eventOwner": "1",}

def getNamedSessions():
    sessions = dict()
    for username, password in USERLIST.items():
        user = requests.Session()
        user.auth = (username, password)
        sessions[username] = user
    return sessions

def test_all(test, sessions):
    for name, user in sessions.items():
        print("###############################")
        print(name)
        test(user)
        print("###############################")

def hello_test(session):
    resp = session.get("http://127.0.0.1:8080/hello")
    print(resp.text)

def user_info_test(session):
    resp = session.get("http://127.0.0.1:8080/user/info")
    print(resp.text)

def registration(session):
    resp = session.post("http://127.0.0.1:8080/events/1/register")
    print(resp.text)
    return resp


def removetest(session, id = 0):
    print('remove')
    resp = session.post("http://127.0.0.1:8080/registrations/{}/remove".format(id))
    print(resp.text)


def registrationslist(session, page = ""):
    print('list')
    resp = session.get("http://127.0.0.1:8080/user/registrations/{}".format(page))
    print(resp.text)

def s_user_reg(sessions):
    sess = sessions['simpleUser']
    user_info_test(sess)
    rsp = registration(sess)
    jsonResp = json.loads(rsp.text)
    print(jsonResp['id'])
    user_info_test(sess)
    registrationslist(sess)
    registrationslist(sess, "3")
    removetest(sess, jsonResp['id'])
    user_info_test(sess)
    

def non_reg_tes():
    res = requests.get("http://127.0.0.1:8080/events/1")  
    print(res)
    print(res.content) 
    res = requests.get("http://127.0.0.1:8080/registrations/2")  
    print(res)
    print(res.content) 

def main():
    sessions = getNamedSessions()
    test_all(hello_test, sessions)
    test_all(user_info_test, sessions) 
    print('########################')
    s_user_reg(sessions)
    print('########################')



if __name__ == '__main__':
    main()