CREATE TABLE IF NOT EXISTS "user" (
  "id" UUID NOT NULL PRIMARY KEY,
  "fullname" VARCHAR(100) NOT NULL,
  "creci_id" CHAR(7) NOT NULL UNIQUE,
  "cellphone" VARCHAR(15) NOT NULL,
  "email" VARCHAR(100) NOT NULL UNIQUE,
  "password" VARCHAR(100) NOT NULL,
  "avatar" VARCHAR(255)
);

---- create above / drop below ----

DROP TABLE IF EXISTS "user";