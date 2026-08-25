create table cities(
    id serial primary key,
    name varchar(32)
);

create table hotels(
    id bigserial primary key,
    name varchar(64) not null,
    city_id int references cities(id),
    address varchar(100) not null,
    stars int check ( stars between 0 and 7)
);