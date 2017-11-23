GOPATHMAKE := $(realpath ../../../../../)

prepare_user_test:
	PGPASSWORD=123456 psql -U postgres -f test_schema.sql

user_test:
	echo $(GOPATHMAKE)
	GOPATH=$(GOPATHMAKE) go test -cover ../application
