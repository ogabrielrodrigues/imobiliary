CREATE TABLE IF NOT EXISTS "user" (
  "id" UUID NOT NULL gen_random_uuid(),
  "fullname" VARCHAR(100) NOT NULL,
  "creci_id" CHAR(8) UNIQUE NOT NULL,
  "cellphone" VARCHAR(15) NOT NULL UNIQUE,
  "email" VARCHAR(100) NOT NULL UNIQUE,
  "password" VARCHAR(255) NOT NULL,
  "avatar" VARCHAR(255),
  PRIMARY KEY (id, creci_id)
);

---- create above / drop below ----

DROP TABLE IF EXISTS "user";