create table plant
(
    id         SERIAL PRIMARY KEY,
    ident      VARCHAR NOT NULL,
    name       VARCHAR NOT NULL,
    price      VARCHAR NOT NULL,
    status     VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

create table porder
(
    id_order         SERIAL PRIMARY KEY,
    ident_order      VARCHAR NOT NULL,
    name_order       VARCHAR NOT NULL,
    price_order      VARCHAR NOT NULL,
    status_order     VARCHAR NOT NULL,
    start_order      VARCHAR NOT NULL,
    end_order        VARCHAR NOT NULL,
    created_at       TIMESTAMP DEFAULT NOW()
);


insert into plant (ident, name, price, status) values ('123', 'Google', 33, 'education');
insert into plant (ident, name, price, status) values ('565', 'd', 2, 'f');
insert into plant (ident, name, price, status) values ('888', 'd', 8, 'g');

insert into porder (ident_order, name_order, price_order, status_order, start_order, end_order) values ('123', 'Google', '33', 'education', '24', '48');
insert into porder (ident_order, name_order, price_order, status_order, start_order, end_order) values ('565', 'd', '2', 'f', '45', '46');
insert into porder (ident_order, name_order, price_order, status_order, start_order, end_order) values ('888', 'd', '8', 'g', '2', '8');