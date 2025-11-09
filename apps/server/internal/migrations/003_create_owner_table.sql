create type marital_status as enum (
  'Solteiro(a)',
  'Casado(a)',
  'Amasiado(a)',
  'Divorciado(a)',
  'União Estável',
  'Viúvo(a)'
);

create table if not exists "owner" (
  "id" uuid not null primary key,
  "manager_id" uuid not null references "manager" ("id"),
  "address_id" uuid not null references "address" ("id"),
  "fullname" varchar(100) not null,
  "cpf" char(14) not null unique,
  "rg" varchar(15) not null unique,
  "phone" varchar(15) not null,
  "email" varchar(100) not null unique,
  "occupation" varchar(100) not null,
  "marital_status" marital_status not null,
  "active" boolean default true
);

---- create above / drop below ----

drop table if exists "owner";
drop type if exists "marital_status";