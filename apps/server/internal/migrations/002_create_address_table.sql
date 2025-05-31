CREATE TABLE IF NOT EXISTS "address" (
  "id" UUID NOT NULL PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
  "full_address" VARCHAR(255) NOT NULL UNIQUE,
  "street" VARCHAR(255) NOT NULL,
  "number" VARCHAR(20) NOT NULL,
  "complement" VARCHAR(255),
  "neighborhood" VARCHAR(100) NOT NULL,
  "city" VARCHAR(100) NOT NULL,
  "state" CHAR(2) NOT NULL,
  "zip_code" CHAR(8) NOT NULL
);

---- create above / drop below ----

DROP TABLE IF EXISTS "address";