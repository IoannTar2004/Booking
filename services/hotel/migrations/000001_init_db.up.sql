create table hotels(
    id bigserial primary key,
    name varchar(64) not null,
    address varchar(100) not null,
    stars int check ( stars between 0 and 7)
);