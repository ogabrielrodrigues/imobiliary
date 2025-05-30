CREATE TABLE IF NOT EXISTS "manager" (
  "id" UUID NOT NULL PRIMARY KEY,
  "fullname" VARCHAR(50) NOT NULL,
  "phone" VARCHAR(15) NOT NULL UNIQUE,
  "email" VARCHAR(100) NOT NULL UNIQUE,
  "password" VARCHAR(100) NOT NULL
);

---- create above / drop below ----

DROP TABLE IF EXISTS "manager";