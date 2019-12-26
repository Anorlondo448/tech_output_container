DROP SCHEMA IF EXISTS app_database;
CREATE SCHEMA app_database;
USE app_database;

DROP TABLE IF EXISTS member;

CREATE TABLE member
(
  id           INT(10),
  name     VARCHAR(40),
  point        INT(10),
  company  VARCHAR(255)
);

INSERT INTO member (id, name, point, company) VALUES (1, "YEBISU", 95, "KIRIN");
INSERT INTO member (id, name, point, company) VALUES (2, "MALTS",  90, "SUNTORY");
INSERT INTO member (id, name, point, company) VALUES (3, "一番搾り",  100, "KIRIN");