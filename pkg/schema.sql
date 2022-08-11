create table if not exists currency_pair (
    id serial not null primary key,
    pair varchar(10)
);

create table if not exists history_deal (
    deal_id uuid not null,
    id_pair int not null primary key,
    deal_time time not null,
    coast float not null,

);
