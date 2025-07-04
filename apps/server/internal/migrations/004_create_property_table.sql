CREATE TYPE property_status AS ENUM (
  'Disponível',
  'Ocupado',
  'Indisponível',
  'Reservado',
  'Reformando'
);

CREATE TYPE property_kind AS ENUM (
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
  "water_id" VARCHAR(20) NOT NULL UNIQUE,
  "energy_id" VARCHAR(20) NOT NULL UNIQUE,
  "address_id" UUID NOT NULL UNIQUE,
  "manager_id" UUID NOT NULL,
  "owner_id" UUID NOT NULL,
  FOREIGN KEY ("address_id")
  REFERENCES "address" ("id")
  ON DELETE CASCADE,
  FOREIGN KEY ("manager_id")
  REFERENCES "manager" ("id")
  ON DELETE CASCADE,
  FOREIGN KEY ("owner_id")
  REFERENCES "owner" ("id")
  ON DELETE CASCADE
);

---- create above / drop below ----

DROP TABLE IF EXISTS "property";
DROP TYPE IF EXISTS "property_kind";
DROP TYPE IF EXISTS "property_status";