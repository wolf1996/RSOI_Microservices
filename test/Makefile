GOPATHMAKE := $(realpath ../../../../../)

prepare_regs_test:
	PGPASSWORD=123456 psql -U postgres -f test_schema.sql

regs_test:
	echo $(GOPATHMAKE)
	GOPATH=$(GOPATHMAKE) go test -cover ../application
