GOPATH= $(realpath ../../../../)
PIDFILE= $(GOPATH)/pid/frontend.pid
LOGFILE= $(GOPATH)/logs/frontend.log


build:
	GOPATH=$(GOPATH) go build -o $(GOPATH)/bin/frontend main.go

start:
	nohup $(GOPATH)/bin/frontend > $(LOGFILE) 2>&1 & echo $$!> $(PIDFILE)

stop:
	cat $(PIDFILE) | xargs kill