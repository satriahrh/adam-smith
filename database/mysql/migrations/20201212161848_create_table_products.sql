-- +goose Up
create table products
(
  id           int unsigned primary key auto_increment,
  sku          varchar(255),
  name         varchar(255),
  images       json,
  marketplaces json,
  descriptions json
);

-- +goose Down
drop table products;
