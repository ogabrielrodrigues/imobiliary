CREATE TYPE marital_status AS ENUM (
  'Solteiro(a)',
  'Casado(a)',
  'Amasiado(a)',
  'Divorciado(a)',
  'União Estável',
  'Viúvo(a)'
);

CREATE TABLE IF NOT EXISTS "owner" (
  "id" UUID NOT NULL PRIMARY KEY,
  "fullname" VARCHAR(100) NOT NULL,
  "cpf" CHAR(14) NOT NULL UNIQUE,
  "rg" VARCHAR(15) NOT NULL UNIQUE,
  "email" VARCHAR(100) NOT NULL UNIQUE,
  "cellphone" VARCHAR(15) NOT NULL,
  "occupation" VARCHAR(100) NOT NULL,
  "marital_status" marital_status NOT NULL,
  "address_id" UUID NOT NULL,
  "manager_id" UUID NOT NULL,
  FOREIGN KEY ("address_id")
  REFERENCES "address" ("id")
  ON DELETE CASCADE,
  FOREIGN KEY ("manager_id")
  REFERENCES "manager" ("id")
  ON DELETE CASCADE
);

ALTER TABLE "property" ADD COLUMN "owner_id" UUID;
ALTER TABLE "property" ADD CONSTRAINT "property_owner_id_fkey"
FOREIGN KEY ("owner_id")
REFERENCES "owner" ("id")
ON DELETE CASCADE;

---- create above / drop below ----

ALTER TABLE "property" DROP CONSTRAINT "property_owner_id_fkey";
ALTER TABLE "property" DROP COLUMN "owner_id";
DROP TABLE IF EXISTS "owner";
DROP TYPE IF EXISTS "marital_status";