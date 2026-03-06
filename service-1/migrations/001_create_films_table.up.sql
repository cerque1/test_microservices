CREATE TABLE films (
    id serial primary key,
    name varchar not null,
    length int not null,
    release_date date not null
)