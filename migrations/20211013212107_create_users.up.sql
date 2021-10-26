CREATE TABLE users (
  id bigserial not null primary key,
  email varchar not null unique,
  name varchar not null,
  password varchar not null
);