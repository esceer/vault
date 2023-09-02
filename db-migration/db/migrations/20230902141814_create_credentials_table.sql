-- migrate:up
CREATE TABLE credentials (
    id integer primary key autoincrement, 
    user varchar(255),
    key varchar(255),
    secret blob
);


-- migrate:down
DROP TABLE credentials IF EXISTS safe;
