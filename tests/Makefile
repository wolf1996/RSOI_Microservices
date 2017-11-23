GOPATHMAKE := $(realpath ../../../../../)

prepare_events_test:
	PGPASSWORD=123456 psql -U postgres -f test_schema.sql

events_test:
	echo $(GOPATHMAKE)
	GOPATH=$(GOPATHMAKE) go test -cover ../application