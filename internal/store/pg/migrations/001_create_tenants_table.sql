CREATE TYPE MARITAL_STATUS AS ENUM (
  'solteiro(a)',
  'casado(a)',
  'amasiado(a)',
  'divorciado(a)',
  'viúvo(a)'
);

CREATE TABLE IF NOT EXISTS tenant (
  "id" UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
  "fullname" TEXT NOT NULL,
  "rg" VARCHAR(22) UNIQUE NOT NULL,
  "cpf" CHAR(14) UNIQUE NOT NULL,
  "occupation" TEXT NOT NULL,
  "marital_status" MARITAL_STATUS NOT NULL
);

---- create above / drop below ----

DROP TABLE IF EXISTS tenant;
DROP TYPE IF EXISTS MARITAL_STATUS;