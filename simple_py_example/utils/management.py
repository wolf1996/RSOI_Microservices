import subprocess
from time import sleep

def stop_service(servicename):
    print("stopping {}".format(servicename))
    cmd = "make {}_stop".format(servicename)
    make = subprocess.Popen(cmd.split(),cwd = "../" )
    make.wait()

def start_service(servicename):
    print("starting {}".format(servicename))
    cmd = "make {}_start".format(servicename)
    make = subprocess.Popen(cmd.split(),cwd = "../" )
    make.wait()
    sleep(10)