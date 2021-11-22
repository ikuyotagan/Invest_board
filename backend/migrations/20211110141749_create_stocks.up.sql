CREATE TABLE stocks (
  id bigserial not null primary key,
  name varchar not null unique,
  figi varchar not null unique
);