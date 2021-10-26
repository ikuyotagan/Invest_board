CREATE TABLE candels (
  id bigserial not null primary key,
  price float not null,
  name varchar not null,
  time timestamp not null
);