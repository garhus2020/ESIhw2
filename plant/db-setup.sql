create table plant
(
    id         SERIAL PRIMARY KEY,
    ident      VARCHAR NOT NULL,
    name       VARCHAR NOT NULL,
    price      VARCHAR NOT NULL,
    status     VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

insert into plant (ident, name, price, status) values ('123', 'Google', 33, 'education');
insert into plant (ident, name, price, status) values ('565', 'd', 2, 'f');
insert into plant (ident, name, price, status) values ('888', 'd', 8, 'g');