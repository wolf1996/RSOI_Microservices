"""
1
"""
import requests
from requests.auth import HTTPBasicAuth

USERLIST = {"simpleUser": "1", "eventOwner": "1",}

def getNamedSessions():
    sessions = dict()
    for username, password in USERLIST.items():
        user = requests.Session()
        user.auth = (username, password)
        sessions[username] = user
    return sessions

def main():
    sessions = getNamedSessions()
    for name, session in sessions.items():
        print(name)
        resp = session.get("http://127.0.0.1:8080/hello")
        print(resp.text)
    

if __name__ == '__main__':
    main()  