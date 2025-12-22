create table if not exists "address" (
  "id" uuid not null primary key default gen_random_uuid(),
  "full_address" varchar(255) not null unique,
  "street" varchar(255) not null,
  "number" varchar(20) not null,
  "complement" varchar(255),
  "neighborhood" varchar(100) not null,
  "city" varchar(100) not null,
  "state" char(2) not null,
  "zip_code" char(8) not null,
  "active" boolean default true
);

---- create above / drop below ----

drop table if exists "address";