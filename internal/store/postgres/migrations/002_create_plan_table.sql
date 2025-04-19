CREATE TYPE "plan_kind" AS ENUM (
  'free',
  'pro'
);

CREATE TABLE IF NOT EXISTS "plan" (
  "id" UUID NOT NULL PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
  "kind" plan_kind NOT NULL UNIQUE,
  "price" NUMERIC(10, 2) NOT NULL
);

CREATE TABLE IF NOT EXISTS "user_plan" (
  "id" UUID NOT NULL PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
  "user_id" UUID NOT NULL UNIQUE,
  "plan_id" UUID NOT NULL,
  "properties_total_quota" INT NOT NULL,
  "properties_used_quota" INT NOT NULL,
  "properties_remaining_quota" INT NOT NULL,
  FOREIGN KEY ("user_id")
  REFERENCES "user" ("id")
  ON DELETE CASCADE,
  FOREIGN KEY ("plan_id")
  REFERENCES "plan" ("id")
  ON DELETE CASCADE
);

---- create above / drop below ----

DROP TABLE IF EXISTS "user_plan";
DROP TABLE IF EXISTS "plan";
DROP TYPE IF EXISTS "plan_kind";