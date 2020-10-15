CREATE TABLE links (
    id bigserial not null primary key,
    initial_link varchar not null,
    shortened_link varchar not null
);