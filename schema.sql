DROP DATABASE IF EXISTS usersdb;
CREATE DATABASE usersdb;

DROP USER IF EXISTS usersuser;
CREATE USER usersuser WITH  PASSWORD '123456';

GRANT ALL PRIVILEGES ON DATABASE usersdb TO usersuser;

\c usersdb;

DROP TABLE IF EXISTS USER_INFO;
CREATE TABLE USER_INFO (
  ID       SERIAL PRIMARY KEY ,
  USERNAME VARCHAR(256),
  EVENTS_NUMBER INT DEFAULT(0)
);
INSERT INTO USER_INFO VALUES  (DEFAULT, 'simpleUser', DEFAULT);
INSERT INTO USER_INFO VALUES  (DEFAULT, 'eventOwner', DEFAULT);

GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO usersuser;
