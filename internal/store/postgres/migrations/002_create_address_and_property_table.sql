CREATE TABLE IF NOT EXISTS "address" (
  "id" UUID NOT NULL PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
  "street" VARCHAR(255) NOT NULL,
  "number" VARCHAR(20) NOT NULL,
  "complement" VARCHAR(255),
  "neighborhood" VARCHAR(100) NOT NULL,
  "city" VARCHAR(100) NOT NULL,
  "state" CHAR(2) NOT NULL,
  "zip_code" CHAR(8) NOT NULL,
  "full_address" VARCHAR(255) NOT NULL UNIQUE,
  "mini_address" VARCHAR(255) NOT NULL UNIQUE
);

CREATE TYPE "property_status" AS ENUM (
  'Disponível',
  'Ocupado',
  'Indisponível',
  'Reservado',
  'Reformando'
);

CREATE TYPE "property_kind" AS ENUM (
  'Residencial',
  'Comercial',
  'Industrial',
  'Terreno',
  'Rural'
);

CREATE TABLE IF NOT EXISTS "property" (
  "id" UUID NOT NULL PRIMARY KEY,
  "status" property_status NOT NULL,
  "kind" property_kind NOT NULL,
  "water_id" VARCHAR(20) NOT NULL,
  "energy_id" VARCHAR(20) NOT NULL,
  "address_id" UUID NOT NULL UNIQUE,
  "manager_id" UUID NOT NULL,
  FOREIGN KEY ("address_id")
  REFERENCES "address" ("id")
  ON DELETE CASCADE,
  FOREIGN KEY ("manager_id")
  REFERENCES "user" ("id")
  ON DELETE CASCADE
);

---- create above / drop below ----

DROP TABLE IF EXISTS "property";
DROP TYPE IF EXISTS "property_kind";
DROP TYPE IF EXISTS "property_status";
DROP TABLE IF EXISTS "address";