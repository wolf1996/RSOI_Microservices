import requests

def getNamedSessions(userpassdict):
    sessions = dict()
    print("Start LogIn")
    for username, password in userpassdict.items():
        user = requests.Session()
        res = sign_in(user,username,password)
        print("user {} logged in with password {}".format(username, password))
        print(res)
        print(res.text)
        sessions[username] = user
    print("LogIn finished \n\n\n")
    return sessions


def sign_in(session, login, password):
    return session.post("http://127.0.0.1:8080/login",json={"login":login, "pass":password} )

def get_user_hello(session):
    return session.get("http://127.0.0.1:8080/hello")

def get_user_info(session):
    return session.get("http://127.0.0.1:8080/user/info")

def registre_user(session, event_id):
    return session.post("http://127.0.0.1:8080/events/{}/register".format(event_id))

def remove_reg(session, reg_id):
    return session.delete("http://127.0.0.1:8080/registrations/{}/remove".format(reg_id))

def registrations(session, page_num = None):
    if page_num is None:
        page_num = ""
    else:
        page_num = int(page_num)
    return session.get("http://127.0.0.1:8080/user/registrations/{}".format(page_num))

def events(session, page_num = None):
    if page_num is None:
        page_num = ""
    else:
        page_num = int(page_num)
    return session.get("http://127.0.0.1:8080/events/{}".format(page_num))

def get_event_info(session, event_id = None):
    if event_id is None:
        event_id = ""
    else:
        event_id = int(event_id)
    return requests.get("http://127.0.0.1:8080/event/{}".format(event_id))  

def get_reg_info(session, reg_id):
    return requests.get("http://127.0.0.1:8080/registrations/{}",reg_id)  
