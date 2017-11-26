GOPATH = $(realpath ../../../../../)
PIDFILE = $(GOPATH)/pid/quemanager.pid
LOGFILE = $(GOPATH)/logs/quemanager.log

build_gateway_queuemanager:
	GOPATH=$(GOPATH) go build -o $(GOPATH)/bin/gateway_queuemanager queuemanager.go

start_queuemanager_server:
	nohup $(GOPATH)/bin/gateway_queuemanager > $(LOGFILE) 2>&1 & echo $$!> $(PIDFILE)

stop_queuemanager_server:
	cat $(PIDFILE) | xargs kill