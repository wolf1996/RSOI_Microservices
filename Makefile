user_prepare_database:
	$(MAKE) -C ../user create_database

event_prepare_database:
	$(MAKE) -C ../events create_database

registration_prepare_database:
	$(MAKE) -C ../registration create_database

auth_prepare_database:
	$(MAKE) -C ../auth create_database

prepare_databases: \
	user_prepare_database \
	event_prepare_database \
	registration_prepare_database \
	auth_prepare_database

user_start:
	$(MAKE) -C ../user start

event_start:
	$(MAKE) -C ../events start

registration_start:
	$(MAKE) -C ../registration start

gateway_start:
	$(MAKE) -C ../gateway start

auth_start:
	$(MAKE) -C ../auth start

frontend_start:
	$(MAKE) -C ../frontend start

start_all: \
	user_start \
	event_start \
	registration_start \
	gateway_start \
	auth_start \
	frontend_start

user_stop:
	$(MAKE) -C ../user stop

event_stop:
	$(MAKE) -C ../events stop

registration_stop:
	$(MAKE) -C ../registration stop

gateway_stop:
	$(MAKE) -C ../gateway stop

auth_stop:
	$(MAKE) -C ../auth stop

frontend_stop:
	$(MAKE) -C ../frontend stop

stop_all: \
	user_stop \
	event_stop \
	registration_stop \
	gateway_stop \
	auth_stop \
	frontend_stop

user_build:
	$(MAKE) -C ../user build

event_build:
	$(MAKE) -C ../events build

registration_build:
	$(MAKE) -C ../registration build

gateway_build:
	$(MAKE) -C ../gateway build

auth_build:
	$(MAKE) -C ../auth build

frontend_build:
	$(MAKE) -C ../frontend build

build_all: \
	user_build \
	event_build \
	registration_build \
	gateway_build  \
	auth_build \
	frontend_build