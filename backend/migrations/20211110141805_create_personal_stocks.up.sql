CREATE TABLE personal_stocks (
  id bigserial not null primary key,
  stock_id bigserial not null,
  user_id bigserial not null,
  user_stock_value float not null
);