CREATE TABLE IF NOT EXISTS "tenant" (
  "id" UUID NOT NULL PRIMARY KEY,
  "manager_id" UUID NOT NULL,
  "address_id" UUID NOT NULL,
  "fullname" VARCHAR(100) NOT NULL,
  "cpf" CHAR(14) NOT NULL UNIQUE,
  "rg" VARCHAR(15) NOT NULL UNIQUE,
  "phone" VARCHAR(15) NOT NULL,
  "occupation" VARCHAR(100) NOT NULL,
  "marital_status" marital_status NOT NULL,
  FOREIGN KEY ("manager_id")
  REFERENCES "manager" ("id")
  ON DELETE CASCADE,
  FOREIGN KEY ("address_id")
  REFERENCES "address" ("id")
  ON DELETE CASCADE
);

---- create above / drop below ----

DROP TABLE IF EXISTS "tenant";