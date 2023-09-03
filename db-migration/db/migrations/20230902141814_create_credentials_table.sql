-- migrate:up
CREATE TABLE credentials (
    id integer primary key autoincrement, 
    user varchar(255) not null,
    site varchar(255) not null,
    secret blob not null,
    created_at datetime default current_timestamp
);


-- migrate:down
DROP TABLE credentials IF EXISTS safe;
