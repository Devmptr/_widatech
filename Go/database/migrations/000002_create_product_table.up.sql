create table products
(
    invoice_no      text    not null,
    item_name       text    not null,
    quantity        int     not null,
    total_cogs      int     not null    default 0,
    total_price     int     not null    default 0
);