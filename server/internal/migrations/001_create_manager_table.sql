create table if not exists "manager" (
  "id" uuid not null primary key,
  "fullname" varchar(50) not null,
  "phone" varchar(15) not null unique,
  "email" varchar(100) not null unique,
  "password" varchar(100) not null,
  "active" boolean default true
);

---- create above / drop below ----

drop table if exists "manager";