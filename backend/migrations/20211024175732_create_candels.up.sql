CREATE TABLE candels (
  id bigserial not null primary key,
  open_price float not null,
  close_price float not null,
  lowest_price float not null,
  highest_price float not null,
  trading_volume float not null,
  stock_id bigserial not null,
  time timestamptz not null
);