create type property_status as enum (
  'Disponível',
  'Ocupado',
  'Indisponível',
  'Reservado',
  'Reformando'
);

create type property_kind as enum (
  'Residencial',
  'Comercial',
  'Industrial',
  'Terreno',
  'Rural'
);

create table if not exists "property" (
  "id" uuid not null primary key,
  "address_id" uuid not null unique references "address" ("id"),
  "manager_id" uuid not null references "manager" ("id"),
  "owner_id" uuid not null references "owner" ("id"),
  "status" property_status not null,
  "kind" property_kind not null,
  "water_id" varchar(20) not null unique,
  "energy_id" varchar(20) not null unique,
  "active" boolean default true
);

---- create above / drop below ----

drop table if exists "property";
drop type if exists "property_kind";
drop type if exists "property_status";