CREATE TABLE IF NOT EXISTS "user" (
  "id" UUID NOT NULL gen_random_uuid(),
  "creci_id" CHAR(8) UNIQUE NOT NULL
  "fullname" VARCHAR(100) NOT NULL,
  "email" VARCHAR(100) NOT NULL UNIQUE,
  "access_code" CHAR(8) UNIQUE,
  PRIMARY KEY (id, creci_id, access_code)
);

---- create above / drop below ----

DROP TABLE IF EXISTS "user";