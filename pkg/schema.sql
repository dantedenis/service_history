set enable_seqscan to off;

--drop table currency_pair;
--drop table history_deal;

create table if not exists currency_pair (
    id serial not null primary key,
    pair varchar(10) unique
);
create unique index on currency_pair(pair);

create table if not exists history_deal (
    id serial not null primary key,
    pair_id int not null,
    deal_time timestamp not null,
    coast float not null,
    constraint fk_pair_id foreign key (pair_id) references currency_pair(id)
);
create index on history_deal(pair_id);
create unique index on history_deal(pair_id, deal_time);

--EURUSD,USDRUB,USDJPY,EURRUB,GBPJPY
insert into currency_pair(pair)
    values
        ('EURUSD'),
        ('USDRUB'),
        ('USDJPY'),
        ('EURRUB'),
        ('GBPJPY');
